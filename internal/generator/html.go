package generator

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"tailwind-v4-to-css-converter/converter"
	"tailwind-v4-to-css-converter/internal/parser"
)

type HTMLGenerator struct {
	classRegex *regexp.Regexp
}

func NewHTMLGenerator() *HTMLGenerator {
	return &HTMLGenerator{
		classRegex: regexp.MustCompile(`(class|className)=["']([^"']+)["']`),
	}
}

func (g *HTMLGenerator) Generate(document *parser.Document, semanticMappings []converter.SemanticMapping, outputPath, moduleName string) error {
	// Create updated HTML content
	updatedContent := g.updateClassReferences(document, semanticMappings, moduleName)

	// Add import statement for CSS module
	updatedContent = g.addCSSModuleImport(updatedContent, moduleName)

	// Write to file
	return os.WriteFile(outputPath, []byte(updatedContent), 0644)
}

func (g *HTMLGenerator) updateClassReferences(document *parser.Document, semanticMappings []converter.SemanticMapping, moduleName string) string {
	content := document.Content

	// Create a mapping from original classes to semantic names
	classMap := make(map[string]string)
	for _, mapping := range semanticMappings {
		classMap[mapping.OriginalClasses] = mapping.SemanticName
	}

	// Process each class reference
	result := g.classRegex.ReplaceAllStringFunc(content, func(match string) string {
		return g.replaceClassAttribute(match, classMap, moduleName)
	})

	return result
}

func (g *HTMLGenerator) replaceClassAttribute(classAttr string, classMap map[string]string, moduleName string) string {
	// Extract the attribute name and class values
	parts := g.classRegex.FindStringSubmatch(classAttr)
	if len(parts) < 3 {
		return classAttr
	}

	attrName := parts[1] // "class" or "className"
	classValues := parts[2]

	// Split class values
	classes := strings.Fields(classValues)

	// Create sets to track processed classes and avoid duplicates
	processedTailwindClasses := make(map[string]bool)
	var remainingClasses []string
	var semanticClasses []string

	// Find the best matching semantic mapping for this set of classes
	var bestMatch string
	var bestMatchClasses []string
	maxMatches := 0

	for originalClasses, semanticName := range classMap {
		originalClassList := strings.Fields(originalClasses)
		matchCount := 0

		// Count how many classes from this element match the mapping
		for _, class := range classes {
			if g.containsClass(originalClassList, class) && g.isTailwindClass(class) {
				matchCount++
			}
		}

		// Use the mapping with the most matches
		if matchCount > maxMatches {
			maxMatches = matchCount
			bestMatch = semanticName
			bestMatchClasses = originalClassList
		}
	}

	// Apply the best match if found
	if bestMatch != "" && maxMatches > 0 {
		// Mark matched Tailwind classes as processed
		for _, class := range bestMatchClasses {
			if g.containsClass(classes, class) {
				processedTailwindClasses[class] = true
			}
		}
		semanticClasses = append(semanticClasses, fmt.Sprintf("styles.%s", bestMatch))
	}

	// Add remaining non-Tailwind classes and unprocessed Tailwind classes
	for _, class := range classes {
		if !processedTailwindClasses[class] && !g.isTailwindClass(class) {
			remainingClasses = append(remainingClasses, class)
		}
	}

	// Build the final class attribute
	var finalClasses []string

	// Add remaining non-Tailwind classes as a string
	if len(remainingClasses) > 0 {
		finalClasses = append(finalClasses, `"`+strings.Join(remainingClasses, " ")+`"`)
	}

	// Add semantic classes
	finalClasses = append(finalClasses, semanticClasses...)

	// Generate the new class attribute
	if len(finalClasses) == 0 {
		return "" // Remove empty class attributes
	} else if len(finalClasses) == 1 && !strings.HasPrefix(finalClasses[0], "styles.") {
		// Single non-semantic class
		return fmt.Sprintf(`%s=%s`, attrName, finalClasses[0])
	} else if len(finalClasses) == 1 {
		// Single semantic class
		return fmt.Sprintf(`%s={%s}`, attrName, finalClasses[0])
	} else {
		// Multiple classes need to be combined
		return fmt.Sprintf(`%s={[%s].join(' ')}`, attrName, strings.Join(finalClasses, ", "))
	}
}

