// Package detector provides format detection capabilities for Aomi
// Smart format detection with pattern matching :D
package detector

import (
	"strings"
	"unicode"
)

// Format represents a supported data format
type Format int

const (
	JSON Format = iota
	CSV
	YAML
	XML
	TOML
	Unknown
)

// String returns the string representation of a format
func (f Format) String() string {
	switch f {
	case JSON:
		return "json"
	case CSV:
		return "csv"
	case YAML:
		return "yaml"
	case XML:
		return "xml"
	case TOML:
		return "toml"
	default:
		return "unknown"
	}
}

// Detector identifies the format of input data
type Detector struct {
	matchers []FormatMatcher
}

// FormatMatcher defines a function that detects a specific format
type FormatMatcher func([]byte) bool

// NewDetector creates a new format detector
func NewDetector() *Detector {
	return &Detector{
		matchers: []FormatMatcher{
			isJSON,
			isCSV,
			isYAML,
			isXML,
			isTOML,
		},
	}
}

// DetectFormat detects the format of the input data
func (d *Detector) DetectFormat(data []byte) Format {
	for i, matcher := range d.matchers {
		if matcher(data) {
			return Format(i)
		}
	}
	return Unknown
}

// isJSON checks if the data is in JSON format
func isJSON(data []byte) bool {
	s := string(data)
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return false
	}

	// Check if it starts with { or [ and ends with } or ]
	return (s[0] == '{' && s[len(s)-1] == '}') ||
		(s[0] == '[' && s[len(s)-1] == ']')
}

// isCSV checks if the data is in CSV format
func isCSV(data []byte) bool {
	s := string(data)
	lines := strings.Split(s, "\n")

	if len(lines) == 0 {
		return false
	}

	// Look for comma-separated values pattern
	for _, line := range lines {
		if strings.Contains(line, ",") && len(strings.TrimSpace(line)) > 0 {
			// Check if it looks like CSV (comparable fields)
			fields := strings.Split(line, ",")
			if len(fields) >= 2 { // At least 2 fields to be considered CSV
				return true
			}
		}
	}
	return false
}

// isYAML checks if the data is in YAML format
func isYAML(data []byte) bool {
	s := string(data)
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 && !strings.HasPrefix(line, "#") { // Skip comments
			// Look for YAML key: value pattern
			if strings.Contains(line, ": ") || strings.HasPrefix(line, "- ") {
				return true
			}
			// Look for indentation in YAML
			if strings.HasPrefix(line, "  ") || strings.HasPrefix(line, "\t") {
				return true
			}
		}
	}
	return false
}

// isXML checks if the data is in XML format
func isXML(data []byte) bool {
	s := string(data)
	s = strings.TrimSpace(s)

	if len(s) < 5 {
		return false
	}

	// Check if it starts with XML declaration or tag
	return strings.HasPrefix(s, "<?xml") ||
		strings.HasPrefix(s, "<") ||
		strings.Contains(s, "</")
}

// isTOML checks if the data is in TOML format
func isTOML(data []byte) bool {
	s := string(data)
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		// Skip empty lines and comments
		if len(line) > 0 && !strings.HasPrefix(line, "#") {
			// Look for TOML key = value pattern or table headers
			if strings.Contains(line, "=") && !strings.Contains(line, ":") {
				parts := strings.SplitN(line, "=", 2)
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					// Valid TOML key should not contain special characters
					isValidKey := true
					for _, r := range key {
						if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' && r != '-' {
							isValidKey = false
							break
						}
					}
					if isValidKey {
						return true
					}
				}
			}
			// Check for TOML table headers [table]
			if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
				return true
			}
		}
	}
	return false
}
// Final implementation
