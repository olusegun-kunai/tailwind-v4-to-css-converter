package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"tailwind-v4-to-css-converter/converter"
	"tailwind-v4-to-css-converter/internal/compiler"
)

type EnhancedGenerator struct {
	tailwindCompiler *compiler.TailwindCompiler
}

type GenerationResult struct {
	TailwindVersion string
	VanillaVersion  string
	ApplyCSS        string
	VanillaCSS      string
	ThemeCSS        string
	Success         bool
	ErrorMessage    string
}

func NewEnhancedGenerator(workingDir string) *EnhancedGenerator {
	return &EnhancedGenerator{
		tailwindCompiler: compiler.NewTailwindCompiler(workingDir),
	}
}

// GenerateEnhancedDualOutput creates complete dual documentation with compiled vanilla CSS
func (eg *EnhancedGenerator) GenerateEnhancedDualOutput(
	originalContent string,
	semanticMappings []converter.SemanticMapping,
	cssRules []converter.CSSRule,
	componentName string,
) (*GenerationResult, error) {

	// Generate @apply CSS from semantic mappings
	applyCSS := eg.generateApplyCSS(cssRules)

	// Compile @apply CSS to vanilla CSS using Tailwind
	compilationResult, err := eg.tailwindCompiler.CompileApplyToVanilla(applyCSS, componentName)
	if err != nil {
		return &GenerationResult{
			Success:      false,
			ErrorMessage: fmt.Sprintf("Failed to compile CSS: %v", err),
		}, err
	}

	if !compilationResult.Success {
		return &GenerationResult{
			Success:      false,
			ErrorMessage: compilationResult.ErrorMessage,
		}, fmt.Errorf("compilation failed: %s", compilationResult.ErrorMessage)
	}

	// Generate component versions
	tailwindVersion := originalContent // Keep original Tailwind version unchanged
	vanillaVersion := eg.generateVanillaComponent(originalContent, semanticMappings, componentName)

	return &GenerationResult{
		TailwindVersion: tailwindVersion,
		VanillaVersion:  vanillaVersion,
		ApplyCSS:        compilationResult.ApplyCSS,
		VanillaCSS:      compilationResult.VanillaCSS,
		ThemeCSS:        compilationResult.ThemeCSS,
		Success:         true,
		ErrorMessage:    "",
	}, nil
}

// SaveEnhancedOutput saves all generated files in the proper structure
func (eg *EnhancedGenerator) SaveEnhancedOutput(
	result *GenerationResult,
	outputDir string,
	fileName string,
	componentName string,
) error {

	// Create directory structure
	tailwindDir := filepath.Join(outputDir, "tailwind")
	vanillaDir := filepath.Join(outputDir, "vanilla")

	if err := os.MkdirAll(tailwindDir, 0755); err != nil {
		return fmt.Errorf("failed to create tailwind directory: %w", err)
	}
	if err := os.MkdirAll(vanillaDir, 0755); err != nil {
		return fmt.Errorf("failed to create vanilla directory: %w", err)
	}

	// Save Tailwind version
	tailwindPath := filepath.Join(tailwindDir, fileName)
	if err := os.WriteFile(tailwindPath, []byte(result.TailwindVersion), 0644); err != nil {
		return fmt.Errorf("failed to write tailwind version: %w", err)
	}

	// Save vanilla component
	vanillaPath := filepath.Join(vanillaDir, fileName)
	if err := os.WriteFile(vanillaPath, []byte(result.VanillaVersion), 0644); err != nil {
		return fmt.Errorf("failed to write vanilla version: %w", err)
	}

	// Save @apply CSS (for development/debugging)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	applyPath := filepath.Join(vanillaDir, fmt.Sprintf("%s.apply.css", baseName))
	if err := os.WriteFile(applyPath, []byte(result.ApplyCSS), 0644); err != nil {
		return fmt.Errorf("failed to write @apply CSS: %w", err)
	}

	// Save compiled vanilla CSS
	vanillaCSSPath := filepath.Join(vanillaDir, fmt.Sprintf("%s.css", baseName))
	if err := os.WriteFile(vanillaCSSPath, []byte(result.VanillaCSS), 0644); err != nil {
		return fmt.Errorf("failed to write vanilla CSS: %w", err)
	}

	// Save theme CSS (both versions use the same theme)
	tailwindThemePath := filepath.Join(tailwindDir, "theme.css")
	vanillaThemePath := filepath.Join(vanillaDir, "theme.css")

	if err := os.WriteFile(tailwindThemePath, []byte(result.ThemeCSS), 0644); err != nil {
		return fmt.Errorf("failed to write tailwind theme CSS: %w", err)
	}
	if err := os.WriteFile(vanillaThemePath, []byte(result.ThemeCSS), 0644); err != nil {
		return fmt.Errorf("failed to write vanilla theme CSS: %w", err)
	}

	return nil
}

