package converter

import (
	"fmt"
	"strings"
	"tailwind-v4-to-css-converter/internal/parser"
)

type Converter struct {
	mappings         *TailwindMappings
	modern           *ModernFeatures
	classCounter     int
	componentContext string
}

type CSSRule struct {
	Selector   string
	Properties []CSSProperty
}

type CSSProperty struct {
	Name  string
	Value string
}

type SemanticMapping struct {
	OriginalClasses string
	SemanticName    string
}

func NewConverter() *Converter {
	return &Converter{
		mappings:     NewTailwindMappings(),
		modern:       NewModernFeatures(),
		classCounter: 0,
	}
}

func (c *Converter) Convert(classes []parser.ExtractedClass) ([]CSSRule, []SemanticMapping) {
	return c.ConvertWithContext(classes, "")
}

func (c *Converter) ConvertWithContext(classes []parser.ExtractedClass, filename string) ([]CSSRule, []SemanticMapping) {
	var cssRules []CSSRule
	var semanticMappings []SemanticMapping

	// Extract component context from filename if available
	c.componentContext = c.extractComponentFromFilename(filename)

	// Group classes by element and create consolidated semantic classes
	elementGroups := c.groupByElement(classes)

	for element, elementClasses := range elementGroups {
		// Create semantic class name
		semanticName := c.generateSemanticName(element, elementClasses)

		// Convert classes to CSS properties and deduplicate
		properties := c.convertAndDeduplicateProperties(elementClasses)

		if len(properties) > 0 {
			cssRules = append(cssRules, CSSRule{
				Selector:   "." + semanticName,
				Properties: properties,
			})

			// Create mapping with all original classes for this element
			var originalClassNames []string
			for _, class := range elementClasses {
				originalClassNames = append(originalClassNames, class.Name)
			}

			semanticMappings = append(semanticMappings, SemanticMapping{
				OriginalClasses: strings.Join(originalClassNames, " "),
				SemanticName:    semanticName,
			})
		}
	}

	return cssRules, semanticMappings
}

func (c *Converter) groupByElement(classes []parser.ExtractedClass) map[string][]parser.ExtractedClass {
	groups := make(map[string][]parser.ExtractedClass)

	for _, class := range classes {
		// Use element as the key for grouping
		groups[class.Context] = append(groups[class.Context], class)
	}

	return groups
}

func (c *Converter) convertAndDeduplicateProperties(classes []parser.ExtractedClass) []CSSProperty {
	propertyMap := make(map[string]CSSProperty)
	var unknownClasses []string

	for _, class := range classes {
		// Try to convert using mappings first
		if cssProps := c.mappings.Convert(class.Name); len(cssProps) > 0 {
			for _, prop := range cssProps {
				// Use property name as key to deduplicate (later values override earlier ones)
				propertyMap[prop.Name] = prop
			}
		} else if modernProps := c.modern.Convert(class.Name); len(modernProps) > 0 {
			for _, prop := range modernProps {
				propertyMap[prop.Name] = prop
			}
		} else {
			// Collect unknown classes to add as comments
			unknownClasses = append(unknownClasses, class.Name)
		}
	}

	var properties []CSSProperty
	for _, prop := range propertyMap {
		properties = append(properties, prop)
	}

	// Add unknown classes as comments
	if len(unknownClasses) > 0 {
		properties = append(properties, CSSProperty{
			Name:  "/* Unknown classes */",
			Value: strings.Join(unknownClasses, " "),
		})
	}

	return properties
}

func (c *Converter) generateSemanticName(element string, classes []parser.ExtractedClass) string {
	c.classCounter++

	// Clean up element name and detect component type
	cleanElement := strings.ReplaceAll(element, ".", "_")
	cleanElement = strings.ToLower(cleanElement)

	// Step 1: Handle component syntax like "Otp.Root", "Accordion.Trigger"
	if strings.Contains(element, ".") {
		parts := strings.Split(element, ".")
		if len(parts) >= 2 {
			componentName := strings.ToLower(parts[0]) // "otp", "accordion"
			elementType := strings.ToLower(parts[1])   // "root", "trigger"
			return fmt.Sprintf("%s-%s", componentName, elementType)
		}
	}

	// Step 2: Extract component context from file/context clues
	componentName := c.detectComponentContext(classes)

	// Step 3: Analyze classes to understand element purpose
	classNames := make([]string, 0, len(classes))
	for _, class := range classes {
		classNames = append(classNames, class.Name)
	}
	allClasses := strings.Join(classNames, " ")

	// Step 4: Detect specific UI patterns first (high confidence)
	if semanticName := c.detectSpecificUIPattern(allClasses, cleanElement, componentName); semanticName != "" {
		return semanticName
	}

	// Step 5: Generate names based on element type and context
	return c.generateContextualName(cleanElement, allClasses, componentName, c.classCounter)
}

