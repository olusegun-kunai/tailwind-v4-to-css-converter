package generator

import (
	"os"
	"strings"
	"tailwind-v4-to-css-converter/converter"
)

type CSSGenerator struct{}

func NewCSSGenerator() *CSSGenerator {
	return &CSSGenerator{}
}

func (g *CSSGenerator) Generate(rules []converter.CSSRule, outputPath string) error {
	var cssContent strings.Builder

	// Add header comment
	cssContent.WriteString("/* Generated CSS Module */\n")
	cssContent.WriteString("/* Converted from Tailwind CSS classes */\n\n")

	// Generate CSS rules
	for _, rule := range rules {
		g.writeRule(&cssContent, rule)
		cssContent.WriteString("\n")
	}

	// Write to file
	return os.WriteFile(outputPath, []byte(cssContent.String()), 0644)
}

func (g *CSSGenerator) writeRule(builder *strings.Builder, rule converter.CSSRule) {
	// Write selector
	builder.WriteString(rule.Selector)
	builder.WriteString(" {\n")

	// Separate properties into regular properties, media queries, hover states, and comments
	var regularProps []converter.CSSProperty
	var mediaQueries []converter.CSSProperty
	var hoverStates []converter.CSSProperty
	var comments []converter.CSSProperty

	for _, prop := range rule.Properties {
		if strings.HasPrefix(prop.Name, "/*") {
			comments = append(comments, prop)
		} else if strings.HasPrefix(prop.Name, "@media") {
			mediaQueries = append(mediaQueries, prop)
		} else if strings.HasPrefix(prop.Name, ":hover") {
			hoverStates = append(hoverStates, prop)
		} else if strings.HasPrefix(prop.Name, "@") {
			// Other at-rules, write as comments for now
			comments = append(comments, converter.CSSProperty{
				Name:  "/* " + prop.Name + " */",
				Value: prop.Value,
			})
		} else {
			regularProps = append(regularProps, prop)
		}
	}

	// Write regular properties
	for _, prop := range regularProps {
		builder.WriteString("  ")
		builder.WriteString(prop.Name)
		builder.WriteString(": ")
		builder.WriteString(prop.Value)
		builder.WriteString(";\n")
	}

	// Write comments
	for _, comment := range comments {
		builder.WriteString("  ")
		builder.WriteString(comment.Name)
		if comment.Value != "" {
			builder.WriteString(" ")
			builder.WriteString(comment.Value)
		}
		builder.WriteString("\n")
	}

	builder.WriteString("}")

	// Write hover states as separate rules
	if len(hoverStates) > 0 {
		builder.WriteString("\n\n")
		builder.WriteString(rule.Selector)
		builder.WriteString(":hover {\n")
		for _, prop := range hoverStates {
			builder.WriteString("  ")
			builder.WriteString(prop.Value)
			builder.WriteString(";\n")
		}
		builder.WriteString("}")
	}

	// Write media queries as separate rules
	for _, mediaQuery := range mediaQueries {
		builder.WriteString("\n\n")
		builder.WriteString(mediaQuery.Name)
		builder.WriteString(" {\n")
		builder.WriteString("  ")
		builder.WriteString(rule.Selector)
		builder.WriteString(" {\n")
		builder.WriteString("    ")
		builder.WriteString(mediaQuery.Value)
		builder.WriteString(";\n")
		builder.WriteString("  }\n")
		builder.WriteString("}")
	}
}

func (g *CSSGenerator) writeAtRule(builder *strings.Builder, prop converter.CSSProperty, rule converter.CSSRule) {
	switch {
	case strings.HasPrefix(prop.Name, "@media"):
		g.writeMediaQuery(builder, prop, rule)
	case strings.HasPrefix(prop.Name, "@container"):
		g.writeContainerQuery(builder, prop, rule)
	case strings.HasPrefix(prop.Name, "@layer"):
		g.writeCascadeLayer(builder, prop, rule)
	default:
		// Unknown at-rule, write as comment
		builder.WriteString("  /* ")
		builder.WriteString(prop.Name)
		builder.WriteString(": ")
		builder.WriteString(prop.Value)
		builder.WriteString(" */\n")
	}
}

