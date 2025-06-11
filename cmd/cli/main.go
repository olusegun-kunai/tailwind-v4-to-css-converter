package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"tailwind-v4-to-css-converter/internal/generator"

	"tailwind-v4-to-css-converter/converter"
	"tailwind-v4-to-css-converter/internal/parser"

	"strings"

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
		cssRules, semanticMapping := conv.ConvertWithContext(classes, path)

		// Read original file content for enhanced dual output
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file content: %w", err)
		}

		// Extract component name from path
		baseName := filepath.Base(path)
		cleanBaseName := getBaseName(baseName)
		componentName := extractComponentName(cleanBaseName)

		// Generate enhanced dual output with compilation
		enhancedGen := generator.NewEnhancedGenerator(filepath.Dir(outputPath))
		result, err := enhancedGen.GenerateEnhancedDualOutput(
			string(content),
			semanticMapping,
			cssRules,
			componentName,
		)
		if err != nil {
			if verbose {
				fmt.Printf("⚠️  Compilation failed, falling back to @apply-only output: %v\n", err)
			}
			// Fallback to original generation method
			return generateFallbackOutput(path, input, output, baseName, cleanBaseName, document, classes, semanticMapping, cssRules, verbose)
		}

		// Save enhanced output
		if err := enhancedGen.SaveEnhancedOutput(result, output, baseName, componentName); err != nil {
			return fmt.Errorf("failed to save enhanced output: %w", err)
		}

		if verbose {
			fmt.Printf("✅ Generated enhanced dual output for %s:\n", path)
			fmt.Printf("   📄 Tailwind: %s/tailwind/%s\n", output, baseName)
			fmt.Printf("   📄 Vanilla:  %s/vanilla/%s\n", output, baseName)
			fmt.Printf("   🎨 @apply:   %s/vanilla/%s.apply.css\n", output, cleanBaseName)
			fmt.Printf("   🎨 Vanilla:  %s/vanilla/%s.css\n", output, cleanBaseName)
			fmt.Printf("   🎨 Theme:    %s/tailwind/theme.css & %s/vanilla/theme.css\n", output, output)
			fmt.Printf("   ✨ Pure vanilla CSS generated successfully!\n")
		}

		return nil
	})
}

// extractComponentName extracts component name from filename
func extractComponentName(baseName string) string {
	// Convert kebab-case or snake_case to camelCase for component naming
	parts := strings.FieldsFunc(baseName, func(c rune) bool {
		return c == '-' || c == '_'
	})

	if len(parts) == 0 {
		return "component"
	}

	// Use the last meaningful part as component name
	for i := len(parts) - 1; i >= 0; i-- {
		part := strings.ToLower(parts[i])
		if part != "component" && part != "index" && part != "page" {
			return part
		}
	}

	return parts[len(parts)-1]
}

// generateFallbackOutput generates output without compilation (fallback mode)
func generateFallbackOutput(path, input, output, baseName, cleanBaseName string, document *parser.Document, classes []parser.ExtractedClass, semanticMapping []converter.SemanticMapping, cssRules []converter.CSSRule, verbose bool) error {
	relPath, _ := filepath.Rel(input, path)
	baseDir := filepath.Dir(relPath)
	outputDir := filepath.Join(output, baseDir)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
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

	// Read original file content for dual output
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file content: %w", err)
	}

	// Generate dual output
	dualOutput := generator.GenerateDualOutput(string(content), semanticMapping)

	// Create output directory structure
	tailwindDir := filepath.Join(outputDir, "tailwind")
	vanillaDir := filepath.Join(outputDir, "vanilla")

	if err := os.MkdirAll(tailwindDir, 0755); err != nil {
		return fmt.Errorf("failed to create tailwind directory: %w", err)
	}
	if err := os.MkdirAll(vanillaDir, 0755); err != nil {
		return fmt.Errorf("failed to create vanilla directory: %w", err)
	}

	// Write Tailwind version (unchanged)
	tailwindPath := filepath.Join(tailwindDir, baseName)
	if err := os.WriteFile(tailwindPath, []byte(dualOutput.TailwindVersion), 0644); err != nil {
		return fmt.Errorf("failed to write tailwind version: %w", err)
	}

	// Write vanilla CSS version
	vanillaPath := filepath.Join(vanillaDir, baseName)
	if err := os.WriteFile(vanillaPath, []byte(dualOutput.VanillaVersion), 0644); err != nil {
		return fmt.Errorf("failed to write vanilla version: %w", err)
	}

	// Write CSS file using @apply strategy
	cssFileName := strings.TrimSuffix(baseName, filepath.Ext(baseName)) + ".css"
	cssApplyPath := filepath.Join(vanillaDir, cssFileName)
	if err := os.WriteFile(cssApplyPath, []byte(dualOutput.CSSFile), 0644); err != nil {
		return fmt.Errorf("failed to write CSS file: %w", err)
	}

	// Generate compilation-ready CSS file for testing Tailwind workflow
	compilationCSSPath := filepath.Join(vanillaDir, "input.css")
	if err := generator.GenerateReadyToCompileCSS(semanticMapping, compilationCSSPath); err != nil {
		return fmt.Errorf("failed to write compilation CSS: %w", err)
	}

	if verbose {
		fmt.Printf("✅ Generated fallback output for %s:\n", path)
		fmt.Printf("   📄 Tailwind: %s\n", tailwindPath)
		fmt.Printf("   📄 Vanilla:  %s\n", vanillaPath)
		fmt.Printf("   🎨 CSS:      %s\n", cssApplyPath)
		fmt.Printf("   🔧 Compile:  %s (ready for: npx tailwindcss -i %s -o output.css)\n", compilationCSSPath, compilationCSSPath)
	}

	return nil
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
