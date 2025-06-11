package compiler

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type TailwindCompiler struct {
	workingDir string
}

type CompilationResult struct {
	VanillaCSS   string
	ApplyCSS     string
	ThemeCSS     string
	Success      bool
	ErrorMessage string
}

func NewTailwindCompiler(workingDir string) *TailwindCompiler {
	return &TailwindCompiler{
		workingDir: workingDir,
	}
}

// CompileApplyToVanilla takes @apply CSS and compiles it to pure vanilla CSS
func (tc *TailwindCompiler) CompileApplyToVanilla(applyCSS string, componentName string) (*CompilationResult, error) {
	// Create temporary directory for compilation
	tempDir, err := ioutil.TempDir("", "tailwind-compiler-")
	if err != nil {
		return &CompilationResult{
			Success:      false,
			ErrorMessage: fmt.Sprintf("Failed to create temp directory: %v", err),
		}, err
	}
	defer os.RemoveAll(tempDir)

	// Create tailwind.config.js
	configPath := filepath.Join(tempDir, "tailwind.config.js")
	if err := tc.createTailwindConfig(configPath); err != nil {
		return &CompilationResult{
			Success:      false,
			ErrorMessage: fmt.Sprintf("Failed to create Tailwind config: %v", err),
		}, err
	}

	// Create input CSS file with @apply styles
	inputPath := filepath.Join(tempDir, "input.css")
	inputCSS := tc.buildInputCSS(applyCSS)
	if err := ioutil.WriteFile(inputPath, []byte(inputCSS), 0644); err != nil {
		return &CompilationResult{
			Success:      false,
			ErrorMessage: fmt.Sprintf("Failed to write input CSS: %v", err),
		}, err
	}

	// Create output path
	outputPath := filepath.Join(tempDir, "output.css")

	// Run Tailwind compilation
	if err := tc.runTailwindCompilation(inputPath, outputPath, configPath); err != nil {
		return &CompilationResult{
			Success:      false,
			ErrorMessage: fmt.Sprintf("Tailwind compilation failed: %v", err),
		}, err
	}

	// Read compiled output
	compiledCSS, err := ioutil.ReadFile(outputPath)
	if err != nil {
		return &CompilationResult{
			Success:      false,
			ErrorMessage: fmt.Sprintf("Failed to read compiled CSS: %v", err),
		}, err
	}

	// Extract vanilla CSS (remove Tailwind base/utilities)
	vanillaCSS := tc.extractVanillaCSS(string(compiledCSS))

	// Generate theme CSS
	themeCSS := tc.generateThemeCSS(applyCSS, componentName)

	return &CompilationResult{
		VanillaCSS:   vanillaCSS,
		ApplyCSS:     applyCSS,
		ThemeCSS:     themeCSS,
		Success:      true,
		ErrorMessage: "",
	}, nil
}

// createTailwindConfig creates a minimal Tailwind config for compilation
func (tc *TailwindCompiler) createTailwindConfig(configPath string) error {
	config := `/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./input.css"],
  theme: {
    extend: {},
  },
  plugins: [],
}`

	return ioutil.WriteFile(configPath, []byte(config), 0644)
}

// buildInputCSS creates the input CSS with Tailwind directives and @apply styles
func (tc *TailwindCompiler) buildInputCSS(applyCSS string) string {
	return fmt.Sprintf(`@tailwind base;
@tailwind components;
@tailwind utilities;

%s`, applyCSS)
}

// runTailwindCompilation executes the Tailwind CLI compilation
func (tc *TailwindCompiler) runTailwindCompilation(inputPath, outputPath, configPath string) error {
	// Check if npx is available
	if _, err := exec.LookPath("npx"); err != nil {
		return fmt.Errorf("npx not found. Please install Node.js and npm")
	}

	// Run: npx tailwindcss -i input.css -o output.css --config tailwind.config.js
	cmd := exec.Command("npx", "tailwindcss",
		"-i", inputPath,
		"-o", outputPath,
		"--config", configPath,
		"--minify")

	// Set working directory to temp directory
	cmd.Dir = filepath.Dir(inputPath)

	// Capture output for debugging
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("compilation failed: %v\nOutput: %s", err, string(output))
	}

	return nil
}

