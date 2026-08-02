// Package profile defines specification workflow profiles.
//
// A Profile bundles spec requirements, synthesis rules, and evaluation
// criteria into a cohesive workflow configuration (e.g., "aws-product",
// "big-tech-feature", "pbhq-lite").
package profile

import "github.com/ProductBuildersHQ/specification-workflow-spec/pkg/synthesis"

// Profile represents a complete specification workflow configuration.
type Profile struct {
	// Name is the profile identifier (e.g., "aws-product", "pbhq-lite").
	Name string `json:"name" jsonschema:"required,description=Profile identifier"`

	// Description explains the profile's purpose and use case.
	Description string `json:"description,omitempty" jsonschema:"description=Profile purpose and target audience"`

	// Extends is the name of a parent profile to inherit from.
	Extends string `json:"extends,omitempty" jsonschema:"description=Parent profile to inherit settings from"`

	// Abstract indicates this profile is a base for other profiles (not directly usable).
	Abstract bool `json:"abstract,omitempty" jsonschema:"description=True if this profile cannot be used directly"`

	// Methodology documents the underlying product methodology.
	Methodology *Methodology `json:"methodology,omitempty" jsonschema:"description=Underlying product methodology documentation"`

	// SpecConfig defines which specs are required/optional.
	SpecConfig map[string]*SpecRequirement `json:"spec_config" jsonschema:"required,description=Spec requirements by spec type ID"`

	// Synthesis defines how specs are generated from other specs.
	Synthesis map[string]*synthesis.Rule `json:"synthesis,omitempty" jsonschema:"description=Synthesis rules by target spec type"`

	// Workflow defines the ordered phases and gates.
	Workflow *Workflow `json:"workflow,omitempty" jsonschema:"description=Phase ordering and gates"`

	// Evaluation defines pass/fail thresholds.
	Evaluation *EvaluationConfig `json:"evaluation,omitempty" jsonschema:"description=Evaluation thresholds"`
}

// Methodology documents the underlying product development methodology.
type Methodology struct {
	// Name is the methodology name (e.g., "Amazon Working Backwards").
	Name string `json:"name" jsonschema:"required,description=Methodology name"`

	// Description explains the methodology.
	Description string `json:"description,omitempty" jsonschema:"description=Methodology overview"`

	// Creator is the person/company who created the methodology.
	Creator string `json:"creator,omitempty" jsonschema:"description=Methodology creator"`

	// Reference is a URL to the canonical methodology documentation.
	Reference string `json:"reference,omitempty" jsonschema:"format=uri,description=URL to methodology documentation"`

	// Principles are the core principles of the methodology.
	Principles []Principle `json:"principles,omitempty" jsonschema:"description=Core methodology principles"`

	// Artifacts are the key artifacts produced by the methodology.
	Artifacts []string `json:"artifacts,omitempty" jsonschema:"description=Key artifacts"`
}

// Principle is a named principle with description.
type Principle struct {
	// ID is the principle identifier (e.g., "customer_obsession").
	ID string `json:"id" jsonschema:"required,description=Principle identifier"`

	// Name is the human-readable name.
	Name string `json:"name" jsonschema:"required,description=Principle name"`

	// Description explains the principle.
	Description string `json:"description,omitempty" jsonschema:"description=Principle explanation"`

	// Source is the origin (e.g., "Amazon", "Google").
	Source string `json:"source,omitempty" jsonschema:"description=Origin company or methodology"`
}

// SpecRequirement defines whether a spec type is required and its configuration.
type SpecRequirement struct {
	// Required indicates whether this spec must be present.
	Required bool `json:"required" jsonschema:"description=Whether this spec is required"`

	// Category overrides the default category for this spec type.
	Category string `json:"category,omitempty" jsonschema:"enum=source,enum=gtm,enum=technical,enum=execution,enum=output,enum=strategic"`

	// Description provides profile-specific context for this spec.
	Description string `json:"description,omitempty" jsonschema:"description=Profile-specific description"`

	// Template specifies a custom template path.
	Template string `json:"template,omitempty" jsonschema:"description=Custom template path"`

	// Rubric specifies a custom rubric path.
	Rubric string `json:"rubric,omitempty" jsonschema:"description=Custom rubric path"`
}

// Workflow defines the ordered execution of specs.
type Workflow struct {
	// Sequence is the ordered list of spec types to produce.
	Sequence []string `json:"sequence,omitempty" jsonschema:"description=Ordered spec type IDs"`

	// Phases groups specs into named phases.
	Phases []Phase `json:"phases,omitempty" jsonschema:"description=Named workflow phases"`

	// IterationTrigger is the spec type that triggers iteration.
	IterationTrigger string `json:"iteration_trigger,omitempty" jsonschema:"description=Spec type that triggers workflow iteration"`

	// ReviewGates are approval checkpoints.
	ReviewGates []ReviewGate `json:"review_gates,omitempty" jsonschema:"description=Approval checkpoints"`
}

// Phase is a named group of specs in the workflow.
type Phase struct {
	// ID is the phase identifier.
	ID string `json:"id" jsonschema:"required,description=Phase identifier"`

	// Name is the human-readable phase name.
	Name string `json:"name" jsonschema:"required,description=Phase name"`

	// Description explains the phase purpose.
	Description string `json:"description,omitempty" jsonschema:"description=Phase purpose"`

	// Specs are the spec types in this phase.
	Specs []string `json:"specs" jsonschema:"required,description=Spec type IDs in this phase"`
}

// ReviewGate is an approval checkpoint after a spec.
type ReviewGate struct {
	// After is the spec type after which this gate applies.
	After string `json:"after" jsonschema:"required,description=Spec type ID after which gate applies"`

	// Action is the required action (e.g., "stakeholder_review", "tech_lead_review").
	Action string `json:"action" jsonschema:"required,description=Required approval action"`

	// Required indicates whether passing this gate is mandatory.
	Required bool `json:"required,omitempty" jsonschema:"description=Whether gate is mandatory"`
}

// EvaluationConfig defines pass/fail thresholds.
type EvaluationConfig struct {
	// PassThreshold is the minimum score (0-100) to pass.
	PassThreshold int `json:"pass_threshold,omitempty" jsonschema:"minimum=0,maximum=100,description=Minimum score to pass"`

	// PartialThreshold is the minimum score (0-100) for partial pass.
	PartialThreshold int `json:"partial_threshold,omitempty" jsonschema:"minimum=0,maximum=100,description=Minimum score for partial pass"`

	// MaxFindingsSeverity defines maximum allowed findings by severity.
	MaxFindingsSeverity *FindingSeverityLimits `json:"max_findings_severity,omitempty" jsonschema:"description=Maximum findings allowed by severity"`
}

// FindingSeverityLimits defines maximum findings by severity level.
type FindingSeverityLimits struct {
	Critical int `json:"critical,omitempty" jsonschema:"description=Max critical findings (-1 = unlimited)"`
	High     int `json:"high,omitempty" jsonschema:"description=Max high findings (-1 = unlimited)"`
	Medium   int `json:"medium,omitempty" jsonschema:"description=Max medium findings (-1 = unlimited)"`
	Low      int `json:"low,omitempty" jsonschema:"description=Max low findings (-1 = unlimited)"`
}
