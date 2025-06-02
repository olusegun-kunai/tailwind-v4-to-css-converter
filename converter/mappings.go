package converter

import (
	"regexp"
	"strconv"
)

type TailwindMappings struct {
	staticMappings map[string][]CSSProperty
	dynamicRegex   []*DynamicMapping
}

type DynamicMapping struct {
	Pattern *regexp.Regexp
	Convert func(matches []string) []CSSProperty
}

func NewTailwindMappings() *TailwindMappings {
	tm := &TailwindMappings{
		staticMappings: make(map[string][]CSSProperty),
		dynamicRegex:   []*DynamicMapping{},
	}

	tm.initStaticMappings()
	tm.initDynamicMappings()

	return tm
}

func (tm *TailwindMappings) Convert(class string) []CSSProperty {
	// Try static mappings first
	if props, exists := tm.staticMappings[class]; exists {
		return props
	}

	// Try dynamic mappings
	for _, mapping := range tm.dynamicRegex {
		if matches := mapping.Pattern.FindStringSubmatch(class); matches != nil {
			return mapping.Convert(matches)
		}
	}

	return []CSSProperty{}
}

func (tm *TailwindMappings) initStaticMappings() {
	// Display
	tm.staticMappings["flex"] = []CSSProperty{{Name: "display", Value: "flex"}}
	tm.staticMappings["inline-flex"] = []CSSProperty{{Name: "display", Value: "inline-flex"}}
	tm.staticMappings["grid"] = []CSSProperty{{Name: "display", Value: "grid"}}
	tm.staticMappings["block"] = []CSSProperty{{Name: "display", Value: "block"}}
	tm.staticMappings["inline"] = []CSSProperty{{Name: "display", Value: "inline"}}
	tm.staticMappings["inline-block"] = []CSSProperty{{Name: "display", Value: "inline-block"}}
	tm.staticMappings["hidden"] = []CSSProperty{{Name: "display", Value: "none"}}

	// Flex Direction
	tm.staticMappings["flex-row"] = []CSSProperty{{Name: "flex-direction", Value: "row"}}
	tm.staticMappings["flex-col"] = []CSSProperty{{Name: "flex-direction", Value: "column"}}
	tm.staticMappings["flex-row-reverse"] = []CSSProperty{{Name: "flex-direction", Value: "row-reverse"}}
	tm.staticMappings["flex-col-reverse"] = []CSSProperty{{Name: "flex-direction", Value: "column-reverse"}}

	// Alignment
	tm.staticMappings["items-start"] = []CSSProperty{{Name: "align-items", Value: "flex-start"}}
	tm.staticMappings["items-center"] = []CSSProperty{{Name: "align-items", Value: "center"}}
	tm.staticMappings["items-end"] = []CSSProperty{{Name: "align-items", Value: "flex-end"}}
	tm.staticMappings["items-stretch"] = []CSSProperty{{Name: "align-items", Value: "stretch"}}
	tm.staticMappings["items-baseline"] = []CSSProperty{{Name: "align-items", Value: "baseline"}}

	tm.staticMappings["justify-start"] = []CSSProperty{{Name: "justify-content", Value: "flex-start"}}
	tm.staticMappings["justify-center"] = []CSSProperty{{Name: "justify-content", Value: "center"}}
	tm.staticMappings["justify-end"] = []CSSProperty{{Name: "justify-content", Value: "flex-end"}}
	tm.staticMappings["justify-between"] = []CSSProperty{{Name: "justify-content", Value: "space-between"}}
	tm.staticMappings["justify-around"] = []CSSProperty{{Name: "justify-content", Value: "space-around"}}
	tm.staticMappings["justify-evenly"] = []CSSProperty{{Name: "justify-content", Value: "space-evenly"}}

	// Text Alignment
	tm.staticMappings["text-left"] = []CSSProperty{{Name: "text-align", Value: "left"}}
	tm.staticMappings["text-center"] = []CSSProperty{{Name: "text-align", Value: "center"}}
	tm.staticMappings["text-right"] = []CSSProperty{{Name: "text-align", Value: "right"}}
	tm.staticMappings["text-justify"] = []CSSProperty{{Name: "text-align", Value: "justify"}}

	// Font Weight
	tm.staticMappings["font-thin"] = []CSSProperty{{Name: "font-weight", Value: "100"}}
	tm.staticMappings["font-light"] = []CSSProperty{{Name: "font-weight", Value: "300"}}
	tm.staticMappings["font-normal"] = []CSSProperty{{Name: "font-weight", Value: "400"}}
	tm.staticMappings["font-medium"] = []CSSProperty{{Name: "font-weight", Value: "500"}}
	tm.staticMappings["font-semibold"] = []CSSProperty{{Name: "font-weight", Value: "600"}}
	tm.staticMappings["font-bold"] = []CSSProperty{{Name: "font-weight", Value: "700"}}
	tm.staticMappings["font-extrabold"] = []CSSProperty{{Name: "font-weight", Value: "800"}}
	tm.staticMappings["font-black"] = []CSSProperty{{Name: "font-weight", Value: "900"}}

	// Position
	tm.staticMappings["static"] = []CSSProperty{{Name: "position", Value: "static"}}
	tm.staticMappings["relative"] = []CSSProperty{{Name: "position", Value: "relative"}}
	tm.staticMappings["absolute"] = []CSSProperty{{Name: "position", Value: "absolute"}}
	tm.staticMappings["fixed"] = []CSSProperty{{Name: "position", Value: "fixed"}}
	tm.staticMappings["sticky"] = []CSSProperty{{Name: "position", Value: "sticky"}}

	// Common Layout Classes
	tm.staticMappings["container"] = []CSSProperty{{Name: "max-width", Value: "1200px"}, {Name: "margin", Value: "0 auto"}}
	tm.staticMappings["mx-auto"] = []CSSProperty{{Name: "margin-left", Value: "auto"}, {Name: "margin-right", Value: "auto"}}

	// Common Visual Classes
	tm.staticMappings["bg-white"] = []CSSProperty{{Name: "background-color", Value: "#ffffff"}}
	tm.staticMappings["bg-black"] = []CSSProperty{{Name: "background-color", Value: "#000000"}}
	tm.staticMappings["text-white"] = []CSSProperty{{Name: "color", Value: "#ffffff"}}
	tm.staticMappings["text-black"] = []CSSProperty{{Name: "color", Value: "#000000"}}

	// Border
	tm.staticMappings["border"] = []CSSProperty{{Name: "border", Value: "1px solid #e5e7eb"}}
	tm.staticMappings["border-b"] = []CSSProperty{{Name: "border-bottom", Value: "1px solid #e5e7eb"}}
	tm.staticMappings["border-t"] = []CSSProperty{{Name: "border-top", Value: "1px solid #e5e7eb"}}
	tm.staticMappings["border-l"] = []CSSProperty{{Name: "border-left", Value: "1px solid #e5e7eb"}}
	tm.staticMappings["border-r"] = []CSSProperty{{Name: "border-right", Value: "1px solid #e5e7eb"}}

	// Rounded
	tm.staticMappings["rounded"] = []CSSProperty{{Name: "border-radius", Value: "0.25rem"}}
	tm.staticMappings["rounded-md"] = []CSSProperty{{Name: "border-radius", Value: "0.375rem"}}
	tm.staticMappings["rounded-lg"] = []CSSProperty{{Name: "border-radius", Value: "0.5rem"}}
	tm.staticMappings["rounded-xl"] = []CSSProperty{{Name: "border-radius", Value: "0.75rem"}}
	tm.staticMappings["rounded-2xl"] = []CSSProperty{{Name: "border-radius", Value: "1rem"}}
	tm.staticMappings["rounded-full"] = []CSSProperty{{Name: "border-radius", Value: "9999px"}}

	// Shadow
	tm.staticMappings["shadow"] = []CSSProperty{{Name: "box-shadow", Value: "0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)"}}
	tm.staticMappings["shadow-md"] = []CSSProperty{{Name: "box-shadow", Value: "0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)"}}
	tm.staticMappings["shadow-lg"] = []CSSProperty{{Name: "box-shadow", Value: "0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)"}}

	// Transition
	tm.staticMappings["transition-shadow"] = []CSSProperty{{Name: "transition", Value: "box-shadow 150ms ease-in-out"}}
	tm.staticMappings["transition-colors"] = []CSSProperty{{Name: "transition", Value: "color, background-color, border-color 150ms ease-in-out"}}

	// Grid
	tm.staticMappings["grid-cols-1"] = []CSSProperty{{Name: "grid-template-columns", Value: "repeat(1, minmax(0, 1fr))"}}
	tm.staticMappings["grid-cols-2"] = []CSSProperty{{Name: "grid-template-columns", Value: "repeat(2, minmax(0, 1fr))"}}
	tm.staticMappings["grid-cols-3"] = []CSSProperty{{Name: "grid-template-columns", Value: "repeat(3, minmax(0, 1fr))"}}
	tm.staticMappings["grid-cols-4"] = []CSSProperty{{Name: "grid-template-columns", Value: "repeat(4, minmax(0, 1fr))"}}

	// Focus states
	tm.staticMappings["focus:outline-none"] = []CSSProperty{{Name: "outline", Value: "none"}}
	tm.staticMappings["focus:ring-2"] = []CSSProperty{{Name: "box-shadow", Value: "0 0 0 2px rgba(59, 130, 246, 0.5)"}}
}

