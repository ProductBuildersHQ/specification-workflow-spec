package diagram

import (
	"strings"
	"testing"

	"github.com/ProductBuildersHQ/specification-workflow-spec/pkg/profile"
	"github.com/ProductBuildersHQ/specification-workflow-spec/pkg/synthesis"
)

func TestGenerateD2(t *testing.T) {
	p := &profile.Profile{
		Name: "aws-product",
		SpecConfig: map[string]*profile.SpecRequirement{
			"mrd":   {Required: true, Category: "source"},
			"press": {Required: true, Category: "gtm"},
			"faq":   {Required: true, Category: "gtm"},
			"prd":   {Required: true, Category: "source"},
		},
		Synthesis: map[string]*synthesis.Rule{
			"press": {Sources: []string{"mrd"}},
			"faq":   {Sources: []string{"mrd", "press"}},
			"prd":   {Sources: []string{"mrd", "press", "faq"}},
		},
	}

	opts := DefaultOptions()
	opts.Title = "AWS Product Flow"

	d2, err := Generate(p, FormatD2, opts)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	// Check for expected content
	if !strings.Contains(d2, "AWS Product Flow") {
		t.Error("D2 output missing title")
	}
	if !strings.Contains(d2, "mrd -> press: synthesize") {
		t.Error("D2 output missing mrd -> press edge")
	}
	if !strings.Contains(d2, "press -> faq: synthesize") {
		t.Error("D2 output missing press -> faq edge")
	}
}

func TestGenerateMermaid(t *testing.T) {
	p := &profile.Profile{
		Name: "aws-product",
		SpecConfig: map[string]*profile.SpecRequirement{
			"mrd":   {Required: true, Category: "source"},
			"press": {Required: true, Category: "gtm"},
			"faq":   {Required: true, Category: "gtm"},
		},
		Synthesis: map[string]*synthesis.Rule{
			"press": {Sources: []string{"mrd"}},
			"faq":   {Sources: []string{"mrd", "press"}},
		},
	}

	opts := DefaultOptions()
	mermaid, err := Generate(p, FormatMermaid, opts)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	if !strings.Contains(mermaid, "flowchart TB") {
		t.Error("Mermaid output missing flowchart header")
	}
	if !strings.Contains(mermaid, "mrd -->|synthesize| press") {
		t.Error("Mermaid output missing mrd -> press edge")
	}
}

func TestGenerateFromDAG(t *testing.T) {
	dag := &synthesis.DAG{
		Nodes: []string{"mrd", "press", "faq"},
		Edges: []synthesis.Edge{
			{Source: "mrd", Target: "press"},
			{Source: "mrd", Target: "faq"},
			{Source: "press", Target: "faq"},
		},
	}

	d2, err := GenerateFromDAG(dag, FormatD2, "Test DAG")
	if err != nil {
		t.Fatalf("GenerateFromDAG failed: %v", err)
	}

	if !strings.Contains(d2, "mrd -> press") {
		t.Error("DAG D2 output missing edge")
	}
}
