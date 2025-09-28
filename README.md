# Aomi - Universal File Converter

**Repository**: https://github.com/loveucifer/aomi  
**Language**: Go  
**License**: MIT

 Command-line tool that converts between any data formats with intelligent schema detection and zero configuration.

## What Aomi Does

Converts between JSON, CSV, YAML, XML, TOML, and more with automatic format detection and smart field mapping.

```bash
aomi input.json output.csv          # JSON to CSV
aomi data.csv config.yaml           # CSV to YAML  
aomi api.xml data.json             # XML to JSON
aomi --format yaml < input.json    # Pipe with format
```

## Features

- **Auto-detection**: Recognizes input format automatically
- **Smart mapping**: Handles nested structures intelligently (flattens nested objects like address.city to address_city)
- **Batch processing**: Convert multiple files at once
- **Schema inference**: Creates optimal output structure
- **Validation**: Ensures data integrity during conversion
- **Streaming**: Handles large files without memory issues
- **Zero configuration**: Works out of the box

## Installation

```bash
go install github.com/loveucifer/aomi@latest
```

## Usage

### Basic Conversion
```bash
aomi input.json output.csv    # Convert JSON to CSV
aomi data.yaml output.xml     # Convert YAML to XML
```

### Format Specification
```bash
aomi --to csv input.json      # Specify target format
aomi --to yaml data.json      # JSON to YAML
```

### Piped Input
```bash
cat data.json | aomi --to csv        # Pipe JSON to CSV
curl api.json | aomi --to yaml > config.yaml  # API to YAML
```

### Batch Processing
```bash
aomi --batch input/ output/ --to json    # Convert all files
```

### Validation Only
```bash
aomi --validate data.json    # Just detect format
```

### Pretty Output
```bash
aomi --pretty data.json output.json    # Formatted output
```

## Supported Formats

- **JSON** - JavaScript Object Notation
- **CSV** - Comma-Separated Values  
- **YAML** - YAML Ain't Markup Language
- **XML** - eXtensible Markup Language
- **TOML** - Tom's Obvious, Minimal Language

## Examples

### JSON to CSV
```bash
# Input: users.json
{
  "users": [
    {"name": "Alice", "age": 30},
    {"name": "Bob", "age": 25}
  ]
}

# Command
aomi users.json users.csv

# Output: users.csv
name,age
Alice,30
Bob,25
```

### CSV to JSON
```bash
# Input: data.csv
name,age,city
Alice,30,NYC
Bob,25,SF

# Command
aomi data.csv data.json

# Output: data.json
[
  {"name": "Alice", "age": 30, "city": "NYC"},
  {"name": "Bob", "age": 25, "city": "SF"}
]
```

### JSON with Nested Objects to CSV
```bash
# Input: user.json
{
  "name": "John Doe",
  "age": 35,
  "address": {
    "street": "123 Main St",
    "city": "Anytown", 
    "zipcode": "12345"
  },
  "hobbies": ["reading", "swimming", "coding"]
}

# Command
aomi user.json user.csv

# Output: user.csv (nested objects are flattened)
hobbies,address_city,address_zipcode,address_street,name,age
"[reading, swimming, coding]",Anytown,12345,"123 Main St",John Doe,35
```

## Changelog

### v0.1.1
- Fixed CSV writer delimiter issue where uninitialized delimiter was set to null character, causing empty output
- Enhanced CSV writer to properly flatten nested JSON objects for CSV conversion
- Added support for flattening nested structures like {"address": {"city": "NYC"}} to address_city=NYC in CSV output

### v0.1.0
- Initial release with support for JSON, CSV, YAML, XML, and TOML conversion
- Automatic format detection
- Batch processing support

## Architecture

```
Input Data → Format Detection → Parsing → Internal Document → Conversion → Format Writing → Output
```

- **Format Detection**: Automatic format detection using pattern matching
- **Parsers**: Format-specific parsers to convert to internal representation
- **Internal Document**: Universal document model for all formats
- **Converters**: Smart conversion between different structures
- **Writers**: Format-specific writers to generate output

## Development

### Project Structure

```
aomi/
├── cmd/aomi/main.go           # CLI entry point
├── pkg/
│   ├── detector/              # Format detection
│   ├── parsers/               # Input format parsers
│   ├── converters/            # Core conversion logic
│   ├── writers/               # Output format writers
│   └── schema/                # Schema inference
├── examples/                  # Sample files
├── tests/                     # Test files
└── README.md
```

### Building from Source

```bash
git clone https://github.com/loveucifer/aomi.git
cd aomi
go build -o aomi ./cmd/aomi/
# Or to install globally:
go install ./cmd/aomi/
```

### Running Tests

```bash
# Run all package tests
go test ./...

# Run specific tests
go test ./pkg/writers/
go test ./pkg/parsers/
```

### Local Development

```bash
# Build and run directly
go run ./cmd/aomi/main.go input.json output.csv

# Build the binary
go build -o aomi ./cmd/aomi/
./aomi input.json output.csv
```

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to contribute to this project.
