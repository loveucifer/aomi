// Package parsers provides format-specific parsing for Aomi
// YAML parser preserving structure and types :0
package parsers

import (
	"github.com/loveucifer/aomi/pkg/schema"
	"gopkg.in/yaml.v3"
)

// YAMLParser parses YAML data into the internal document model
type YAMLParser struct{}

// Parse parses YAML data into a Document
func (p *YAMLParser) Parse(data []byte) (*schema.Document, error) {
	var raw interface{}
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return nil, err // :0 parsing failed
	}

	// Create schema based on the YAML structure
	schemaObj := inferSchema(raw) // :D auto-detect structure
	doc := &schema.Document{
		Schema: schemaObj,
		Data:   raw,
	}

	return doc, nil // :) success
}
