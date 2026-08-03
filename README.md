# Specification Workflow Spec

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Docs][docs-mkdoc-svg]][docs-mkdoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/ProductBuildersHQ/specification-workflow-spec/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/ProductBuildersHQ/specification-workflow-spec/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/ProductBuildersHQ/specification-workflow-spec/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/ProductBuildersHQ/specification-workflow-spec/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/ProductBuildersHQ/specification-workflow-spec/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/ProductBuildersHQ/specification-workflow-spec/actions/workflows/go-sast-codeql.yaml
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/ProductBuildersHQ/specification-workflow-spec
 [docs-godoc-url]: https://pkg.go.dev/github.com/ProductBuildersHQ/specification-workflow-spec
 [docs-mkdoc-svg]: https://img.shields.io/badge/Go-dev%20guide-blue.svg
 [docs-mkdoc-url]: https://productbuildershq.com/specification-workflow-spec
 [viz-svg]: https://img.shields.io/badge/Go-visualizaton-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=ProductBuildersHQ%2Fspecification-workflow-spec
 [loc-svg]: https://tokei.rs/b1/github/ProductBuildersHQ/specification-workflow-spec
 [repo-url]: https://github.com/ProductBuildersHQ/specification-workflow-spec
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/ProductBuildersHQ/specification-workflow-spec/blob/main/LICENSE

A formal specification for defining product specification workflows.

## Overview

`specification-workflow-spec` provides standardized types for defining:

- **Spec Types** - Registry of specification document types (PRD, MRD, Press Release, FAQ, 6-Pager, etc.)
- **Workflows** - Methodology configurations bundling spec requirements, synthesis rules, and evaluation criteria
- **Templates** - Document structure definitions with required/optional sections and embedded content
- **Rubrics** - LLM-as-Judge evaluation criteria using structured-evaluation's `rubric.RubricSet`
- **Synthesis Rules** - Dependency graphs for generating specs from other specs
- **Phase Gates** - Approval checkpoints and workflow control

## Architecture

```
┌───────────────────────────────────────────────────────────────────────────┐
│                       Workflow (Methodology Configuration)                │
├───────────────────────────────────────────────────────────────────────────┤
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐   │
│  │  SpecConfig  │  │  Synthesis   │  │  Templates   │  │   Rubrics    │   │
│  │  (required/  │  │  (DAG of     │  │  (document   │  │  (evaluation │   │
│  │   optional)  │  │   sources)   │  │   structure) │  │   criteria)  │   │
│  └──────────────┘  └──────────────┘  └──────────────┘  └──────────────┘   │
└───────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌───────────────────────────────────────────────────────────────────────────┐
│                                Execution                                  │
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
| `pkg/workflow` | Workflow configuration (spec requirements, synthesis, execution, evaluation) |
| `pkg/workflows` | Embedded default workflows with loaders (embedded, file, chain, resolving) |
| `pkg/template` | Spec template structure definitions |
| `pkg/synthesis` | Synthesis rule DAG for spec generation |
| `pkg/gate` | Phase gates and approval checkpoints |
| `pkg/layout` | Filesystem layout conventions for spec projects |
| `pkg/diagram` | D2 and Mermaid diagram generation from workflows |
| `schema` | Generated JSON Schema files |

Rubric definitions use [structured-evaluation](https://github.com/plexusone/structured-evaluation)'s canonical `rubric.RubricSet` type.

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

## Workflows

Workflows bundle spec requirements, synthesis rules, templates, and rubrics for
specific methodologies. Default workflows (aws-product, big-tech-feature,
lean-startup, etc.) are embedded and load with no filesystem access:

```go
import "github.com/ProductBuildersHQ/specification-workflow-spec/pkg/workflows"

// Load with inheritance resolution (aws-feature extends enterprise)
w, err := workflows.DefaultLoader().Load("aws-feature")
if err != nil {
    // handle error
}

w.Workflow.Name              // "aws-feature"
w.Workflow.RequiredSpecs()   // required spec type IDs
w.Templates["press"].Content // raw markdown template
w.Rubrics["press"].Categories // structured-evaluation rubric categories
```

Loaders compose for customization:

```go
// Organization overrides from a directory, falling back to embedded defaults
loader := workflows.NewResolvingLoader(workflows.NewChainLoader(
    workflows.NewFileLoader("./custom-workflows"),
    workflows.DefaultLoader(),
))
```

## Diagram Generation

Generate D2 or Mermaid diagrams from workflows:

```go
import "github.com/ProductBuildersHQ/specification-workflow-spec/pkg/diagram"

// Generate D2 diagram
opts := diagram.DefaultOptions()
opts.Title = "AWS Product Flow"
d2, _ := diagram.Generate(w.Workflow, diagram.FormatD2, opts)

// Generate Mermaid diagram
mermaid, _ := diagram.Generate(w.Workflow, diagram.FormatMermaid, opts)
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
- `schema/workflow.schema.json`
- `schema/template.schema.json`
- `schema/synthesis.schema.json`
- `schema/gate.schema.json`
- `schema/layout.schema.json`

## Ecosystem

This repository is the contract layer of the ProductBuildersHQ spec stack
(`visionstudio → visionspec → specification-workflow-spec`): it defines
workflow types, schemas, and the embedded default workflow library, and holds
no execution logic. The layers above act on it.

| Project | Role |
|---------|------|
| [visionspec](https://github.com/ProductBuildersHQ/visionspec) | The engine: CLI, MCP server, and importable SDK executing these workflows (scaffolding, LLM synthesis, LLM-as-Judge evaluation, lint/drift/status) |
| [visionstudio](https://github.com/ProductBuildersHQ/visionstudio) | The studio: LLM-powered app loading workflow data from this library directly and executing via visionspec |
| [structured-evaluation](https://github.com/plexusone/structured-evaluation) | Canonical rubric and evaluation-report types; rubrics here are its `rubric.RubricSet` |
| [multi-agent-spec](https://github.com/plexusone/multi-agent-spec) | Multi-agent system definitions |

## License

MIT
