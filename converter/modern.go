package converter

import (
	"regexp"
	"strings"
)

type ModernFeatures struct {
	containerQueries map[string][]CSSProperty
	cascadeLayers    map[string][]CSSProperty
	customProperties map[string][]CSSProperty
}

func NewModernFeatures() *ModernFeatures {
	mf := &ModernFeatures{
		containerQueries: make(map[string][]CSSProperty),
		cascadeLayers:    make(map[string][]CSSProperty),
		customProperties: make(map[string][]CSSProperty),
	}

	mf.initModernFeatures()
	return mf
}

func (mf *ModernFeatures) Convert(class string) []CSSProperty {
	// Handle container queries
	if strings.HasPrefix(class, "@container") {
		return mf.convertContainerQuery(class)
	}

	// Handle cascade layers
	if strings.HasPrefix(class, "@layer") {
		return mf.convertCascadeLayer(class)
	}

	// Handle custom properties (CSS variables)
	if strings.Contains(class, "--") {
		return mf.convertCustomProperty(class)
	}

	// Handle modern pseudo-classes
	if strings.Contains(class, ":") {
		return mf.convertModernPseudo(class)
	}

	// Handle new Tailwind v4+ utilities
	return mf.convertV4Utilities(class)
}

func (mf *ModernFeatures) initModernFeatures() {
	// Initialize container query mappings
	mf.containerQueries["@container-sm"] = []CSSProperty{
		{Name: "@container", Value: "(min-width: 640px)"},
	}
	mf.containerQueries["@container-md"] = []CSSProperty{
		{Name: "@container", Value: "(min-width: 768px)"},
	}
	mf.containerQueries["@container-lg"] = []CSSProperty{
		{Name: "@container", Value: "(min-width: 1024px)"},
	}

	// Initialize cascade layer mappings
	mf.cascadeLayers["@layer-base"] = []CSSProperty{
		{Name: "@layer", Value: "base"},
	}
	mf.cascadeLayers["@layer-components"] = []CSSProperty{
		{Name: "@layer", Value: "components"},
	}
	mf.cascadeLayers["@layer-utilities"] = []CSSProperty{
		{Name: "@layer", Value: "utilities"},
	}
}

func (mf *ModernFeatures) convertContainerQuery(class string) []CSSProperty {
	// Extract container query logic
	if strings.HasPrefix(class, "@container") {
		// Parse container query syntax
		parts := strings.Split(class, ":")
		if len(parts) > 1 {
			query := parts[0]
			property := parts[1]

			return []CSSProperty{
				{Name: "/* Container Query */", Value: query},
				{Name: "/* Applied when */", Value: property},
			}
		}
	}

	return []CSSProperty{}
}

func (mf *ModernFeatures) convertCascadeLayer(class string) []CSSProperty {
	// Handle cascade layers
	if layer, exists := mf.cascadeLayers[class]; exists {
		return layer
	}

	return []CSSProperty{
		{Name: "/* Cascade Layer */", Value: class},
	}
}

func (mf *ModernFeatures) convertCustomProperty(class string) []CSSProperty {
	// Handle CSS custom properties
	if strings.Contains(class, "--") {
		parts := strings.Split(class, "-")
		if len(parts) >= 3 {
			varName := strings.Join(parts[2:], "-")
			return []CSSProperty{
				{Name: "--" + varName, Value: "/* Custom property value */"},
			}
		}
	}

	return []CSSProperty{}
}

func (mf *ModernFeatures) convertModernPseudo(class string) []CSSProperty {
	// Handle modern pseudo-classes and variants
	parts := strings.Split(class, ":")
	if len(parts) < 2 {
		return []CSSProperty{}
	}

	variant := parts[0]
	baseClass := strings.Join(parts[1:], ":")

	switch variant {
	case "hover":
		return []CSSProperty{
			{Name: "/* On hover */", Value: baseClass},
		}
	case "focus":
		return []CSSProperty{
			{Name: "/* On focus */", Value: baseClass},
		}
	case "active":
		return []CSSProperty{
			{Name: "/* When active */", Value: baseClass},
		}
	case "disabled":
		return []CSSProperty{
			{Name: "/* When disabled */", Value: baseClass},
		}
	case "group-hover":
		return []CSSProperty{
			{Name: "/* When parent group hovered */", Value: baseClass},
		}
	case "peer-focus":
		return []CSSProperty{
			{Name: "/* When peer focused */", Value: baseClass},
		}
	default:
		// Handle responsive variants
		if mf.isResponsiveVariant(variant) {
			return []CSSProperty{
				{Name: "/* Responsive: " + variant + " */", Value: baseClass},
			}
		}
	}

	return []CSSProperty{
		{Name: "/* Unknown variant: " + variant + " */", Value: baseClass},
	}
}

