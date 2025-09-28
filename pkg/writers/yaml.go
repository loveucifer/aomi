// Package writers provides format-specific writing for Aomi
// YAML writer with clean formatting :D
package writers

import (
	"github.com/loveucifer/aomi/pkg/schema"
	"gopkg.in/yaml.v3"
)

// YAMLWriter writes documents in YAML format
type YAMLWriter struct{}

// Write converts a document to YAML bytes
func (w *YAMLWriter) Write(doc *schema.Document) ([]byte, error) {
	result, err := yaml.Marshal(doc.Data) // :D clean YAML output
	if err != nil {
		return nil, err // :0 marshaling failed
	}

	return result, nil // :) success
}