func (tm *TailwindMappings) initDynamicMappings() {
	// Gap
	tm.dynamicRegex = append(tm.dynamicRegex, &DynamicMapping{
		Pattern: regexp.MustCompile(`^gap-(\d+(?:\.\d+)?)$`),
		Convert: func(matches []string) []CSSProperty {
			value := tm.convertSpacing(matches[1])
			return []CSSProperty{{Name: "gap", Value: value}}
		},
	})

	// Padding
	tm.dynamicRegex = append(tm.dynamicRegex, &DynamicMapping{
		Pattern: regexp.MustCompile(`^p-(\d+(?:\.\d+)?)$`),
		Convert: func(matches []string) []CSSProperty {
			value := tm.convertSpacing(matches[1])
			return []CSSProperty{{Name: "padding", Value: value}}
		},
	})

	tm.dynamicRegex = append(tm.dynamicRegex, &DynamicMapping{
		Pattern: regexp.MustCompile(`^p([xytrbl])-(\d+(?:\.\d+)?)$`),
		Convert: func(matches []string) []CSSProperty {
			direction := matches[1]
			value := tm.convertSpacing(matches[2])
			return tm.getPaddingProperties(direction, value)
		},
	})

	// Margin
	tm.dynamicRegex = append(tm.dynamicRegex, &DynamicMapping{
		Pattern: regexp.MustCompile(`^m-(\d+(?:\.\d+)?)$`),
		Convert: func(matches []string) []CSSProperty {
			value := tm.convertSpacing(matches[1])
			return []CSSProperty{{Name: "margin", Value: value}}
		},
	})

	tm.dynamicRegex = append(tm.dynamicRegex, &DynamicMapping{
		Pattern: regexp.MustCompile(`^m([xytrbl])-(\d+(?:\.\d+)?)$`),
		Convert: func(matches []string) []CSSProperty {
			direction := matches[1]
			value := tm.convertSpacing(matches[2])
			return tm.getMarginProperties(direction, value)
		},
	})

	// Width
	tm.dynamicRegex = append(tm.dynamicRegex, &DynamicMapping{
		Pattern: regexp.MustCompile(`^w-(\d+(?:\.\d+)?)$`),
		Convert: func(matches []string) []CSSProperty {
			value := tm.convertSize(matches[1])
			return []CSSProperty{{Name: "width", Value: value}}
		},
	})

	// Height
	tm.dynamicRegex = append(tm.dynamicRegex, &DynamicMapping{
		Pattern: regexp.MustCompile(`^h-(\d+(?:\.\d+)?)$`),
		Convert: func(matches []string) []CSSProperty {
			value := tm.convertSize(matches[1])
			return []CSSProperty{{Name: "height", Value: value}}
		},
	})

	// Text Size
	tm.dynamicRegex = append(tm.dynamicRegex, &DynamicMapping{
		Pattern: regexp.MustCompile(`^text-(xs|sm|base|lg|xl|2xl|3xl|4xl|5xl|6xl|7xl|8xl|9xl)$`),
		Convert: func(matches []string) []CSSProperty {
			size := tm.getTextSize(matches[1])
			return []CSSProperty{{Name: "font-size", Value: size}}
		},
	})

	// Background Colors (including simple colors like bg-white, bg-red)
	tm.dynamicRegex = append(tm.dynamicRegex, &DynamicMapping{
		Pattern: regexp.MustCompile(`^bg-(\w+)-(\d+)$`),
		Convert: func(matches []string) []CSSProperty {
			color := tm.getColor(matches[1], matches[2])
			return []CSSProperty{{Name: "background-color", Value: color}}
		},
	})

	// Text Colors
	tm.dynamicRegex = append(tm.dynamicRegex, &DynamicMapping{
		Pattern: regexp.MustCompile(`^text-(\w+)-(\d+)$`),
		Convert: func(matches []string) []CSSProperty {
			color := tm.getColor(matches[1], matches[2])
			return []CSSProperty{{Name: "color", Value: color}}
		},
	})

	// Border Colors
	tm.dynamicRegex = append(tm.dynamicRegex, &DynamicMapping{
		Pattern: regexp.MustCompile(`^border-(\w+)-(\d+)$`),
		Convert: func(matches []string) []CSSProperty {
			color := tm.getColor(matches[1], matches[2])
			return []CSSProperty{{Name: "border-color", Value: color}}
		},
	})

	// Grid responsive classes (md:, lg:)
	tm.dynamicRegex = append(tm.dynamicRegex, &DynamicMapping{
		Pattern: regexp.MustCompile(`^(md|lg):grid-cols-(\d+)$`),
		Convert: func(matches []string) []CSSProperty {
			breakpoint := matches[1]
			cols := matches[2]
			mediaQuery := tm.getMediaQuery(breakpoint)
			return []CSSProperty{{Name: "@media " + mediaQuery, Value: "grid-template-columns: repeat(" + cols + ", minmax(0, 1fr))"}}
		},
	})

	// Hover states
	tm.dynamicRegex = append(tm.dynamicRegex, &DynamicMapping{
		Pattern: regexp.MustCompile(`^hover:(\w+)-(\w+)-(\d+)$`),
		Convert: func(matches []string) []CSSProperty {
			property := matches[1]
			colorName := matches[2]
			shade := matches[3]
			color := tm.getColor(colorName, shade)

			var cssProperty string
			switch property {
			case "bg":
				cssProperty = "background-color"
			case "text":
				cssProperty = "color"
			case "border":
				cssProperty = "border-color"
			default:
				cssProperty = "color"
			}

			return []CSSProperty{{Name: ":hover", Value: cssProperty + ": " + color}}
		},
	})

	// Focus ring colors
	tm.dynamicRegex = append(tm.dynamicRegex, &DynamicMapping{
		Pattern: regexp.MustCompile(`^focus:ring-(\w+)-(\d+)$`),
		Convert: func(matches []string) []CSSProperty {
			color := tm.getColor(matches[1], matches[2])
			return []CSSProperty{{Name: "box-shadow", Value: "0 0 0 2px " + color}}
		},
	})
}