func (mf *ModernFeatures) convertV4Utilities(class string) []CSSProperty {
	// Handle new Tailwind v4+ utilities
	v4Patterns := map[string]func(string) []CSSProperty{
		"size-":    mf.convertSizeUtility,
		"grid-":    mf.convertGridUtility,
		"place-":   mf.convertPlaceUtility,
		"content-": mf.convertContentUtility,
	}

	for prefix, converter := range v4Patterns {
		if strings.HasPrefix(class, prefix) {
			return converter(class)
		}
	}

	return []CSSProperty{}
}

func (mf *ModernFeatures) convertSizeUtility(class string) []CSSProperty {
	// Handle size-* utilities (sets both width and height)
	re := regexp.MustCompile(`^size-(\d+(?:\.\d+)?)$`)
	if matches := re.FindStringSubmatch(class); matches != nil {
		value := matches[1] + "rem"
		return []CSSProperty{
			{Name: "width", Value: value},
			{Name: "height", Value: value},
		}
	}
	return []CSSProperty{}
}

func (mf *ModernFeatures) convertGridUtility(class string) []CSSProperty {
	// Handle advanced grid utilities
	switch {
	case strings.HasPrefix(class, "grid-cols-"):
		re := regexp.MustCompile(`^grid-cols-(\d+)$`)
		if matches := re.FindStringSubmatch(class); matches != nil {
			return []CSSProperty{
				{Name: "grid-template-columns", Value: "repeat(" + matches[1] + ", minmax(0, 1fr))"},
			}
		}
	case strings.HasPrefix(class, "grid-rows-"):
		re := regexp.MustCompile(`^grid-rows-(\d+)$`)
		if matches := re.FindStringSubmatch(class); matches != nil {
			return []CSSProperty{
				{Name: "grid-template-rows", Value: "repeat(" + matches[1] + ", minmax(0, 1fr))"},
			}
		}
	}
	return []CSSProperty{}
}

func (mf *ModernFeatures) convertPlaceUtility(class string) []CSSProperty {
	// Handle place-* utilities
	switch class {
	case "place-content-center":
		return []CSSProperty{{Name: "place-content", Value: "center"}}
	case "place-content-start":
		return []CSSProperty{{Name: "place-content", Value: "start"}}
	case "place-content-end":
		return []CSSProperty{{Name: "place-content", Value: "end"}}
	case "place-items-center":
		return []CSSProperty{{Name: "place-items", Value: "center"}}
	case "place-items-start":
		return []CSSProperty{{Name: "place-items", Value: "start"}}
	case "place-items-end":
		return []CSSProperty{{Name: "place-items", Value: "end"}}
	}
	return []CSSProperty{}
}

func (mf *ModernFeatures) convertContentUtility(class string) []CSSProperty {
	// Handle content-* utilities
	switch class {
	case "content-center":
		return []CSSProperty{{Name: "align-content", Value: "center"}}
	case "content-start":
		return []CSSProperty{{Name: "align-content", Value: "flex-start"}}
	case "content-end":
		return []CSSProperty{{Name: "align-content", Value: "flex-end"}}
	case "content-between":
		return []CSSProperty{{Name: "align-content", Value: "space-between"}}
	case "content-around":
		return []CSSProperty{{Name: "align-content", Value: "space-around"}}
	case "content-evenly":
		return []CSSProperty{{Name: "align-content", Value: "space-evenly"}}
	}
	return []CSSProperty{}
}

func (mf *ModernFeatures) isResponsiveVariant(variant string) bool {
	responsiveVariants := []string{"sm", "md", "lg", "xl", "2xl"}
	for _, rv := range responsiveVariants {
		if variant == rv {
			return true
		}
	}
	return false
}