// detectComponentContext tries to infer component name from classes or context
func (c *Converter) detectComponentContext(classes []parser.ExtractedClass) string {
	// First, check if we have component context from filename
	if c.componentContext != "" {
		return c.componentContext
	}

	// Look for component-specific class patterns
	for _, class := range classes {
		className := strings.ToLower(class.Name)

		// Check for component indicators in class names
		if strings.Contains(className, "otp") {
			return "otp"
		}
		if strings.Contains(className, "modal") {
			return "modal"
		}
		if strings.Contains(className, "accordion") {
			return "accordion"
		}
		if strings.Contains(className, "dropdown") {
			return "dropdown"
		}
		if strings.Contains(className, "card") {
			return "card"
		}
		if strings.Contains(className, "hero") {
			return "hero"
		}
		if strings.Contains(className, "nav") {
			return "navigation"
		}
	}

	return "" // No component context detected
}

// detectSpecificUIPattern identifies specific UI component patterns with high confidence
func (c *Converter) detectSpecificUIPattern(allClasses, element, componentName string) string {
	// Modal overlay pattern (very specific)
	if strings.Contains(allClasses, "fixed") && strings.Contains(allClasses, "inset-0") &&
		(strings.Contains(allClasses, "bg-black") || strings.Contains(allClasses, "bg-gray")) {
		if componentName != "" {
			return fmt.Sprintf("%s-overlay", componentName)
		}
		return "modal-overlay"
	}

	// Dropdown menu pattern (specific positioning + styling)
	if strings.Contains(allClasses, "absolute") &&
		(strings.Contains(allClasses, "top-") || strings.Contains(allClasses, "bottom-")) &&
		strings.Contains(allClasses, "shadow") && strings.Contains(allClasses, "bg-white") {
		if componentName != "" {
			return fmt.Sprintf("%s-menu", componentName)
		}
		return "dropdown-menu"
	}

	// Card container pattern (rounded + shadow + padding + background)
	if strings.Contains(allClasses, "rounded") && strings.Contains(allClasses, "shadow") &&
		strings.Contains(allClasses, "bg-white") &&
		(strings.Contains(allClasses, "p-") || strings.Contains(allClasses, "px-")) {
		if componentName != "" {
			return fmt.Sprintf("%s-card", componentName)
		}
		return "card-container"
	}

	// Toast/notification pattern (fixed positioning + colored background)
	if strings.Contains(allClasses, "fixed") &&
		(strings.Contains(allClasses, "top-") || strings.Contains(allClasses, "bottom-")) &&
		(strings.Contains(allClasses, "bg-green") || strings.Contains(allClasses, "bg-red") ||
			strings.Contains(allClasses, "bg-yellow") || strings.Contains(allClasses, "bg-blue")) {
		return "toast-notification"
	}

	// Hero section pattern (full height + center alignment)
	if strings.Contains(allClasses, "min-h-screen") && strings.Contains(allClasses, "flex") &&
		strings.Contains(allClasses, "items-center") && strings.Contains(allClasses, "justify-center") {
		if componentName != "" {
			return fmt.Sprintf("%s-hero", componentName)
		}
		return "hero-section"
	}

	return "" // No specific pattern detected
}

/*
Utilitites class
div
@apply doing the div


*/