// generateApplyCSS creates @apply CSS from semantic mappings and CSS rules
func (eg *EnhancedGenerator) generateApplyCSS(cssRules []converter.CSSRule) string {
	var cssBuilder strings.Builder

	cssBuilder.WriteString("/* Generated CSS using @apply strategy */\n")
	cssBuilder.WriteString("/* This CSS will be compiled to vanilla CSS */\n\n")

	for _, rule := range cssRules {
		cssBuilder.WriteString(fmt.Sprintf("%s {\n", rule.Selector))

		// Extract Tailwind classes from properties (we need to reverse-engineer this)
		tailwindClasses := eg.extractTailwindClasses(rule.Properties)
		if len(tailwindClasses) > 0 {
			cssBuilder.WriteString(fmt.Sprintf("  @apply %s;\n", strings.Join(tailwindClasses, " ")))
		}

		cssBuilder.WriteString("}\n\n")
	}

	return cssBuilder.String()
}

// extractTailwindClasses converts CSS properties back to Tailwind classes for @apply
func (eg *EnhancedGenerator) extractTailwindClasses(properties []converter.CSSProperty) []string {
	var classes []string

	for _, prop := range properties {
		// Skip comment properties
		if strings.HasPrefix(prop.Name, "/*") {
			continue
		}

		// Map CSS properties back to Tailwind classes
		switch prop.Name {
		case "display":
			if prop.Value == "flex" {
				classes = append(classes, "flex")
			} else if prop.Value == "grid" {
				classes = append(classes, "grid")
			} else if prop.Value == "block" {
				classes = append(classes, "block")
			}
		case "flex-direction":
			if prop.Value == "column" {
				classes = append(classes, "flex-col")
			}
		case "align-items":
			if prop.Value == "center" {
				classes = append(classes, "items-center")
			}
		case "justify-content":
			if prop.Value == "center" {
				classes = append(classes, "justify-center")
			} else if prop.Value == "space-between" {
				classes = append(classes, "justify-between")
			}
		case "background-color":
			classes = append(classes, eg.colorToTailwind("bg", prop.Value))
		case "color":
			classes = append(classes, eg.colorToTailwind("text", prop.Value))
		case "padding":
			classes = append(classes, eg.spacingToTailwind("p", prop.Value))
		case "padding-left", "padding-right":
			classes = append(classes, eg.spacingToTailwind("px", prop.Value))
		case "padding-top", "padding-bottom":
			classes = append(classes, eg.spacingToTailwind("py", prop.Value))
		case "width":
			classes = append(classes, eg.sizeToTailwind("w", prop.Value))
		case "height":
			classes = append(classes, eg.sizeToTailwind("h", prop.Value))
		case "border-radius":
			classes = append(classes, eg.borderRadiusToTailwind(prop.Value))
		case "font-weight":
			if prop.Value == "600" || prop.Value == "bold" {
				classes = append(classes, "font-semibold")
			}
		case "font-size":
			classes = append(classes, eg.fontSizeToTailwind(prop.Value))
		}
	}

	return classes
}

// Helper functions to map CSS values back to Tailwind classes
func (eg *EnhancedGenerator) colorToTailwind(prefix, value string) string {
	colorMap := map[string]string{
		"#2563eb": "blue-600",
		"#1d4ed8": "blue-700",
		"#f9fafb": "gray-50",
		"#d1d5db": "gray-300",
		"#111827": "gray-900",
		"#4b5563": "gray-600",
		"#ffffff": "white",
	}

	if tailwindColor, exists := colorMap[value]; exists {
		return fmt.Sprintf("%s-%s", prefix, tailwindColor)
	}
	return ""
}

