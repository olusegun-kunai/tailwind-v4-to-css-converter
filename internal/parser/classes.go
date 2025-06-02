package parser

import (
	"sort"
	"strings"
)

type ClassExtractor struct{}

type ExtractedClass struct {
	Name     string
	Category string
	Context  string // Element context where it was found
}

func NewClassExtractor() *ClassExtractor {
	return &ClassExtractor{}
}

func (e *ClassExtractor) Extract(doc *Document) []ExtractedClass {
	classMap := make(map[string]ExtractedClass)

	for _, ref := range doc.ClassRefs {
		for _, class := range ref.Classes {
			category := e.categorizeClass(class)
			key := class + "_" + category

			if _, exists := classMap[key]; !exists {
				classMap[key] = ExtractedClass{
					Name:     class,
					Category: category,
					Context:  ref.Element,
				}
			}
		}
	}

	// Convert map to slice and sort
	classes := make([]ExtractedClass, 0, len(classMap))
	for _, class := range classMap {
		classes = append(classes, class)
	}

	sort.Slice(classes, func(i, j int) bool {
		return classes[i].Name < classes[j].Name
	})

	return classes
}

func (e *ClassExtractor) categorizeClass(class string) string {
	switch {
	case strings.HasPrefix(class, "flex") || strings.HasPrefix(class, "grid") ||
		strings.HasPrefix(class, "block") || strings.HasPrefix(class, "inline") ||
		class == "hidden":
		return "display"

	case strings.HasPrefix(class, "items-") || strings.HasPrefix(class, "justify-") ||
		strings.HasPrefix(class, "place-") || strings.HasPrefix(class, "content-"):
		return "alignment"

	case strings.HasPrefix(class, "w-") || strings.HasPrefix(class, "h-") ||
		strings.HasPrefix(class, "min-") || strings.HasPrefix(class, "max-"):
		return "sizing"

	case strings.HasPrefix(class, "p-") || strings.HasPrefix(class, "m-") ||
		strings.HasPrefix(class, "space-") || strings.HasPrefix(class, "gap-"):
		return "spacing"

	case strings.HasPrefix(class, "text-") || strings.HasPrefix(class, "font-") ||
		strings.HasPrefix(class, "leading-") || strings.HasPrefix(class, "tracking-"):
		return "typography"

	case strings.HasPrefix(class, "bg-") || strings.HasPrefix(class, "border-") ||
		strings.HasPrefix(class, "ring-") || strings.HasPrefix(class, "shadow-"):
		return "visual"

	case strings.HasPrefix(class, "rounded") || strings.HasPrefix(class, "opacity-") ||
		strings.HasPrefix(class, "scale-") || strings.HasPrefix(class, "rotate-"):
		return "effects"

	case strings.Contains(class, ":"):
		return "responsive"

	default:
		return "utility"
	}
}
