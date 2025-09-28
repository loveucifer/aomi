// Package parsers provides format-specific parsing for Aomi
// JSON parser with automatic schema detection :)
package parsers

import (
	"encoding/json"
	"github.com/loveucifer/aomi/pkg/schema"
)

// JSONParser parses JSON data into the internal document model
type JSONParser struct{}

// Parse parses JSON data into a Document
func (p *JSONParser) Parse(data []byte) (*schema.Document, error) {
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err // :0 parsing failed
	}

	// Create schema based on the JSON structure
	schemaObj := inferSchema(raw) // :D auto-detect structure
	doc := &schema.Document{
		Schema: schemaObj,
		Data:   raw,
	}

	return doc, nil // :) success
}

// inferSchema infers the schema from the raw data
func inferSchema(data interface{}) *schema.Schema {
	switch v := data.(type) {
	case string:
		return &schema.Schema{Type: schema.String}
	case float64: // JSON numbers are float64
		return &schema.Schema{Type: schema.Number}
	case bool:
		return &schema.Schema{Type: schema.Boolean}
	case []interface{}:
		// Array - get schema of first element as representative
		s := &schema.Schema{Type: schema.Array}
		if len(v) > 0 {
			s.Items = &schema.FieldSchema{
				Name:   "item",
				Type:   inferSchema(v[0]).Type,
				Nested: inferSchema(v[0]),
			}
		}
		return s
	case map[string]interface{}:
		// Object - create field schema for each key
		fields := make(map[string]*schema.FieldSchema)
		for key, value := range v {
			fields[key] = &schema.FieldSchema{
				Name:     key,
				Type:     inferSchema(value).Type,
				Required: true,
				Nested:   inferSchema(value),
			}
		}
		return &schema.Schema{
			Type:   schema.Object,
			Fields: fields,
		}
	default:
		return &schema.Schema{Type: schema.String} // Default to string
	}
}
