package workflows

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultLoader(t *testing.T) {
	loader := DefaultLoader()

	t.Run("loads embedded workflow", func(t *testing.T) {
		w, err := loader.Load("aws-feature")
		if err != nil {
			t.Fatalf("Load(aws-feature) error: %v", err)
		}
		if w.Workflow.Name != "aws-feature" {
			t.Errorf("Name = %q, want %q", w.Workflow.Name, "aws-feature")
		}
	})

	t.Run("resolves inheritance", func(t *testing.T) {
		w, err := loader.Load("aws-feature")
		if err != nil {
			t.Fatalf("Load(aws-feature) error: %v", err)
		}

		// aws-feature extends enterprise, so should have enterprise's specs merged
		if w.Workflow.SpecConfig == nil {
			t.Fatal("SpecConfig is nil after inheritance resolution")
		}

		// Check that enterprise's "mrd" requirement is inherited
		if _, ok := w.Workflow.SpecConfig["mrd"]; !ok {
			t.Error("Expected inherited 'mrd' spec config from enterprise")
		}

		// Check that aws-feature's "opportunity-spec" is present
		if _, ok := w.Workflow.SpecConfig["opportunity-spec"]; !ok {
			t.Error("Expected 'opportunity-spec' from aws-feature")
		}
	})

	t.Run("inherits templates from parent", func(t *testing.T) {
		w, err := loader.Load("aws-feature")
		if err != nil {
			t.Fatalf("Load(aws-feature) error: %v", err)
		}

		// Should have templates from both aws-feature and enterprise
		if len(w.Templates) == 0 {
			t.Error("Expected templates after inheritance")
		}
	})

	t.Run("lists available workflows", func(t *testing.T) {
		available := loader.Available()
		if len(available) < 3 {
			t.Errorf("Available() = %d, want at least 3", len(available))
		}
	})
}

// TestAllEmbeddedWorkflowsResolve verifies every embedded workflow parses and
// its full inheritance chain resolves without error.
func TestAllEmbeddedWorkflowsResolve(t *testing.T) {
	loader := DefaultLoader()
	names := loader.Available()
	if len(names) < 24 {
		t.Errorf("Available() = %d workflows, want at least 24", len(names))
	}

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			w, err := loader.Load(name)
			if err != nil {
				t.Fatalf("Load(%q) error: %v", name, err)
			}
			if w.Workflow.Name != name {
				t.Errorf("Name = %q, want %q", w.Workflow.Name, name)
			}
			if w.Workflow.Extends != "" {
				t.Errorf("Extends = %q after resolution, want empty", w.Workflow.Extends)
			}
			for specType, rs := range w.Rubrics {
				if len(rs.Categories) == 0 {
					t.Errorf("rubric %q has no categories", specType)
				}
			}
		})
	}
}

