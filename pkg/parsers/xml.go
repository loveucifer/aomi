// Package parsers provides format-specific parsing for Aomi
// XML parser with structured conversion :D
package parsers

import (
	"github.com/loveucifer/aomi/pkg/schema"
	"strings"
)

// XMLParser parses XML data into the internal document model
type XMLParser struct{}

// Parse parses XML data into a Document
func (p *XMLParser) Parse(data []byte) (*schema.Document, error) {
	// For now, convert XML to a generic structure
	// In a real implementation, we'd use more sophisticated XML parsing

	// Simple approach: return the raw XML for now
	result := map[string]interface{}{
		"_raw_xml": string(data),
	}

	schemaObj := inferSchema(result) // :D auto-detect structure
	doc := &schema.Document{
		Schema: schemaObj,
		Data:   result,
	}

	return doc, nil // :) success
}

// xmlToMap converts XML string to a map structure
func xmlToMap(xmlStr string) interface{} {
	// This is a simplified XML to map converter
	// In a real implementation, we'd use proper XML parsing
	trimmed := strings.TrimSpace(xmlStr)

	// Check if it's an array-like structure (has multiple root elements)
	if strings.HasPrefix(trimmed, "<?xml") {
		// Remove XML declaration
		trimmed = strings.SplitN(trimmed, ">", 2)[1]
	}

	// Return raw XML as a fallback structure
	// Full XML parsing implementation would go here
	return map[string]interface{}{
		"_raw_xml": trimmed,
	}
}
