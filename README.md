# Specification Workflow Spec

A formal specification for defining product specification workflows.

## Overview

`specification-workflow-spec` provides standardized types for defining:

- **Spec Types** - Registry of specification document types (PRD, MRD, Press Release, FAQ, 6-Pager, etc.)
- **Profiles** - Workflow configurations bundling spec requirements, synthesis rules, and evaluation criteria
- **Templates** - Document structure definitions with required/optional sections
- **Rubric Definitions** - LLM-as-Judge evaluation criteria (distinct from evaluation reports)
- **Synthesis Rules** - Dependency graphs for generating specs from other specs
- **Phase Gates** - Approval checkpoints and workflow control

## Architecture

```
┌────────────────────────────────────────────────────────────────────────────┐
│                         Profile (Workflow Configuration)                   │
├────────────────────────────────────────────────────────────────────────────┤
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐   │
│  │  SpecConfig  │  │  Synthesis   │  │  Templates   │  │   Rubrics    │   │
│  │  (required/  │  │  (DAG of     │  │  (document   │  │  (evaluation │   │
│  │   optional)  │  │   sources)   │  │   structure) │  │   criteria)  │   │
│  └──────────────┘  └──────────────┘  └──────────────┘  └──────────────┘   │
└────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌────────────────────────────────────────────────────────────────────────────┐
│                           Workflow Execution                               │
├────────────────────────────────────────────────────────────────────────────┤
│  Phase 1: Discovery  →  Gate  →  Phase 2: Vision  →  Gate  →  Phase 3...  │
│  (MRD)                          (Press, FAQ)                  (PRD, UXD)   │
└────────────────────────────────────────────────────────────────────────────┘
```

## Installation

```bash
go get github.com/ProductBuildersHQ/specification-workflow-spec
```

## Packages

| Package | Description |
|---------|-------------|
| `pkg/spectype` | Spec type registry and category definitions |
| `pkg/profile` | Workflow profile configuration |
| `pkg/template` | Spec template structure definitions |
| `pkg/rubricdef` | Rubric definition for LLM-as-Judge evaluation |
| `pkg/synthesis` | Synthesis rule DAG for spec generation |
| `pkg/gate` | Phase gates and approval checkpoints |
| `schema` | Generated JSON Schema files |

## Spec Types

The registry defines canonical spec types across methodologies:

| ID | Name | Category | Origins |
|----|------|----------|---------|
| `prd` | Product Requirements Document | source | startup, enterprise, big-tech |
| `mrd` | Market Requirements Document | source | enterprise, aws-product |
| `press` | Press Release | gtm | aws-product, big-tech |
| `faq` | Frequently Asked Questions | gtm | aws-product, big-tech |
| `narrative-6p` | Six-Pager Narrative | gtm | aws-product, big-tech |
| `trd` | Technical Requirements Document | technical | enterprise, google |
| `plan` | Implementation Plan | execution | pbhq-lite |
| `roadmap` | Roadmap | execution | pbhq-lite |

See `pkg/spectype/spectype.go` for the full registry.

## Profiles

Profiles bundle workflow configurations for specific methodologies:

```go
import "github.com/ProductBuildersHQ/specification-workflow-spec/pkg/profile"

p := profile.Profile{
    Name:        "aws-product",
    Description: "Amazon Working Backwards for new products",
    SpecConfig: map[string]*profile.SpecRequirement{
        "mrd":   {Required: true},
        "press": {Required: true},
        "faq":   {Required: true},
        "prd":   {Required: true},
    },
    Synthesis: map[string]*synthesis.Rule{
        "press": {Sources: []string{"mrd"}},
        "faq":   {Sources: []string{"mrd", "press"}},
        "prd":   {Sources: []string{"mrd", "press", "faq"}},
    },
}
```

## Schema Generation

JSON Schema files are generated from Go types:

```bash
go generate ./schema/...
```

This produces:

- `schema/spectype.schema.json`
- `schema/profile.schema.json`
- `schema/template.schema.json`
- `schema/rubricdef.schema.json`
- `schema/synthesis.schema.json`
- `schema/gate.schema.json`

## Related Projects

| Project | Purpose |
|---------|---------|
| [visionspec](https://github.com/ProductBuildersHQ/visionspec) | CLI for spec workflow execution |
| [structured-evaluation](https://github.com/plexusone/structured-evaluation) | LLM-as-Judge evaluation reports |
| [multi-agent-spec](https://github.com/plexusone/multi-agent-spec) | Multi-agent system definitions |

## License

MIT