func TestChainLoader(t *testing.T) {
	// Create a temp directory with a custom workflow
	tmpDir, err := os.MkdirTemp("", "workflow-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create custom workflow directory
	customDir := filepath.Join(tmpDir, "custom-workflow")
	if err := os.MkdirAll(customDir, 0755); err != nil {
		t.Fatalf("Failed to create custom workflow dir: %v", err)
	}

	// Write profile.yaml
	profileYAML := `name: custom-workflow
description: A custom workflow for testing
`
	if err := os.WriteFile(filepath.Join(customDir, "profile.yaml"), []byte(profileYAML), 0600); err != nil {
		t.Fatalf("Failed to write profile.yaml: %v", err)
	}

	// Create chain loader: file first, then embedded
	loader := NewChainLoader(
		NewFileLoader(tmpDir),
		&embeddedLoader{},
	)

	t.Run("loads from file loader first", func(t *testing.T) {
		w, err := loader.Load("custom-workflow")
		if err != nil {
			t.Fatalf("Load(custom-workflow) error: %v", err)
		}
		if w.Workflow.Name != "custom-workflow" {
			t.Errorf("Name = %q, want %q", w.Workflow.Name, "custom-workflow")
		}
	})

	t.Run("falls back to embedded", func(t *testing.T) {
		w, err := loader.Load("aws-feature")
		if err != nil {
			t.Fatalf("Load(aws-feature) error: %v", err)
		}
		if w.Workflow.Name != "aws-feature" {
			t.Errorf("Name = %q, want %q", w.Workflow.Name, "aws-feature")
		}
	})

	t.Run("combines available from all loaders", func(t *testing.T) {
		available := loader.Available()

		hasCustom := false
		hasEmbedded := false
		for _, name := range available {
			if name == "custom-workflow" {
				hasCustom = true
			}
			if name == "aws-feature" {
				hasEmbedded = true
			}
		}

		if !hasCustom {
			t.Error("Available() missing custom-workflow")
		}
		if !hasEmbedded {
			t.Error("Available() missing aws-feature")
		}
	})
}

func TestFileLoader(t *testing.T) {
	// Create a temp directory with a workflow
	tmpDir, err := os.MkdirTemp("", "workflow-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create workflow directory structure
	workflowDir := filepath.Join(tmpDir, "test-workflow")
	templatesDir := filepath.Join(workflowDir, "templates")
	rubricsDir := filepath.Join(workflowDir, "rubrics")

	for _, dir := range []string{workflowDir, templatesDir, rubricsDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create dir: %v", err)
		}
	}

	// Write profile.yaml
	profileYAML := `name: test-workflow
description: Test workflow
spec_config:
  prd:
    required: true
    category: source
`
	if err := os.WriteFile(filepath.Join(workflowDir, "profile.yaml"), []byte(profileYAML), 0600); err != nil {
		t.Fatalf("Failed to write profile.yaml: %v", err)
	}

	// Write template
	templateContent := "# PRD Template\n\nThis is a test template."
	if err := os.WriteFile(filepath.Join(templatesDir, "prd.md"), []byte(templateContent), 0600); err != nil {
		t.Fatalf("Failed to write template: %v", err)
	}

	// Write a template with Windows line endings (as produced by
	// core.autocrlf=true checkouts) to verify normalization.
	crlfContent := "# TRD Template\r\n\r\nCRLF line endings."
	if err := os.WriteFile(filepath.Join(templatesDir, "trd.md"), []byte(crlfContent), 0600); err != nil {
		t.Fatalf("Failed to write CRLF template: %v", err)
	}

	// Write rubric
	rubricYAML := `id: prd-rubric
name: PRD Rubric
evaluationType: analytic
categories:
  - id: completeness
    name: Completeness
    weight: 1.0
`
	if err := os.WriteFile(filepath.Join(rubricsDir, "prd.rubric.yaml"), []byte(rubricYAML), 0600); err != nil {
		t.Fatalf("Failed to write rubric: %v", err)
	}

	loader := NewFileLoader(tmpDir)

	t.Run("loads workflow", func(t *testing.T) {
		w, err := loader.Load("test-workflow")
		if err != nil {
			t.Fatalf("Load error: %v", err)
		}
		if w.Workflow.Name != "test-workflow" {
			t.Errorf("Name = %q, want %q", w.Workflow.Name, "test-workflow")
		}
	})

	t.Run("loads templates", func(t *testing.T) {
		w, err := loader.Load("test-workflow")
		if err != nil {
			t.Fatalf("Load error: %v", err)
		}
		tmpl, ok := w.Templates["prd"]
		if !ok {
			t.Fatal("Template 'prd' not found")
		}
		if tmpl.Content != templateContent {
			t.Errorf("Template content mismatch")
		}
	})

	t.Run("normalizes CRLF template content", func(t *testing.T) {
		w, err := loader.Load("test-workflow")
		if err != nil {
			t.Fatalf("Load error: %v", err)
		}
		tmpl, ok := w.Templates["trd"]
		if !ok {
			t.Fatal("Template 'trd' not found")
		}
		if want := "# TRD Template\n\nCRLF line endings."; tmpl.Content != want {
			t.Errorf("Content = %q, want %q", tmpl.Content, want)
		}
	})

	t.Run("loads rubrics", func(t *testing.T) {
		w, err := loader.Load("test-workflow")
		if err != nil {
			t.Fatalf("Load error: %v", err)
		}
		rubric, ok := w.Rubrics["prd"]
		if !ok {
			t.Fatal("Rubric 'prd' not found")
		}
		if rubric.ID != "prd-rubric" {
			t.Errorf("Rubric ID = %q, want %q", rubric.ID, "prd-rubric")
		}
	})

	t.Run("lists available", func(t *testing.T) {
		available := loader.Available()
		if len(available) != 1 || available[0] != "test-workflow" {
			t.Errorf("Available() = %v, want [test-workflow]", available)
		}
	})
}

func TestResolvingLoader_CircularInheritance(t *testing.T) {
	// Create workflows that reference each other
	tmpDir, err := os.MkdirTemp("", "workflow-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// workflow-a extends workflow-b
	workflowA := filepath.Join(tmpDir, "workflow-a")
	if err := os.MkdirAll(workflowA, 0750); err != nil {
		t.Fatalf("Failed to create workflow-a dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(workflowA, "profile.yaml"), []byte(`name: workflow-a
extends: workflow-b
`), 0600); err != nil {
		t.Fatalf("Failed to write workflow-a profile: %v", err)
	}

	// workflow-b extends workflow-a (circular!)
	workflowB := filepath.Join(tmpDir, "workflow-b")
	if err := os.MkdirAll(workflowB, 0750); err != nil {
		t.Fatalf("Failed to create workflow-b dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(workflowB, "profile.yaml"), []byte(`name: workflow-b
extends: workflow-a
`), 0600); err != nil {
		t.Fatalf("Failed to write workflow-b profile: %v", err)
	}

	loader := NewResolvingLoader(NewFileLoader(tmpDir))

	_, err = loader.Load("workflow-a")
	if err == nil {
		t.Error("Expected error for circular inheritance")
	}
	if err != nil && !contains(err.Error(), "circular inheritance") {
		t.Errorf("Expected circular inheritance error, got: %v", err)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsAt(s, substr, 0))
}

func containsAt(s, substr string, start int) bool {
	for i := start; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
