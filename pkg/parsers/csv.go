// Package parsers provides format-specific parsing for Aomi
// CSV parser with header detection and smart mapping :D
package parsers

import (
	"encoding/csv"
	"github.com/loveucifer/aomi/pkg/schema"
	"strconv"
	"strings"
)

// CSVParser parses CSV data into the internal document model
type CSVParser struct {
	Delimiter rune
	HasHeader bool
}

// NewCSVParser creates a new CSV parser with default settings
func NewCSVParser() *CSVParser {
	return &CSVParser{
		Delimiter: ',',
		HasHeader: true,
	}
}

// Parse parses CSV data into a Document
func (p *CSVParser) Parse(data []byte) (*schema.Document, error) {
	reader := csv.NewReader(strings.NewReader(string(data)))
	reader.Comma = rune(p.Delimiter)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err // :0 parsing failed
	}

	if len(records) == 0 {
		return &schema.Document{
			Schema: &schema.Schema{Type: schema.Array},
			Data:   []interface{}{},
		}, nil
	}

	// Determine headers
	var headers []string
	if p.HasHeader && len(records) > 0 {
		headers = records[0]
		records = records[1:] // Skip header row
	} else {
		// Generate generic headers
		if len(records) > 0 {
			for i := 0; i < len(records[0]); i++ {
				headers = append(headers, "field_"+strconv.Itoa(i))
			}
		}
	}

	// Convert records to array of objects
	var result []interface{}
	for _, record := range records {
		row := make(map[string]interface{})
		for i, value := range record {
			if i < len(headers) {
				// Try to infer data type
				row[headers[i]] = inferType(value) // :D smart type detection
			}
		}
		result = append(result, row)
	}

	// Create schema for the CSV data
	schemaObj := p.inferCSVSchema(headers, result)

	doc := &schema.Document{
		Schema: schemaObj,
		Data:   result,
	}

	return doc, nil // :) success
}

// inferType tries to infer the data type from string value
func inferType(value string) interface{} {
	// Try to parse as boolean
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "true", "1", "yes", "on":
		return true
	case "false", "0", "no", "off":
		return false
	}

	// Try to parse as number
	if num, err := strconv.ParseFloat(value, 64); err == nil {
		return num
	}

	// Default to string
	return value
}

// inferCSVSchema infers schema from CSV headers and sample data
func (p *CSVParser) inferCSVSchema(headers []string, data []interface{}) *schema.Schema {
	// For CSV, we typically have an array of objects
	arraySchema := &schema.Schema{Type: schema.Array}

	if len(data) > 0 {
		// Use the first record to determine schema
		if first, ok := data[0].(map[string]interface{}); ok {
			fields := make(map[string]*schema.FieldSchema)

			for _, header := range headers {
				if value, exists := first[header]; exists {
					dataType := inferSchema(value).Type
					fields[header] = &schema.FieldSchema{
						Name:     header,
						Type:     dataType,
						Required: true,
						Nested:   inferSchema(value),
					}
				}
			}

			arraySchema.Items = &schema.FieldSchema{
				Name:   "record",
				Type:   schema.Object,
				Nested: &schema.Schema{Type: schema.Object, Fields: fields},
			}
		}
	}

	return arraySchema
}
