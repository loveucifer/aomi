package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content := `// Package main implements the Aomi universal file converter
// Your gateway to format conversion magic! :)
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/loveucifer/aomi/pkg/detector"
	"github.com/loveucifer/aomi/pkg/parsers"
	"github.com/loveucifer/aomi/pkg/schema"
	"github.com/loveucifer/aomi/pkg/writers"
)

var (
	toFormat   = flag.String("to", "", "Target format (json, csv, yaml, xml, toml)")
	pretty     = flag.Bool("pretty", false, "Pretty print output")
	batch      = flag.Bool("batch", false, "Batch process directory")
	validate   = flag.Bool("validate", false, "Validate input format only")
	help       = flag.Bool("help", false, "Show help message")
	version    = flag.Bool("version", false, "Show version information")
)

const versionString = "Aomi v0.1.0 - Universal File Converter"

func main() {
	flag.Parse()

	if *help {
		printUsage()
		os.Exit(0)
	}

	if *version {
		fmt.Println(versionString)
		os.Exit(0)
	}

	args := flag.Args()
	
	if len(args) == 0 && !isPipeInput() {
		fmt.Println("Error: No input files specified")
		printUsage()
		os.Exit(1)
	}

	if *validate {
		// Validation mode: just detect format and report
		err := validateInput(args)
		if err != nil {
			fmt.Printf("Validation error: %%v\n", err)
			os.Exit(1)
		}
		return
	}

	if *batch {
		// Batch mode: process directory
		if len(args) < 2 {
			fmt.Println("Batch mode requires input and output directories")
			os.Exit(1)
		}
		err := processBatch(args[0], args[1], *toFormat)
		if err != nil {
			fmt.Printf("Batch processing error: %%v\n", err)
			os.Exit(1)
		}
		return
	}

	// Single file or piped input mode
	if isPipeInput() {
		err := processPipedInput(*toFormat, *pretty)
		if err != nil {
			fmt.Printf("Error processing input: %%v\n", err)
			os.Exit(1)
		}
	} else {
		// File-to-file conversion
		if len(args) < 2 {
			fmt.Println("Error: Input and output files required")
			printUsage()
			os.Exit(1)
		}
		err := processFile(args[0], args[1], *toFormat, *pretty)
		if err != nil {
			fmt.Printf("Error converting files: %%v\n", err)
			os.Exit(1)
		}
	}
}

// isPipeInput checks if input is being piped
func isPipeInput() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

// validateInput validates the input format
func validateInput(inputs []string) error {
	detectorInst := detector.NewDetector()
	
	for _, input := range inputs {
		data, err := ioutil.ReadFile(input)
		if err != nil {
			return fmt.Errorf("reading %%s: %%v", input, err)
		}
		
		format := detectorInst.DetectFormat(data)
		fmt.Printf("%%s: %%s\n", input, format.String()) // :D detected format
	}
	
	return nil
}

// processPipedInput handles piped input
func processPipedInput(targetFormat string, pretty bool) error {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("reading stdin: %%v", err)
	}
	
	detectorInst := detector.NewDetector()
	sourceFormat := detectorInst.DetectFormat(data)
	
	if sourceFormat == detector.Unknown {
		return fmt.Errorf("unknown input format")
	}
	
	// Parse the input
	doc, err := parseData(data, sourceFormat)
	if err != nil {
		return fmt.Errorf("parsing input: %%v", err)
	}
	
	// Convert to target format
	target := detector.Unknown
	if targetFormat != "" {
		target = stringToFormat(targetFormat)
	} else {
		// If not specified, default to JSON for piped output
		target = detector.JSON
	}
	
	if target == detector.Unknown {
		return fmt.Errorf("unknown target format: %%s", targetFormat)
	}
	
	// Write the output
	output, err := writeData(doc, target, pretty)
	if err != nil {
		return fmt.Errorf("writing output: %%v", err)
	}
	
	fmt.Print(string(output))
	return nil
}

// processFile converts a single file
func processFile(inputFile, outputFile, targetFormat string, pretty bool) error {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("reading %%s: %%v", inputFile, err)
	}
	
	detectorInst := detector.NewDetector()
	sourceFormat := detectorInst.DetectFormat(data)
	
	if sourceFormat == detector.Unknown {
		return fmt.Errorf("unknown input format for %%s", inputFile)
	}
	
	// Determine target format
	target := sourceFormat
	if targetFormat != "" {
		target = stringToFormat(targetFormat)
	} else {
		// Infer from output file extension
		ext := strings.TrimPrefix(filepath.Ext(outputFile), ".")
		target = stringToFormat(ext)
	}
	
	if target == detector.Unknown {
		return fmt.Errorf("unknown target format: %%s", targetFormat)
	}
	
	// Parse the input
	doc, err := parseData(data, sourceFormat)
	if err != nil {
		return fmt.Errorf("parsing input: %%v", err)
	}
	
	// Write the output
	output, err := writeData(doc, target, pretty)
	if err != nil {
		return fmt.Errorf("writing output: %%v", err)
	}
	
	err = ioutil.WriteFile(outputFile, output, 0644)
	if err != nil {
		return fmt.Errorf("writing %%s: %%v", outputFile, err)
	}
	
	fmt.Printf("Converted %%s (%%s) -> %%s (%%s)\n", inputFile, sourceFormat.String(), outputFile, target.String()) // :D conversion complete
	return nil
}

// processBatch processes all files in a directory
func processBatch(inputDir, outputDir, targetFormat string) error {
	// Create output directory if it doesn't exist
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		return fmt.Errorf("creating output directory: %%v", err)
	}
	
	// Read all files in input directory
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		return fmt.Errorf("reading directory: %%v", err)
	}
	
	detectorInst := detector.NewDetector()
	
	for _, file := range files {
		if file.IsDir() {
			continue // Skip subdirectories for now
		}
		
		inputPath := filepath.Join(inputDir, file.Name())
		ext := strings.TrimPrefix(filepath.Ext(file.Name()), ".")
		sourceFormat := stringToFormat(ext)
		
		// If extension doesn't match a known format, try detection
		if sourceFormat == detector.Unknown {
			data, err := ioutil.ReadFile(inputPath)
			if err != nil {
				fmt.Printf("Warning: could not read %%s: %%v\n", inputPath, err)
				continue
			}
			sourceFormat = detectorInst.DetectFormat(data)
		}
		
		if sourceFormat == detector.Unknown {
			fmt.Printf("Warning: unknown format for %%s, skipping\n", inputPath)
			continue
		}
		
		// Determine output file path and format
		outputExt := ext
		if targetFormat != "" {
			outputExt = targetFormat
		}
		
		// Change extension to target format
		baseName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		outputPath := filepath.Join(outputDir, baseName+"."+outputExt)
		
		target := stringToFormat(outputExt)
		if target == detector.Unknown {
			fmt.Printf("Warning: unknown target format %%s, skipping %%s\n", outputExt, inputPath)
			continue
		}
		
		// Process the file
		err = processFile(inputPath, outputPath, outputExt, false)
		if err != nil {
			fmt.Printf("Warning: error processing %%s: %%v\n", inputPath, err)
			continue
		}
	}
	
	return nil
}

// parseData parses data based on its format
func parseData(data []byte, format detector.Format) (*schema.Document, error) {
	switch format {
	case detector.JSON:
		return (&parsers.JSONParser{}).Parse(data)
	case detector.CSV:
		return (&parsers.CSVParser{}).Parse(data)
	case detector.YAML:
		return (&parsers.YAMLParser{}).Parse(data)
	case detector.XML:
		return (&parsers.XMLParser{}).Parse(data)
	case detector.TOML:
		return (&parsers.TOMLParser{}).Parse(data)
	default:
		return nil, fmt.Errorf("unsupported format: %%s", format.String())
	}
}

// writeData writes document data in the specified format
func writeData(doc *schema.Document, format detector.Format, pretty bool) ([]byte, error) {
	switch format {
	case detector.JSON:
		writer := &writers.JSONWriter{Indent: pretty}
		return writer.Write(doc)
	case detector.CSV:
		writer := &writers.CSVWriter{}
		return writer.Write(doc)
	case detector.YAML:
		writer := &writers.YAMLWriter{}
		return writer.Write(doc)
	case detector.XML:
		writer := &writers.XMLWriter{}
		return writer.Write(doc)
	case detector.TOML:
		writer := &writers.TOMLWriter{}
		return writer.Write(doc)
	default:
		return nil, fmt.Errorf("unsupported format: %%s", format.String())
	}
}

// stringToFormat converts a string to a Format
func stringToFormat(s string) detector.Format {
	switch strings.ToLower(s) {
	case "json":
		return detector.JSON
	case "csv":
		return detector.CSV
	case "yaml", "yml":
		return detector.YAML
	case "xml":
		return detector.XML
	case "toml":
		return detector.TOML
	default:
		return detector.Unknown
	}
}

// printUsage shows the help message
func printUsage() {
	fmt.Println(versionString)
	fmt.Println("Converts between JSON, CSV, YAML, XML, TOML and more formats")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  aomi [options] input output        # Convert input file to output file")
	fmt.Println("  aomi --to format < input > output  # Pipe with format")
	fmt.Println("  aomi --batch input_dir output_dir  # Batch convert directory")
	fmt.Println()
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  aomi data.json output.csv          # JSON to CSV")
	fmt.Println("  aomi --to yaml data.json           # JSON to YAML")
	fmt.Println("  cat data.csv | aomi --to json      # Pipe conversion")
	fmt.Println("  aomi --batch ./in ./out --to json  # Batch conversion")
}
`

	err := ioutil.WriteFile("cmd/aomi/main.go", []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully created cmd/aomi/main.go")
}
