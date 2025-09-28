// Package writers provides format-specific writing for Aomi
// CSV writer with proper escaping and header support :D
package writers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/loveucifer/aomi/pkg/converters"
	"github.com/loveucifer/aomi/pkg/schema"
)

// CSVWriter writes documents in CSV format
type CSVWriter struct {
	Delimiter rune
	Headers   []string
}

// Write converts a document to CSV bytes
func (w *CSVWriter) Write(doc *schema.Document) ([]byte, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	writer.Comma = w.Delimiter

	// Handle different data types
	switch data := doc.Data.(type) {
	case []interface{}:
		// Array of objects - each object becomes a row
		if len(data) == 0 {
			// No data, just write headers if any
			if len(w.Headers) > 0 {
				writer.Write(w.Headers)
			}
		} else {
			// Get headers from first record if not specified
			headers := w.Headers
			if len(headers) == 0 {
				headers = getCSVHeaders(data[0])
			}

			// Write headers
			writer.Write(headers)

			// Write data rows
			for _, record := range data {
				var row []string
				if recordMap, ok := record.(map[string]interface{}); ok {
					for _, header := range headers {
						value := recordMap[header]
						row = append(row, formatCSVValue(value))
					}
				} else {
					// Flatten complex structures for CSV
					flat := converters.FlattenForCSV(record)
					for _, header := range headers {
						value := flat[header]
						row = append(row, formatCSVValue(value))
					}
				}
				writer.Write(row)
			}
		}
	case map[string]interface{}:
		// Single object - flatten to single row
		headers := w.Headers
		if len(headers) == 0 {
			headers = getMapKeys(data)
		}

		var row []string
		for _, header := range headers {
			value := data[header]
			row = append(row, formatCSVValue(value))
		}

		writer.Write(headers)
		writer.Write(row)
	default:
		// For other types, create a simple key-value mapping
		row := []string{fmt.Sprintf("%v", data)}
		writer.Write([]string{"value"}) // :D single column header
		writer.Write(row)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return nil, err // :0 writing failed
	}

	return buf.Bytes(), nil // :) success
}

// getCSVHeaders extracts headers from a record
func getCSVHeaders(record interface{}) []string {
	var headers []string

	if recordMap, ok := record.(map[string]interface{}); ok {
		for key := range recordMap {
			headers = append(headers, key)
		}
	} else {
		// If it's not a map, flatten it first
		flat := converters.FlattenForCSV(record)
		for key := range flat {
			headers = append(headers, key)
		}
	}

	return headers
}

// getMapKeys gets all keys from a map
func getMapKeys(m map[string]interface{}) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// formatCSVValue formats a value for CSV output
func formatCSVValue(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case float64:
		return fmt.Sprintf("%g", v) // :0 avoid scientific notation
	case bool:
		return fmt.Sprintf("%t", v)
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", v) // :D convert to string
	}
}