func (tm *TailwindMappings) convertSpacing(value string) string {
	// Convert Tailwind spacing to rem
	if val, err := strconv.ParseFloat(value, 64); err == nil {
		return strconv.FormatFloat(val*0.25, 'f', -1, 64) + "rem"
	}
	return value
}

func (tm *TailwindMappings) convertSize(value string) string {
	// Convert Tailwind size to rem
	if val, err := strconv.ParseFloat(value, 64); err == nil {
		return strconv.FormatFloat(val*0.25, 'f', -1, 64) + "rem"
	}
	return value
}

func (tm *TailwindMappings) getPaddingProperties(direction, value string) []CSSProperty {
	switch direction {
	case "x":
		return []CSSProperty{
			{Name: "padding-left", Value: value},
			{Name: "padding-right", Value: value},
		}
	case "y":
		return []CSSProperty{
			{Name: "padding-top", Value: value},
			{Name: "padding-bottom", Value: value},
		}
	case "t":
		return []CSSProperty{{Name: "padding-top", Value: value}}
	case "r":
		return []CSSProperty{{Name: "padding-right", Value: value}}
	case "b":
		return []CSSProperty{{Name: "padding-bottom", Value: value}}
	case "l":
		return []CSSProperty{{Name: "padding-left", Value: value}}
	}
	return []CSSProperty{}
}

