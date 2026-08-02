package profile

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// ParseYAML parses a Profile from YAML bytes.
func ParseYAML(data []byte) (*Profile, error) {
	var p Profile
	if err := yaml.Unmarshal(data, &p); err != nil {
		return nil, fmt.Errorf("parsing profile YAML: %w", err)
	}
	return &p, nil
}

// ParseYAMLFile parses a Profile from a YAML file path.
func ParseYAMLFile(path string) (*Profile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading profile file: %w", err)
	}
	return ParseYAML(data)
}

// ParseYAMLFromFS parses a Profile from an fs.FS at the given path.
func ParseYAMLFromFS(fsys fs.FS, path string) (*Profile, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("reading profile from FS: %w", err)
	}
	return ParseYAML(data)
}

// LoadFromFS loads a profile directory from an fs.FS.
// Expects profile.yaml at the root of the directory.
func LoadFromFS(fsys fs.FS, dir string) (*Profile, error) {
	profilePath := filepath.Join(dir, "profile.yaml")
	return ParseYAMLFromFS(fsys, profilePath)
}

// ToYAML serializes a Profile to YAML bytes.
func (p *Profile) ToYAML() ([]byte, error) {
	data, err := yaml.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("marshaling profile: %w", err)
	}
	return data, nil
}

// Clone creates a deep copy of the Profile.
func (p *Profile) Clone() *Profile {
	if p == nil {
		return nil
	}

	clone := &Profile{
		Name:        p.Name,
		Description: p.Description,
		Extends:     p.Extends,
		Abstract:    p.Abstract,
	}

	// Clone Methodology
	if p.Methodology != nil {
		clone.Methodology = &Methodology{
			Name:        p.Methodology.Name,
			Description: p.Methodology.Description,
			Creator:     p.Methodology.Creator,
			Reference:   p.Methodology.Reference,
		}
		if p.Methodology.Principles != nil {
			clone.Methodology.Principles = make([]Principle, len(p.Methodology.Principles))
			copy(clone.Methodology.Principles, p.Methodology.Principles)
		}
		if p.Methodology.Artifacts != nil {
			clone.Methodology.Artifacts = make([]string, len(p.Methodology.Artifacts))
			copy(clone.Methodology.Artifacts, p.Methodology.Artifacts)
		}
	}

	// Clone SpecConfig
	if p.SpecConfig != nil {
		clone.SpecConfig = make(map[string]*SpecRequirement)
		for k, v := range p.SpecConfig {
			clone.SpecConfig[k] = &SpecRequirement{
				Required:    v.Required,
				Category:    v.Category,
				Description: v.Description,
				Template:    v.Template,
				Rubric:      v.Rubric,
			}
		}
	}

	// Clone Synthesis
	if p.Synthesis != nil {
		clone.Synthesis = make(map[string]*SynthesisRule)
		for k, v := range p.Synthesis {
			rule := &SynthesisRule{
				Guidance:      v.Guidance,
				PromptContext: v.PromptContext,
				Required:      v.Required,
				Priority:      v.Priority,
			}
			if v.Sources != nil {
				rule.Sources = make([]string, len(v.Sources))
				copy(rule.Sources, v.Sources)
			}
			clone.Synthesis[k] = rule
		}
	}

	// Clone Workflow
	if p.Workflow != nil {
		clone.Workflow = &Workflow{
			IterationTrigger: p.Workflow.IterationTrigger,
		}
		if p.Workflow.Sequence != nil {
			clone.Workflow.Sequence = make([]string, len(p.Workflow.Sequence))
			copy(clone.Workflow.Sequence, p.Workflow.Sequence)
		}
		if p.Workflow.Phases != nil {
			clone.Workflow.Phases = make([]Phase, len(p.Workflow.Phases))
			copy(clone.Workflow.Phases, p.Workflow.Phases)
		}
		if p.Workflow.ReviewGates != nil {
			clone.Workflow.ReviewGates = make([]ReviewGate, len(p.Workflow.ReviewGates))
			copy(clone.Workflow.ReviewGates, p.Workflow.ReviewGates)
		}
	}

	// Clone Evaluation
	if p.Evaluation != nil {
		clone.Evaluation = &EvaluationConfig{
			PassThreshold:    p.Evaluation.PassThreshold,
			PartialThreshold: p.Evaluation.PartialThreshold,
		}
		if p.Evaluation.MaxFindingsSeverity != nil {
			clone.Evaluation.MaxFindingsSeverity = &FindingSeverityLimits{
				Critical: p.Evaluation.MaxFindingsSeverity.Critical,
				High:     p.Evaluation.MaxFindingsSeverity.High,
				Medium:   p.Evaluation.MaxFindingsSeverity.Medium,
				Low:      p.Evaluation.MaxFindingsSeverity.Low,
			}
		}
	}

	return clone
}

// Merge combines this profile with a parent profile.
// Settings from this profile override the parent.
func (p *Profile) Merge(parent *Profile) *Profile {
	if parent == nil {
		return p.Clone()
	}

	// Start with a clone of the parent
	merged := parent.Clone()
	merged.Name = p.Name
	merged.Description = p.Description
	merged.Extends = "" // Clear extends since we've resolved it

	// Override with child's methodology if present
	if p.Methodology != nil {
		merged.Methodology = p.Methodology
	}

	// Merge SpecConfig (child overrides parent)
	if p.SpecConfig != nil {
		if merged.SpecConfig == nil {
			merged.SpecConfig = make(map[string]*SpecRequirement)
		}
		for k, v := range p.SpecConfig {
			merged.SpecConfig[k] = v
		}
	}

	// Merge Synthesis (child overrides parent)
	if p.Synthesis != nil {
		if merged.Synthesis == nil {
			merged.Synthesis = make(map[string]*SynthesisRule)
		}
		for k, v := range p.Synthesis {
			merged.Synthesis[k] = v
		}
	}

	// Override workflow if child has one
	if p.Workflow != nil {
		merged.Workflow = p.Workflow
	}

	// Override evaluation if child has one
	if p.Evaluation != nil {
		merged.Evaluation = p.Evaluation
	}

	return merged
}

// Validate checks if the profile is valid.
func (p *Profile) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("profile name is required")
	}
	return nil
}

// RequiredSpecs returns the list of required spec type IDs.
func (p *Profile) RequiredSpecs() []string {
	if p.SpecConfig == nil {
		return nil
	}

	var required []string
	for id, req := range p.SpecConfig {
		if req.Required {
			required = append(required, id)
		}
	}
	return required
}

// IsRequired returns whether a spec type is required.
func (p *Profile) IsRequired(specType string) bool {
	if p.SpecConfig == nil {
		return false
	}
	if req, ok := p.SpecConfig[specType]; ok {
		return req.Required
	}
	return false
}

// GetCategory returns the category for a spec type.
func (p *Profile) GetCategory(specType string) string {
	if p.SpecConfig == nil {
		return ""
	}
	if req, ok := p.SpecConfig[specType]; ok {
		return req.Category
	}
	return ""
}