func (eg *EnhancedGenerator) spacingToTailwind(prefix, value string) string {
	spacingMap := map[string]string{
		"0.25rem": "1",
		"0.5rem":  "2",
		"0.75rem": "3",
		"1rem":    "4",
		"1.25rem": "5",
		"1.5rem":  "6",
		"2rem":    "8",
	}

	if tailwindSpacing, exists := spacingMap[value]; exists {
		return fmt.Sprintf("%s-%s", prefix, tailwindSpacing)
	}
	return ""
}

func (eg *EnhancedGenerator) sizeToTailwind(prefix, value string) string {
	sizeMap := map[string]string{
		"100%":    "full",
		"3rem":    "12",
		"50%":     "1/2",
		"33.333%": "1/3",
		"25%":     "1/4",
	}

	if tailwindSize, exists := sizeMap[value]; exists {
		return fmt.Sprintf("%s-%s", prefix, tailwindSize)
	}
	return ""
}

func (eg *EnhancedGenerator) borderRadiusToTailwind(value string) string {
	radiusMap := map[string]string{
		"0.25rem":  "rounded",
		"0.375rem": "rounded-md",
		"0.5rem":   "rounded-lg",
		"9999px":   "rounded-full",
	}

	if tailwindRadius, exists := radiusMap[value]; exists {
		return tailwindRadius
	}
	return ""
}

func (eg *EnhancedGenerator) fontSizeToTailwind(value string) string {
	fontSizeMap := map[string]string{
		"0.75rem":  "text-xs",
		"0.875rem": "text-sm",
		"1rem":     "text-base",
		"1.125rem": "text-lg",
		"1.25rem":  "text-xl",
		"1.5rem":   "text-2xl",
	}

	if tailwindSize, exists := fontSizeMap[value]; exists {
		return tailwindSize
	}
	return ""
}

// generateVanillaComponent creates the vanilla version with CSS module imports
func (eg *EnhancedGenerator) generateVanillaComponent(
	originalContent string,
	semanticMappings []converter.SemanticMapping,
	componentName string,
) string {

	content := originalContent

	// Add CSS module import at the top
	baseName := strings.ToLower(componentName)
	importStatement := fmt.Sprintf("import styles from './%s.css';\n", baseName)

	// Find the first import or beginning of file to insert our import
	lines := strings.Split(content, "\n")
	var newLines []string
	importInserted := false

	for i, line := range lines {
		if !importInserted && (strings.HasPrefix(strings.TrimSpace(line), "import") ||
			(!strings.HasPrefix(strings.TrimSpace(line), "import") && strings.TrimSpace(line) != "")) {
			// Insert our import before the first non-import line or after existing imports
			if strings.HasPrefix(strings.TrimSpace(line), "import") {
				newLines = append(newLines, line)
				// Add our import after existing imports
				if i+1 < len(lines) && !strings.HasPrefix(strings.TrimSpace(lines[i+1]), "import") {
					newLines = append(newLines, "", importStatement)
					importInserted = true
				}
			} else {
				newLines = append(newLines, importStatement, "", line)
				importInserted = true
			}
		} else {
			newLines = append(newLines, line)
		}
	}

	if !importInserted {
		newLines = append([]string{importStatement, ""}, newLines...)
	}

	content = strings.Join(newLines, "\n")

	// Replace Tailwind classes with semantic class names
	for _, mapping := range semanticMappings {
		// Replace className="tailwind classes" with className={styles.semanticName}
		oldPattern := fmt.Sprintf(`className="%s"`, mapping.OriginalClasses)
		newPattern := fmt.Sprintf(`className={styles.%s}`, mapping.SemanticName)
		content = strings.ReplaceAll(content, oldPattern, newPattern)

		// Also handle class attribute (for Qwik)
		oldPattern = fmt.Sprintf(`class="%s"`, mapping.OriginalClasses)
		newPattern = fmt.Sprintf(`class={styles.%s}`, mapping.SemanticName)
		content = strings.ReplaceAll(content, oldPattern, newPattern)
	}

	return content
}
