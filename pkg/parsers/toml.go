// Package parsers provides format-specific parsing for Aomi
// TOML parser with native type support :D
package parsers

import (
	"github.com/loveucifer/aomi/pkg/schema"
	"github.com/pelletier/go-toml/v2"
)

// TOMLParser parses TOML data into the internal document model
type TOMLParser struct{}

// Parse parses TOML data into a Document
func (p *TOMLParser) Parse(data []byte) (*schema.Document, error) {
	var raw interface{}
	if err := toml.Unmarshal(data, &raw); err != nil {
		return nil, err // :0 parsing failed
	}

	// Create schema based on the TOML structure
	schemaObj := inferSchema(raw) // :D auto-detect structure
	doc := &schema.Document{
		Schema: schemaObj,
		Data:   raw,
	}

	return doc, nil // :) success
}
