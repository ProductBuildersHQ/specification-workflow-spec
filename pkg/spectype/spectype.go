// Package spectype defines the registry of specification document types.
//
// Each SpecType represents a distinct document artifact in a specification
// workflow (e.g., PRD, MRD, Press Release, FAQ, 6-Pager).
package spectype

// Category groups spec types by their role in the workflow.
type Category string

const (
	// CategorySource represents human-authored discovery documents.
	CategorySource Category = "source"

	// CategoryGTM represents go-to-market documents (often synthesized).
	CategoryGTM Category = "gtm"

	// CategoryTechnical represents technical specification documents.
	CategoryTechnical Category = "technical"

	// CategoryExecution represents execution tracking documents.
	CategoryExecution Category = "execution"

	// CategoryOutput represents reconciled output documents.
	CategoryOutput Category = "output"

	// CategoryStrategic represents strategic planning documents (e.g., V2MOM).
	CategoryStrategic Category = "strategic"
)

// AuthorshipMode indicates whether a spec is typically human-authored or LLM-synthesized.
type AuthorshipMode string

const (
	// AuthorshipHuman indicates the spec is primarily human-authored.
	AuthorshipHuman AuthorshipMode = "human"

	// AuthorshipSynthesized indicates the spec is primarily LLM-synthesized.
	AuthorshipSynthesized AuthorshipMode = "synthesized"

	// AuthorshipHybrid indicates the spec involves both human and LLM authorship.
	AuthorshipHybrid AuthorshipMode = "hybrid"
)

// SpecType defines a specification document type in the registry.
type SpecType struct {
	// ID is the canonical identifier (e.g., "prd", "press", "narrative-6p").
	ID string `json:"id" jsonschema:"required,description=Canonical identifier for the spec type"`

	// Name is the human-readable name (e.g., "Product Requirements Document").
	Name string `json:"name" jsonschema:"required,description=Human-readable name"`

	// ShortName is an abbreviated name for UI display (e.g., "PRD", "6-Pager").
	ShortName string `json:"shortName,omitempty" jsonschema:"description=Abbreviated name for compact display"`

	// Description explains the purpose of this spec type.
	Description string `json:"description,omitempty" jsonschema:"description=Purpose and usage of this spec type"`

	// Category groups this spec type (source, gtm, technical, etc.).
	Category Category `json:"category" jsonschema:"required,enum=source,enum=gtm,enum=technical,enum=execution,enum=output,enum=strategic"`

	// Authorship indicates typical authorship mode.
	Authorship AuthorshipMode `json:"authorship" jsonschema:"required,enum=human,enum=synthesized,enum=hybrid"`

	// Filename is the canonical filename (e.g., "PRD.md", "press.md").
	Filename string `json:"filename" jsonschema:"required,description=Canonical filename for this spec"`

	// EvalFilename is the filename for evaluation results (e.g., "PRD.eval.json").
	EvalFilename string `json:"evalFilename,omitempty" jsonschema:"description=Filename for LLM-as-Judge evaluation results"`

	// Origins lists the methodologies that use this spec type.
	Origins []string `json:"origins,omitempty" jsonschema:"description=Methodologies that originated or use this spec type"`

	// Aliases are alternative names for this spec type.
	Aliases []string `json:"aliases,omitempty" jsonschema:"description=Alternative names or abbreviations"`
}

// SpecTypeRegistry is a collection of spec type definitions.
type SpecTypeRegistry struct {
	// Version is the schema version for this registry.
	Version string `json:"version" jsonschema:"required,description=Schema version"`

	// Types is the list of registered spec types.
	Types []SpecType `json:"types" jsonschema:"required,description=Registered specification types"`
}

