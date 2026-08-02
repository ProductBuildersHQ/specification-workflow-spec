// Package schema provides embedded JSON Schema files generated from Go types.
//
//go:generate go run generate.go
package schema

import "embed"

// FS embeds the generated JSON Schema files.
//
//go:embed *.schema.json
var FS embed.FS

// SchemaFiles lists the available schema files.
var SchemaFiles = []string{
	"spectype.schema.json",
	"profile.schema.json",
	"template.schema.json",
	"rubricdef.schema.json",
	"synthesis.schema.json",
	"gate.schema.json",
}