func (g *CSSGenerator) writeMediaQuery(builder *strings.Builder, prop converter.CSSProperty, rule converter.CSSRule) {
	builder.WriteString("  ")
	builder.WriteString(prop.Name)
	builder.WriteString(" ")
	builder.WriteString(prop.Value)
	builder.WriteString(" {\n")

	// Write properties inside media query
	builder.WriteString("    /* Media query content would go here */\n")

	builder.WriteString("  }\n")
}

func (g *CSSGenerator) writeContainerQuery(builder *strings.Builder, prop converter.CSSProperty, rule converter.CSSRule) {
	builder.WriteString("  @container ")
	builder.WriteString(prop.Value)
	builder.WriteString(" {\n")

	// Write properties inside container query
	builder.WriteString("    /* Container query content would go here */\n")

	builder.WriteString("  }\n")
}

func (g *CSSGenerator) writeCascadeLayer(builder *strings.Builder, prop converter.CSSProperty, rule converter.CSSRule) {
	builder.WriteString("  /* @layer ")
	builder.WriteString(prop.Value)
	builder.WriteString(" */\n")
}

func (g *CSSGenerator) GenerateWithOptions(rules []converter.CSSRule, outputPath string, options CSSOptions) error {
	var cssContent strings.Builder

	// Add custom header if provided
	if options.Header != "" {
		cssContent.WriteString(options.Header)
		cssContent.WriteString("\n\n")
	}

	// Add imports if provided
	for _, importStmt := range options.Imports {
		cssContent.WriteString("@import \"")
		cssContent.WriteString(importStmt)
		cssContent.WriteString("\";\n")
	}

	if len(options.Imports) > 0 {
		cssContent.WriteString("\n")
	}

	// Generate CSS rules with formatting options
	for i, rule := range rules {
		if options.Minify {
			g.writeMinifiedRule(&cssContent, rule)
		} else {
			g.writeFormattedRule(&cssContent, rule, options.IndentSize)
		}

		// Add spacing between rules
		if !options.Minify && i < len(rules)-1 {
			cssContent.WriteString("\n")
		}
	}

	// Write to file
	return os.WriteFile(outputPath, []byte(cssContent.String()), 0644)
}

func (g *CSSGenerator) writeFormattedRule(builder *strings.Builder, rule converter.CSSRule, indentSize int) {
	indent := strings.Repeat(" ", indentSize)

	builder.WriteString(rule.Selector)
	builder.WriteString(" {\n")

	for _, prop := range rule.Properties {
		if !strings.HasPrefix(prop.Name, "/*") && !strings.HasPrefix(prop.Name, "@") {
			builder.WriteString(indent)
			builder.WriteString(prop.Name)
			builder.WriteString(": ")
			builder.WriteString(prop.Value)
			builder.WriteString(";\n")
		}
	}

	builder.WriteString("}")
	builder.WriteString("\n")
}

func (g *CSSGenerator) writeMinifiedRule(builder *strings.Builder, rule converter.CSSRule) {
	builder.WriteString(rule.Selector)
	builder.WriteString("{")

	for _, prop := range rule.Properties {
		if !strings.HasPrefix(prop.Name, "/*") && !strings.HasPrefix(prop.Name, "@") {
			builder.WriteString(prop.Name)
			builder.WriteString(":")
			builder.WriteString(prop.Value)
			builder.WriteString(";")
		}
	}

	builder.WriteString("}")
}

type CSSOptions struct {
	Header     string
	Imports    []string
	Minify     bool
	IndentSize int
}

func DefaultCSSOptions() CSSOptions {
	return CSSOptions{
		Header:     "/* Generated CSS Module */\n/* Converted from Tailwind CSS classes */",
		Imports:    []string{},
		Minify:     false,
		IndentSize: 2,
	}
}
