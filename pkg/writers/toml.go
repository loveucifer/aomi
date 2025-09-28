// Package writers provides format-specific writing for Aomi
// TOML writer with native type preservation :D
package writers

import (
	"github.com/loveucifer/aomi/pkg/schema"
	"github.com/pelletier/go-toml/v2"
)

// TOMLWriter writes documents in TOML format
type TOMLWriter struct{}

// Write converts a document to TOML bytes
func (w *TOMLWriter) Write(doc *schema.Document) ([]byte, error) {
	result, err := toml.Marshal(doc.Data) // :D preserve native types
	if err != nil {
		return nil, err // :0 marshaling failed
	}

	return result, nil // :) success
}
