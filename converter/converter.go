package converter

import (
	"fmt"
	"strings"
	"tailwind-v4-to-css-converter/internal/parser"
)

type Converter struct {
	mappings     *TailwindMappings
	modern       *ModernFeatures
	classCounter int
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
	var cssRules []CSSRule
	var semanticMappings []SemanticMapping

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

	// Convert map back to slice
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

	// Clean up element name
	cleanElement := strings.ReplaceAll(element, ".", "_")
	cleanElement = strings.ToLower(cleanElement)

	// Analyze classes to determine primary purpose
	hasLayout := false
	hasTypography := false
	hasVisual := false
	isButton := false

	for _, class := range classes {
		switch class.Category {
		case "display", "alignment", "sizing", "spacing":
			hasLayout = true
		case "typography":
			hasTypography = true
		case "visual", "effects":
			hasVisual = true
		}

		// Check if it's a button element or has button-like classes
		if cleanElement == "button" || strings.Contains(class.Name, "btn") {
			isButton = true
		}
	}

	// Generate semantic name based on analysis
	if isButton {
		return fmt.Sprintf("button_%d", c.classCounter)
	} else if hasLayout && hasTypography {
		return fmt.Sprintf("%s_container_%d", cleanElement, c.classCounter)
	} else if hasLayout {
		return fmt.Sprintf("%s_layout_%d", cleanElement, c.classCounter)
	} else if hasTypography {
		return fmt.Sprintf("%s_text_%d", cleanElement, c.classCounter)
	} else if hasVisual {
		return fmt.Sprintf("%s_visual_%d", cleanElement, c.classCounter)
	}

	return fmt.Sprintf("%s_%d", cleanElement, c.classCounter)
}
