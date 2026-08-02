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
┌───────────────────────────────────────────────────────────────────────────┐
│                           Workflow Execution                              │
├───────────────────────────────────────────────────────────────────────────┤
│  Phase 1: Discovery  →  Gate  →  Phase 2: Vision  →  Gate  →  Phase 3...  │
│  (MRD)                          (Press, FAQ)                  (PRD, UXD)  │
└───────────────────────────────────────────────────────────────────────────┘
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
| `pkg/diagram` | D2 and Mermaid diagram generation from profiles |
| `schema` | Generated JSON Schema files |

## Spec Types

The registry defines canonical spec types across methodologies:

### Source Specs (Human-Authored)

| ID | Name | Category | Origins |
|----|------|----------|---------|
| `mrd` | Market Requirements Document | source | enterprise, aws-product, big-tech-product |
| `prd` | Product Requirements Document | source | startup, enterprise, big-tech |
| `uxd` | User Experience Design | source | design-thinking, big-tech |
| `opportunity-spec` | Opportunity Specification | source | aws-feature, big-tech-feature |
| `hypothesis` | Hypothesis Document | source | lean-startup, 0-1 |
| `shapeup-pitch` | Shape Up Pitch | source | shapeup |
| `ost` | Opportunity Solution Tree | source | continuous-discovery |

### GTM Specs (Synthesized)

| ID | Name | Category | Origins |
|----|------|----------|---------|
| `press` | Press Release | gtm | aws-product, big-tech |
| `faq` | Frequently Asked Questions | gtm | aws-product, big-tech |
| `narrative-6p` | Six-Pager Narrative | gtm | aws-product, big-tech-product |
| `narrative-1p` | One-Pager Executive Summary | gtm | enterprise, big-tech |
| `bmc` | Business Model Canvas | gtm | enterprise, lean-startup |

### Technical Specs (Synthesized)

| ID | Name | Category | Origins |
|----|------|----------|---------|
| `trd` | Technical Requirements Document | technical | enterprise, google, big-tech |
| `tpd` | Test Plan Document | technical | enterprise, big-tech |
| `ird` | Infrastructure Requirements Document | technical | enterprise, big-tech-product |

### Execution Specs

| ID | Name | Category | Origins |
|----|------|----------|---------|
| `plan` | Implementation Plan | execution | pbhq-lite |
| `roadmap` | Roadmap | execution | pbhq-lite |
| `spec` | Reconciled Specification | output | enterprise |

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

## Diagram Generation

Generate D2 or Mermaid diagrams from profiles:

```go
import "github.com/ProductBuildersHQ/specification-workflow-spec/pkg/diagram"

// Generate D2 diagram
opts := diagram.DefaultOptions()
opts.Title = "AWS Product Flow"
d2, _ := diagram.Generate(profile, diagram.FormatD2, opts)

// Generate Mermaid diagram
mermaid, _ := diagram.Generate(profile, diagram.FormatMermaid, opts)
```

Output formats:

- **D2** - [D2 language](https://d2lang.com/) for SVG generation via `d2` CLI
- **Mermaid** - [Mermaid](https://mermaid.js.org/) for embedding in Markdown

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