func (tm *TailwindMappings) getMarginProperties(direction, value string) []CSSProperty {
	switch direction {
	case "x":
		return []CSSProperty{
			{Name: "margin-left", Value: value},
			{Name: "margin-right", Value: value},
		}
	case "y":
		return []CSSProperty{
			{Name: "margin-top", Value: value},
			{Name: "margin-bottom", Value: value},
		}
	case "t":
		return []CSSProperty{{Name: "margin-top", Value: value}}
	case "r":
		return []CSSProperty{{Name: "margin-right", Value: value}}
	case "b":
		return []CSSProperty{{Name: "margin-bottom", Value: value}}
	case "l":
		return []CSSProperty{{Name: "margin-left", Value: value}}
	}
	return []CSSProperty{}
}

func (tm *TailwindMappings) getTextSize(size string) string {
	sizes := map[string]string{
		"xs":   "0.75rem",
		"sm":   "0.875rem",
		"base": "1rem",
		"lg":   "1.125rem",
		"xl":   "1.25rem",
		"2xl":  "1.5rem",
		"3xl":  "1.875rem",
		"4xl":  "2.25rem",
		"5xl":  "3rem",
		"6xl":  "3.75rem",
		"7xl":  "4.5rem",
		"8xl":  "6rem",
		"9xl":  "8rem",
	}

	if val, exists := sizes[size]; exists {
		return val
	}
	return "1rem"
}

