package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"tailwind-v4-to-css-converter/internal/generator"

	"tailwind-v4-to-css-converter/converter"
	"tailwind-v4-to-css-converter/internal/parser"

	"github.com/spf13/cobra"
)

var (
	inputPath  string
	outputPath string
	verbose    bool
)

var rootCmd = &cobra.Command{
	Use:   "tailwind-converter",
	Short: "Convert Tailwind CSS to vanilla CSS modules",
	Long:  "A tool to convert Tailwind CSS classes in HTML/JSX/Qwik files to semantic CSS modules",
	Run:   run,
}

func init() {
	rootCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Input file or directory")
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output directory")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("output")
}

func run(cmd *cobra.Command, args []string) {
	if verbose {
		fmt.Printf("Converting from: %s\n", inputPath)
		fmt.Printf("Output to: %s\n", outputPath)
	}

	// Ensure output directory exists
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Process files
	if err := processPath(inputPath, outputPath); err != nil {
		fmt.Printf("Error processing files: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Conversion completed successfully!")
}

func processPath(input, output string) error {
	return filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// Only process HTML-like files
		ext := filepath.Ext(path)
		if ext != ".html" && ext != ".jsx" && ext != ".tsx" && ext != ".vue" {
			return nil
		}

		if verbose {
			fmt.Printf("Processing: %s\n", path)
		}

		// Parse file
		htmlParser := parser.NewHTMLParser()
		document, err := htmlParser.ParseFile(path)
		if err != nil {
			return fmt.Errorf("error parsing %s: %v", path, err)
		}

		// Extract classes
		classExtractor := parser.NewClassExtractor()
		classes := classExtractor.Extract(document)

		if len(classes) == 0 {
			return nil // No Tailwind classes found
		}

		// Convert classes
		conv := converter.NewConverter()
		cssRules, semanticMapping := conv.Convert(classes)

		// Generate output files
		relPath, _ := filepath.Rel(input, path)
		baseDir := filepath.Dir(relPath)
		baseName := filepath.Base(relPath)
		outputDir := filepath.Join(output, baseDir)

		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return err
		}

		// Get clean base name (without extension)
		cleanBaseName := getBaseName(baseName)
		if cleanBaseName == "" {
			return fmt.Errorf("invalid filename: %s", baseName)
		}

		// Generate CSS file
		cssGen := generator.NewCSSGenerator()
		cssPath := filepath.Join(outputDir, cleanBaseName+".module.css")
		if err := cssGen.Generate(cssRules, cssPath); err != nil {
			return err
		}

		// Generate updated HTML file
		htmlGen := generator.NewHTMLGenerator()
		htmlPath := filepath.Join(outputDir, baseName)
		if err := htmlGen.Generate(document, semanticMapping, htmlPath, cleanBaseName); err != nil {
			return err
		}

		return nil
	})
}

func getBaseName(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