// generateContextualName creates meaningful names based on element type and styling context
func (c *Converter) generateContextualName(element, allClasses, componentName string, counter int) string {
	// Analyze styling to understand purpose
	isInteractive := strings.Contains(allClasses, "hover:") || strings.Contains(allClasses, "focus:") ||
		strings.Contains(allClasses, "active:") || strings.Contains(allClasses, "cursor-pointer")
	isContainer := strings.Contains(allClasses, "flex") || strings.Contains(allClasses, "grid") ||
		strings.Contains(allClasses, "p-") || strings.Contains(allClasses, "px-") ||
		strings.Contains(allClasses, "space-")
	isTypography := strings.Contains(allClasses, "text-") || strings.Contains(allClasses, "font-")
	isInput := strings.Contains(allClasses, "border-") &&
		(strings.Contains(allClasses, "w-") || strings.Contains(allClasses, "h-"))

	// Generate names based on element type and component context
	switch element {
	case "button":
		if componentName != "" {
			if isInteractive {
				return fmt.Sprintf("%s-button", componentName)
			}
			return fmt.Sprintf("%s-trigger", componentName)
		}
		if isInteractive {
			return "button-primary"
		}
		return "button"

	case "input":
		if componentName != "" {
			return fmt.Sprintf("%s-input", componentName)
		}
		return "input-field"

	case "h1", "h2", "h3", "h4", "h5", "h6":
		if componentName != "" {
			return fmt.Sprintf("%s-heading", componentName)
		}
		return fmt.Sprintf("%s-text", element)

	case "p":
		if componentName != "" {
			return fmt.Sprintf("%s-text", componentName)
		}
		return "p-text"

	case "div":
		// For divs, be more specific about their purpose
		if componentName != "" {
			// Root container detection
			if strings.Contains(allClasses, "min-h-screen") ||
				(isContainer && strings.Contains(allClasses, "justify-center")) {
				return fmt.Sprintf("%s-root", componentName)
			}
			// Input group detection
			if isContainer && isInput {
				return fmt.Sprintf("%s-input-group", componentName)
			}
			// Content container
			if isContainer && isTypography {
				return fmt.Sprintf("%s-content", componentName)
			}
			// Generic container
			if isContainer {
				return fmt.Sprintf("%s-container", componentName)
			}
			// Text wrapper
			if isTypography {
				return fmt.Sprintf("%s-text", componentName)
			}
			// Fallback with component context
			return fmt.Sprintf("%s-element-%d", componentName, counter)
		}

		// Without component context, use generic but meaningful names
		if isContainer && strings.Contains(allClasses, "grid-cols") {
			return "grid-container"
		}
		if isContainer && isTypography {
			return "content-container"
		}
		if isContainer {
			return "layout-container"
		}
		if isTypography {
			return "text-container"
		}
		return fmt.Sprintf("container-%d", counter)

	case "form":
		if componentName != "" {
			return fmt.Sprintf("%s-form", componentName)
		}
		return "form-container"

	case "nav":
		return "navigation"

	case "header":
		return "header-container"

	case "section":
		if componentName != "" {
			return fmt.Sprintf("%s-section", componentName)
		}
		return "section-container"

	default:
		// For custom elements or unknown tags
		if componentName != "" {
			return fmt.Sprintf("%s-%s", componentName, element)
		}
		if isInteractive {
			return fmt.Sprintf("%s-interactive", element)
		}
		if isContainer {
			return fmt.Sprintf("%s-container", element)
		}
		return element
	}
}

func (c *Converter) extractComponentFromFilename(filename string) string {
	if filename == "" {
		return ""
	}

	// Remove file extension and path
	baseName := filename
	if lastSlash := strings.LastIndex(filename, "/"); lastSlash != -1 {
		baseName = filename[lastSlash+1:]
	}
	if lastBackslash := strings.LastIndex(baseName, "\\"); lastBackslash != -1 {
		baseName = baseName[lastBackslash+1:]
	}
	if lastDot := strings.LastIndex(baseName, "."); lastDot != -1 {
		baseName = baseName[:lastDot]
	}

	// Convert to lowercase for processing
	baseName = strings.ToLower(baseName)

	// Extract component patterns from common naming conventions

	// Pattern: "qwik-otp" -> "otp"
	// Pattern: "react-modal" -> "modal"
	// Pattern: "vue-accordion" -> "accordion"
	if parts := strings.Split(baseName, "-"); len(parts) >= 2 {
		// If it starts with framework name, use the rest
		if parts[0] == "qwik" || parts[0] == "react" || parts[0] == "vue" || parts[0] == "svelte" {
			return strings.Join(parts[1:], "-")
		}
		// Otherwise use the last part as the component name
		return parts[len(parts)-1]
	}

	// Pattern: "otp_input" -> "otp"
	if parts := strings.Split(baseName, "_"); len(parts) >= 2 {
		return parts[0]
	}

	// Single word component names
	commonComponents := map[string]string{
		"modal":      "modal",
		"dropdown":   "dropdown",
		"accordion":  "accordion",
		"carousel":   "carousel",
		"tooltip":    "tooltip",
		"popover":    "popover",
		"dialog":     "dialog",
		"sidebar":    "sidebar",
		"navbar":     "navbar",
		"navigation": "navigation",
		"header":     "header",
		"footer":     "footer",
		"card":       "card",
		"button":     "button",
		"input":      "input",
		"form":       "form",
		"table":      "table",
		"grid":       "grid",
		"list":       "list",
		"menu":       "menu",
		"breadcrumb": "breadcrumb",
		"pagination": "pagination",
		"tabs":       "tabs",
		"badge":      "badge",
		"avatar":     "avatar",
		"spinner":    "spinner",
		"loader":     "loader",
		"toast":      "toast",
		"alert":      "alert",
		"otp":        "otp",
		"hero":       "hero",
		"banner":     "banner",
		"section":    "section",
	}

	if component, exists := commonComponents[baseName]; exists {
		return component
	}

	// If filename contains known component keywords, extract them
	for keyword := range commonComponents {
		if strings.Contains(baseName, keyword) {
			return keyword
		}
	}

	return baseName // Fallback to the full basename
}