func (g *HTMLGenerator) containsClass(classList []string, targetClass string) bool {
	for _, class := range classList {
		if class == targetClass {
			return true
		}
	}
	return false
}

func (g *HTMLGenerator) isTailwindClass(class string) bool {
	// Same logic as in parser/html.go
	tailwindPrefixes := []string{
		"flex", "grid", "block", "inline", "hidden",
		"text-", "bg-", "border-", "p-", "m-", "w-", "h-",
		"items-", "justify-", "gap-", "space-", "rounded",
		"font-", "leading-", "tracking-", "opacity-",
		"hover:", "focus:", "active:", "disabled:",
		"sm:", "md:", "lg:", "xl:", "2xl:",
	}

	for _, prefix := range tailwindPrefixes {
		if strings.HasPrefix(class, prefix) || class == strings.TrimSuffix(prefix, "-") {
			return true
		}
	}
	return false
}

func (g *HTMLGenerator) addCSSModuleImport(content, moduleName string) string {
	// Check if import already exists
	importRegex := regexp.MustCompile(`import\s+.*from\s+['"].*\.module\.css['"]`)
	if importRegex.MatchString(content) {
		return content // Import already exists
	}

	// Find the position to insert the import
	lines := strings.Split(content, "\n")
	insertIndex := g.findImportInsertPosition(lines)

	// Create the import statement
	importStatement := fmt.Sprintf("import styles from './%s.module.css';", moduleName)

	// Insert the import
	if insertIndex >= 0 && insertIndex < len(lines) {
		newLines := make([]string, 0, len(lines)+1)
		newLines = append(newLines, lines[:insertIndex]...)
		newLines = append(newLines, importStatement)
		newLines = append(newLines, lines[insertIndex:]...)
		return strings.Join(newLines, "\n")
	} else {
		// If no good position found, add at the beginning
		return importStatement + "\n\n" + content
	}
}

func (g *HTMLGenerator) findImportInsertPosition(lines []string) int {
	// Find the best position to insert CSS module import
	lastImportIndex := -1

	for i, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Skip comments and empty lines
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "//") || strings.HasPrefix(trimmedLine, "/*") {
			continue
		}

		// Check if this is an import statement
		if strings.HasPrefix(trimmedLine, "import ") {
			lastImportIndex = i
		} else if lastImportIndex >= 0 {
			// Found first non-import line after imports
			return lastImportIndex + 1
		} else if strings.HasPrefix(trimmedLine, "export ") {
			// Found export before any imports, insert before it
			return i
		}
	}

	// If we found imports but no code after, insert after last import
	if lastImportIndex >= 0 {
		return lastImportIndex + 1
	}

	// If no imports found, insert at the beginning
	return 0
}

func (g *HTMLGenerator) GenerateWithCustomTemplate(document *parser.Document, semanticMappings []converter.SemanticMapping, outputPath, moduleName string, template HTMLTemplate) error {
	// Apply custom template transformations
	content := document.Content

	if template.PreProcessor != nil {
		content = template.PreProcessor(content)
	}

	// Update class references
	updatedContent := g.updateClassReferences(&parser.Document{Content: content, ClassRefs: document.ClassRefs}, semanticMappings, moduleName)

	// Add imports with custom logic
	if template.ImportGenerator != nil {
		updatedContent = template.ImportGenerator(updatedContent, moduleName)
	} else {
		updatedContent = g.addCSSModuleImport(updatedContent, moduleName)
	}

	if template.PostProcessor != nil {
		updatedContent = template.PostProcessor(updatedContent)
	}

	// Write to file
	return os.WriteFile(outputPath, []byte(updatedContent), 0644)
}

type HTMLTemplate struct {
	PreProcessor    func(string) string
	ImportGenerator func(content, moduleName string) string
	PostProcessor   func(string) string
}