func (tm *TailwindMappings) getColor(colorName, shade string) string {
	// Comprehensive color mapping
	colors := map[string]map[string]string{
		"blue": {
			"50":  "#eff6ff",
			"100": "#dbeafe",
			"200": "#bfdbfe",
			"300": "#93c5fd",
			"400": "#60a5fa",
			"500": "#3b82f6",
			"600": "#2563eb",
			"700": "#1d4ed8",
			"800": "#1e40af",
			"900": "#1e3a8a",
		},
		"red": {
			"50":  "#fef2f2",
			"100": "#fee2e2",
			"200": "#fecaca",
			"300": "#fca5a5",
			"400": "#f87171",
			"500": "#ef4444",
			"600": "#dc2626",
			"700": "#b91c1c",
			"800": "#991b1b",
			"900": "#7f1d1d",
		},
		"green": {
			"50":  "#f0fdf4",
			"100": "#dcfce7",
			"200": "#bbf7d0",
			"300": "#86efac",
			"400": "#4ade80",
			"500": "#22c55e",
			"600": "#16a34a",
			"700": "#15803d",
			"800": "#166534",
			"900": "#14532d",
		},
		"gray": {
			"50":  "#f9fafb",
			"100": "#f3f4f6",
			"200": "#e5e7eb",
			"300": "#d1d5db",
			"400": "#9ca3af",
			"500": "#6b7280",
			"600": "#4b5563",
			"700": "#374151",
			"800": "#1f2937",
			"900": "#111827",
		},
		"purple": {
			"50":  "#faf5ff",
			"100": "#f3e8ff",
			"200": "#e9d5ff",
			"300": "#d8b4fe",
			"400": "#c084fc",
			"500": "#a855f7",
			"600": "#9333ea",
			"700": "#7c3aed",
			"800": "#6b21a8",
			"900": "#581c87",
		},
	}

	if colorMap, exists := colors[colorName]; exists {
		if color, exists := colorMap[shade]; exists {
			return color
		}
	}

	return "#000000" // fallback
}

func (tm *TailwindMappings) getMediaQuery(breakpoint string) string {
	breakpoints := map[string]string{
		"sm":  "(min-width: 640px)",
		"md":  "(min-width: 768px)",
		"lg":  "(min-width: 1024px)",
		"xl":  "(min-width: 1280px)",
		"2xl": "(min-width: 1536px)",
	}

	if query, exists := breakpoints[breakpoint]; exists {
		return query
	}
	return "(min-width: 768px)" // fallback to md
}
