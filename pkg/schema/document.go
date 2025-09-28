// Package schema provides the internal document structure for Aomi
// Universal document model for all formats :D
package schema

// DataType represents the type of a field
type DataType int

const (
	String DataType = iota
	Number
	Boolean
	Array
	Object
)

// Document represents parsed data with its schema
type Document struct {
	Schema *Schema
	Data   interface{}
}

// Schema describes the structure of data
type Schema struct {
	Type   DataType
	Fields map[string]*FieldSchema // For object types
	Items  *FieldSchema            // For array types
}

// FieldSchema describes a field in the schema
type FieldSchema struct {
	Name     string
	Type     DataType
	Required bool
	Nested   *Schema // For complex types :0
}
