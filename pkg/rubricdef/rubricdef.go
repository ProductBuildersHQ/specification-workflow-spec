// Package rubricdef defines evaluation rubric definitions.
//
// A rubric definition specifies the evaluation criteria for a spec type,
// distinct from the evaluation report (which is in structured-evaluation).
// This package defines how to AUTHOR rubrics, not how to REPORT results.
package rubricdef

// RubricDefinition defines evaluation criteria for a spec type.
type RubricDefinition struct {
	// ID is the rubric identifier (e.g., "prd-rubric").
	ID string `json:"id" jsonschema:"required,description=Rubric identifier"`

	// Name is the human-readable name.
	Name string `json:"name" jsonschema:"required,description=Rubric name"`

	// Version is the rubric version.
	Version string `json:"version,omitempty" jsonschema:"description=Rubric version"`

	// Description explains what this rubric evaluates.
	Description string `json:"description,omitempty" jsonschema:"description=Rubric purpose"`

	// SpecType is the spec type this rubric evaluates.
	SpecType string `json:"specType" jsonschema:"required,description=Spec type ID this rubric evaluates"`

	// EvaluationType is the type of evaluation (analytic, holistic).
	EvaluationType EvaluationType `json:"evaluationType" jsonschema:"required,enum=analytic,enum=holistic"`

	// Categories are the evaluation dimensions.
	Categories []Category `json:"categories" jsonschema:"required,description=Evaluation categories"`

	// PassCriteria defines what constitutes a passing evaluation.
	PassCriteria *PassCriteria `json:"passCriteria,omitempty" jsonschema:"description=Pass/fail criteria"`

	// Methodology is the source methodology influencing this rubric.
	Methodology string `json:"methodology,omitempty" jsonschema:"description=Source methodology"`
}

// EvaluationType indicates the evaluation approach.
type EvaluationType string

const (
	// EvaluationAnalytic evaluates each category independently.
	EvaluationAnalytic EvaluationType = "analytic"

	// EvaluationHolistic evaluates the document as a whole.
	EvaluationHolistic EvaluationType = "holistic"
)

// Category defines an evaluation dimension.
type Category struct {
	// ID is the category identifier (e.g., "problem_definition").
	ID string `json:"id" jsonschema:"required,description=Category identifier"`

	// Name is the human-readable category name.
	Name string `json:"name" jsonschema:"required,description=Category name"`

	// Description explains what this category evaluates.
	Description string `json:"description,omitempty" jsonschema:"description=Category purpose"`

	// Weight is the category weight (0.0-1.0, all weights should sum to 1.0).
	Weight float64 `json:"weight,omitempty" jsonschema:"minimum=0,maximum=1,description=Category weight"`

	// Required indicates whether this category must pass.
	Required bool `json:"required,omitempty" jsonschema:"description=Whether category is required to pass"`

	// Scale defines the scoring scale for this category.
	Scale *Scale `json:"scale,omitempty" jsonschema:"description=Scoring scale"`
}

// Scale defines the scoring scale for a category.
type Scale struct {
	// Type is the scale type (categorical, integer, numeric).
	Type ScaleType `json:"type" jsonschema:"required,enum=categorical,enum=integer,enum=numeric"`

	// Options are the scale options (for categorical scales).
	Options []ScaleOption `json:"options,omitempty" jsonschema:"description=Scale options"`

	// Min is the minimum value (for numeric scales).
	Min float64 `json:"min,omitempty" jsonschema:"description=Minimum value"`

	// Max is the maximum value (for numeric scales).
	Max float64 `json:"max,omitempty" jsonschema:"description=Maximum value"`
}

// ScaleType indicates the type of scoring scale.
type ScaleType string

const (
	// ScaleCategorical uses discrete categories (pass/partial/fail).
	ScaleCategorical ScaleType = "categorical"

	// ScaleInteger uses integer scores (1-5).
	ScaleInteger ScaleType = "integer"

	// ScaleNumeric uses continuous scores (0.0-1.0).
	ScaleNumeric ScaleType = "numeric"
)

// ScaleOption defines a categorical scale option.
type ScaleOption struct {
	// Value is the option value (e.g., "pass", "partial", "fail").
	Value string `json:"value" jsonschema:"required,description=Option value"`

	// Label is the human-readable label.
	Label string `json:"label" jsonschema:"required,description=Option label"`

	// Criteria describe when to select this option.
	Criteria []string `json:"criteria,omitempty" jsonschema:"description=Selection criteria"`

	// Score is the numeric equivalent (for aggregation).
	Score float64 `json:"score,omitempty" jsonschema:"description=Numeric score equivalent"`
}

// PassCriteria defines what constitutes a passing evaluation.
type PassCriteria struct {
	// MinCategoriesPassing is the minimum categories that must pass.
	// Use "all_required" for all required categories, or a number.
	MinCategoriesPassing string `json:"minCategoriesPassing,omitempty" jsonschema:"description=Minimum passing categories"`

	// MaxFindingsSeverity defines maximum findings by severity.
	MaxFindingsSeverity *SeverityLimits `json:"maxFindingsSeverity,omitempty" jsonschema:"description=Maximum findings by severity"`

	// MinIntScore is the minimum integer score (1-5) to pass.
	MinIntScore int `json:"minIntScore,omitempty" jsonschema:"minimum=1,maximum=5,description=Minimum integer score to pass"`

	// MinNumericScore is the minimum numeric score (0.0-1.0) to pass.
	MinNumericScore float64 `json:"minNumericScore,omitempty" jsonschema:"minimum=0,maximum=1,description=Minimum numeric score to pass"`
}

// SeverityLimits defines maximum findings by severity level.
type SeverityLimits struct {
	// Critical is the maximum critical findings allowed (-1 = unlimited).
	Critical int `json:"critical,omitempty" jsonschema:"description=Max critical findings"`

	// High is the maximum high findings allowed (-1 = unlimited).
	High int `json:"high,omitempty" jsonschema:"description=Max high findings"`

	// Medium is the maximum medium findings allowed (-1 = unlimited).
	Medium int `json:"medium,omitempty" jsonschema:"description=Max medium findings"`

	// Low is the maximum low findings allowed (-1 = unlimited).
	Low int `json:"low,omitempty" jsonschema:"description=Max low findings"`
}
