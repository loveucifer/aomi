// Package converters provides cross-format conversion for Aomi
// Core conversion engine with smart mapping :D
package converters

import (
	"fmt"
	"github.com/loveucifer/aomi/pkg/detector"
	"github.com/loveucifer/aomi/pkg/schema"
)

// Converter handles conversion between different formats
type Converter struct {
	sourceFormat detector.Format
	targetFormat detector.Format
}

// NewConverter creates a new converter instance
func NewConverter(source, target detector.Format) *Converter {
	return &Converter{
		sourceFormat: source,
		targetFormat: target,
	}
}

// Convert converts a document from one format to another
func (c *Converter) Convert(doc *schema.Document, target detector.Format) (*schema.Document, error) {
	// In this basic implementation, we'll return the same document
	// In a full implementation, we'd transform the structure as needed
	return doc, nil // :) for now, just pass through
}

// FlattenForCSV flattens nested structures for CSV output
func FlattenForCSV(data interface{}) map[string]interface{} {
	// Flatten nested structures for CSV output
	// user.name -> user_name
	// address.street -> address_street :0
	flat := make(map[string]interface{})

	if m, ok := data.(map[string]interface{}); ok {
		for k, v := range m {
			switch val := v.(type) {
			case map[string]interface{}:
				// Flatten nested objects
				for nestedK, nestedV := range val {
					flat[k+"_"+nestedK] = nestedV
				}
			case []interface{}:
				// Convert arrays to strings for CSV
				flat[k] = arrayToString(val)
			default:
				flat[k] = val
			}
		}
	}

	return flat
}

// arrayToString converts an array to a string representation
func arrayToString(arr []interface{}) string {
	var result []string
	for _, item := range arr {
		result = append(result, stringOrNumberToString(item))
	}
	return "[" + stringSliceToString(result, ", ") + "]" // :D array as string
}

// stringOrNumberToString converts a value to string
func stringOrNumberToString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case float64:
		return fmt.Sprintf("%g", val)
	case bool:
		return fmt.Sprintf("%t", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

// stringSliceToString joins a string slice with a separator
func stringSliceToString(slice []string, sep string) string {
	if len(slice) == 0 {
		return ""
	}

	result := slice[0]
	for _, s := range slice[1:] {
		result += sep + s
	}
	return result
}