// CoreSpecTypes returns the built-in spec type definitions.
func CoreSpecTypes() []SpecType {
	return []SpecType{
		// Source specs (human-authored)
		{
			ID:           "mrd",
			Name:         "Market Requirements Document",
			ShortName:    "MRD",
			Description:  "Defines the market problem, customer segments, and business opportunity.",
			Category:     CategorySource,
			Authorship:   AuthorshipHuman,
			Filename:     "MRD.md",
			EvalFilename: "MRD.eval.json",
			Origins:      []string{"enterprise", "aws-product", "big-tech-product"},
		},
		{
			ID:           "prd",
			Name:         "Product Requirements Document",
			ShortName:    "PRD",
			Description:  "Defines product goals, user stories, and functional requirements.",
			Category:     CategorySource,
			Authorship:   AuthorshipHuman,
			Filename:     "PRD.md",
			EvalFilename: "PRD.eval.json",
			Origins:      []string{"startup", "enterprise", "big-tech"},
		},
		{
			ID:           "uxd",
			Name:         "User Experience Design",
			ShortName:    "UXD",
			Description:  "Defines user journeys, wireframes, and interaction patterns.",
			Category:     CategorySource,
			Authorship:   AuthorshipHuman,
			Filename:     "UXD.md",
			EvalFilename: "UXD.eval.json",
			Origins:      []string{"design-thinking", "big-tech"},
		},
		{
			ID:           "opportunity-spec",
			Name:         "Opportunity Specification",
			ShortName:    "OpportunitySpec",
			Description:  "12-box canvas combining Patton discovery + Cagan business case for features.",
			Category:     CategorySource,
			Authorship:   AuthorshipHuman,
			Filename:     "OPPORTUNITY-SPEC.md",
			EvalFilename: "OPPORTUNITY-SPEC.eval.json",
			Origins:      []string{"aws-feature", "big-tech-feature"},
			Aliases:      []string{"opportunity-canvas"},
		},

		// GTM specs (synthesized)
		{
			ID:           "press",
			Name:         "Press Release",
			ShortName:    "Press",
			Description:  "Working Backwards customer announcement written before building.",
			Category:     CategoryGTM,
			Authorship:   AuthorshipSynthesized,
			Filename:     "PRESS.md",
			EvalFilename: "PRESS.eval.json",
			Origins:      []string{"aws-product", "aws-feature", "big-tech"},
			Aliases:      []string{"press-release", "pr"},
		},
		{
			ID:           "faq",
			Name:         "Frequently Asked Questions",
			ShortName:    "FAQ",
			Description:  "Challenges assumptions with internal and external questions.",
			Category:     CategoryGTM,
			Authorship:   AuthorshipSynthesized,
			Filename:     "FAQ.md",
			EvalFilename: "FAQ.eval.json",
			Origins:      []string{"aws-product", "aws-feature", "big-tech"},
		},
		{
			ID:           "narrative-6p",
			Name:         "Six-Pager Narrative",
			ShortName:    "6-Pager",
			Description:  "Amazon-style narrative document for stakeholder alignment.",
			Category:     CategoryGTM,
			Authorship:   AuthorshipSynthesized,
			Filename:     "NARRATIVE-6P.md",
			EvalFilename: "NARRATIVE-6P.eval.json",
			Origins:      []string{"aws-product", "big-tech-product"},
			Aliases:      []string{"6-pager", "six-pager", "narrative"},
		},
		{
			ID:           "narrative-1p",
			Name:         "One-Pager Executive Summary",
			ShortName:    "1-Pager",
			Description:  "Executive summary for quick stakeholder review.",
			Category:     CategoryGTM,
			Authorship:   AuthorshipSynthesized,
			Filename:     "NARRATIVE-1P.md",
			EvalFilename: "NARRATIVE-1P.eval.json",
			Origins:      []string{"enterprise", "big-tech"},
			Aliases:      []string{"1-pager", "one-pager", "exec-summary"},
		},
		{
			ID:           "bmc",
			Name:         "Business Model Canvas",
			ShortName:    "BMC",
			Description:  "9-block business model visualization extracted from MRD.",
			Category:     CategoryGTM,
			Authorship:   AuthorshipSynthesized,
			Filename:     "BMC.md",
			EvalFilename: "BMC.eval.json",
			Origins:      []string{"enterprise", "lean-startup"},
		},

		// Technical specs (synthesized)
		{
			ID:           "trd",
			Name:         "Technical Requirements Document",
			ShortName:    "TRD",
			Description:  "Defines architecture, components, APIs, and technical approach.",
			Category:     CategoryTechnical,
			Authorship:   AuthorshipSynthesized,
			Filename:     "TRD.md",
			EvalFilename: "TRD.eval.json",
			Origins:      []string{"enterprise", "google", "big-tech"},
			Aliases:      []string{"design-doc", "technical-design"},
		},
		{
			ID:           "tpd",
			Name:         "Test Plan Document",
			ShortName:    "TPD",
			Description:  "Defines test strategy, coverage requirements, and validation approach.",
			Category:     CategoryTechnical,
			Authorship:   AuthorshipSynthesized,
			Filename:     "TPD.md",
			EvalFilename: "TPD.eval.json",
			Origins:      []string{"enterprise", "big-tech"},
			Aliases:      []string{"test-plan"},
		},
		{
			ID:           "ird",
			Name:         "Infrastructure Requirements Document",
			ShortName:    "IRD",
			Description:  "Defines deployment, scaling, and operational requirements.",
			Category:     CategoryTechnical,
			Authorship:   AuthorshipSynthesized,
			Filename:     "IRD.md",
			EvalFilename: "IRD.eval.json",
			Origins:      []string{"enterprise", "big-tech-product"},
			Aliases:      []string{"infra-doc"},
		},

		// Execution specs
		{
			ID:          "plan",
			Name:        "Implementation Plan",
			ShortName:   "PLAN",
			Description: "Sequences work into phases with dependencies and milestones.",
			Category:    CategoryExecution,
			Authorship:  AuthorshipHybrid,
			Filename:    "PLAN.md",
			Origins:     []string{"pbhq-lite"},
		},
		{
			ID:          "roadmap",
			Name:        "Roadmap",
			ShortName:   "ROADMAP",
			Description: "Tracks RMIs (Roadmap Items) with status and progress.",
			Category:    CategoryExecution,
			Authorship:  AuthorshipHybrid,
			Filename:    "ROADMAP.md",
			Origins:     []string{"pbhq-lite"},
		},

		// Output specs
		{
			ID:          "spec",
			Name:        "Reconciled Specification",
			ShortName:   "SPEC",
			Description: "Unified execution spec reconciling all source documents.",
			Category:    CategoryOutput,
			Authorship:  AuthorshipSynthesized,
			Filename:    "spec.md",
			Origins:     []string{"enterprise"},
		},

		// Methodology-specific specs
		{
			ID:          "hypothesis",
			Name:        "Hypothesis Document",
			ShortName:   "Hypothesis",
			Description: "Core hypothesis to validate in Build-Measure-Learn cycles.",
			Category:    CategorySource,
			Authorship:  AuthorshipHuman,
			Filename:    "HYPOTHESIS.md",
			Origins:     []string{"lean-startup", "0-1"},
		},
		{
			ID:          "shapeup-pitch",
			Name:        "Shape Up Pitch",
			ShortName:   "Pitch",
			Description: "Shaped problem and solution with appetite (Basecamp).",
			Category:    CategorySource,
			Authorship:  AuthorshipHuman,
			Filename:    "PITCH.md",
			Origins:     []string{"shapeup"},
		},
		{
			ID:          "ost",
			Name:        "Opportunity Solution Tree",
			ShortName:   "OST",
			Description: "Maps outcomes to opportunities to solutions (Teresa Torres).",
			Category:    CategorySource,
			Authorship:  AuthorshipHuman,
			Filename:    "OST.md",
			Origins:     []string{"continuous-discovery"},
		},
	}
}
