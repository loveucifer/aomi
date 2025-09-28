// Package writers provides format-specific writing for Aomi
// XML writer with proper structure :D
package writers

import (
	"fmt"
	"github.com/loveucifer/aomi/pkg/schema"
	"strings"
)

// XMLWriter writes documents in XML format
type XMLWriter struct {
	RootTag string
}

// Write converts a document to XML bytes
func (w *XMLWriter) Write(doc *schema.Document) ([]byte, error) {
	rootTag := w.RootTag
	if rootTag == "" {
		rootTag = "root" // :D default root element
	}

	var result strings.Builder
	result.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
	result.WriteString(fmt.Sprintf("<%s>\n", rootTag))

	// Convert the document data to XML
	xmlContent := convertToXML(doc.Data, "  ") // :D indented XML
	result.WriteString(xmlContent)

	result.WriteString(fmt.Sprintf("</%s>\n", rootTag))

	return []byte(result.String()), nil // :) success
}

// convertToXML converts data to XML string representation
func convertToXML(data interface{}, indent string) string {
	var result strings.Builder

	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			result.WriteString(indent)
			result.WriteString(fmt.Sprintf("<%s>", key))

			switch val := value.(type) {
			case map[string]interface{}:
				result.WriteString("\n")
				result.WriteString(convertToXML(val, indent+"  ")) // :D nested indent
				result.WriteString(indent)
				result.WriteString(fmt.Sprintf("</%s>\n", key))
			case []interface{}:
				result.WriteString("\n")
				for _, item := range val {
					result.WriteString(indent + "  ")
					result.WriteString(fmt.Sprintf("<%s_item>", key)) // :D array items
					result.WriteString(fmt.Sprintf("%v", item))
					result.WriteString(fmt.Sprintf("</%s_item>\n", key))
				}
				result.WriteString(indent)
				result.WriteString(fmt.Sprintf("</%s>\n", key))
			default:
				result.WriteString(fmt.Sprintf("%v", val))
				result.WriteString(fmt.Sprintf("</%s>\n", key))
			}
		}
	case []interface{}:
		for i, item := range v {
			itemTag := fmt.Sprintf("item_%d", i)
			result.WriteString(indent)
			result.WriteString(fmt.Sprintf("<%s>", itemTag))

			switch val := item.(type) {
			case map[string]interface{}:
				result.WriteString("\n")
				result.WriteString(convertToXML(val, indent+"  "))
				result.WriteString(indent)
				result.WriteString(fmt.Sprintf("</%s>\n", itemTag))
			default:
				result.WriteString(fmt.Sprintf("%v", val))
				result.WriteString(fmt.Sprintf("</%s>\n", itemTag))
			}
		}
	default:
		result.WriteString(indent)
		result.WriteString(fmt.Sprintf("%v\n", v)) // :D simple value
	}

	return result.String()
}