// extractVanillaCSS removes Tailwind base styles and extracts only our component styles
func (tc *TailwindCompiler) extractVanillaCSS(compiledCSS string) string {
	lines := strings.Split(compiledCSS, "\n")
	var vanillaLines []string
	inComponentSection := false

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Skip Tailwind base styles and utilities
		if strings.Contains(trimmedLine, "/*! tailwindcss") ||
			strings.Contains(trimmedLine, "/* tailwind") ||
			strings.HasPrefix(trimmedLine, "/*") && strings.Contains(trimmedLine, "Tailwind") {
			continue
		}

		// Look for our component classes (they start with . and have our naming pattern)
		if strings.HasPrefix(trimmedLine, ".") &&
			(strings.Contains(trimmedLine, "-root") ||
				strings.Contains(trimmedLine, "-container") ||
				strings.Contains(trimmedLine, "-input") ||
				strings.Contains(trimmedLine, "-button") ||
				strings.Contains(trimmedLine, "-heading") ||
				strings.Contains(trimmedLine, "-text")) {
			inComponentSection = true
		}

		// Skip empty lines at the beginning
		if inComponentSection && trimmedLine != "" {
			vanillaLines = append(vanillaLines, line)
		}

		// End of a CSS rule
		if inComponentSection && trimmedLine == "}" {
			// Add a blank line after each rule for readability
			vanillaLines = append(vanillaLines, "")
		}
	}

	// Clean up the result
	result := strings.Join(vanillaLines, "\n")
	result = strings.TrimSpace(result)

	// Add header comment
	if result != "" {
		result = "/* Generated vanilla CSS */\n/* Converted from Tailwind CSS */\n\n" + result
	}

	return result
}

// generateThemeCSS extracts design tokens and creates CSS variables
func (tc *TailwindCompiler) generateThemeCSS(applyCSS string, componentName string) string {
	themeVars := make(map[string]string)

	// Extract color values
	if strings.Contains(applyCSS, "bg-blue-600") {
		themeVars["--color-blue-600"] = "#2563eb"
	}
	if strings.Contains(applyCSS, "bg-blue-700") {
		themeVars["--color-blue-700"] = "#1d4ed8"
	}
	if strings.Contains(applyCSS, "bg-gray-50") {
		themeVars["--color-gray-50"] = "#f9fafb"
	}
	if strings.Contains(applyCSS, "bg-gray-300") {
		themeVars["--color-gray-300"] = "#d1d5db"
	}
	if strings.Contains(applyCSS, "text-gray-900") {
		themeVars["--color-gray-900"] = "#111827"
	}
	if strings.Contains(applyCSS, "text-gray-600") {
		themeVars["--color-gray-600"] = "#4b5563"
	}
	if strings.Contains(applyCSS, "text-white") {
		themeVars["--color-white"] = "#ffffff"
	}

	// Extract spacing values
	if strings.Contains(applyCSS, "p-2") {
		themeVars["--spacing-2"] = "0.5rem"
	}
	if strings.Contains(applyCSS, "p-4") {
		themeVars["--spacing-4"] = "1rem"
	}
	if strings.Contains(applyCSS, "p-6") {
		themeVars["--spacing-6"] = "1.5rem"
	}
	if strings.Contains(applyCSS, "px-4") {
		themeVars["--spacing-4"] = "1rem"
	}
	if strings.Contains(applyCSS, "py-2") {
		themeVars["--spacing-2"] = "0.5rem"
	}

	// Extract border radius
	if strings.Contains(applyCSS, "rounded-md") {
		themeVars["--border-radius-md"] = "0.375rem"
	}
	if strings.Contains(applyCSS, "rounded-lg") {
		themeVars["--border-radius-lg"] = "0.5rem"
	}
	if strings.Contains(applyCSS, "rounded") {
		themeVars["--border-radius"] = "0.25rem"
	}

	// Extract font sizes
	if strings.Contains(applyCSS, "text-sm") {
		themeVars["--font-size-sm"] = "0.875rem"
	}
	if strings.Contains(applyCSS, "text-lg") {
		themeVars["--font-size-lg"] = "1.125rem"
	}
	if strings.Contains(applyCSS, "text-2xl") {
		themeVars["--font-size-2xl"] = "1.5rem"
	}

	// Build theme CSS
	if len(themeVars) == 0 {
		return "/* No theme variables extracted */"
	}

	var themeCSS strings.Builder
	themeCSS.WriteString(fmt.Sprintf("/* Theme variables for %s component */\n", componentName))
	themeCSS.WriteString(":root {\n")

	for variable, value := range themeVars {
		themeCSS.WriteString(fmt.Sprintf("  %s: %s;\n", variable, value))
	}

	themeCSS.WriteString("}")

	return themeCSS.String()
}
