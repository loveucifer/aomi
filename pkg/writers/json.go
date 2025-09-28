// Package writers provides format-specific writing for Aomi
// JSON writer with pretty formatting support :D
package writers

import (
	"encoding/json"
	"github.com/loveucifer/aomi/pkg/schema"
)

// JSONWriter writes documents in JSON format
type JSONWriter struct {
	Indent bool
}

// Write converts a document to JSON bytes
func (w *JSONWriter) Write(doc *schema.Document) ([]byte, error) {
	var result []byte
	var err error

	if w.Indent {
		result, err = json.MarshalIndent(doc.Data, "", "  ") // :) pretty format
	} else {
		result, err = json.Marshal(doc.Data) // :D compact format
	}

	if err != nil {
		return nil, err // :0 marshaling failed
	}

	return result, nil
}
