package parser

import (
	"io/ioutil"
	"regexp"
	"strings"
)

type HTMLParser struct {
	classRegex *regexp.Regexp
}

type Document struct {
	Content   string
	ClassRefs []ClassRef
}

type ClassRef struct {
	Classes []string
	Start   int
	End     int
	Element string
}

func NewHTMLParser() *HTMLParser {
	// Regex to match class attributes in HTML/JSX
	classRegex := regexp.MustCompile(`(class|className)=["']([^"']+)["']`)
	return &HTMLParser{
		classRegex: classRegex,
	}
}

func (p *HTMLParser) ParseFile(filepath string) (*Document, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return p.ParseContent(string(content))
}

func (p *HTMLParser) ParseContent(content string) (*Document, error) {
	doc := &Document{
		Content:   content,
		ClassRefs: []ClassRef{},
	}

	// Find all class attributes
	matches := p.classRegex.FindAllStringSubmatchIndex(content, -1)

	for _, match := range matches {
		if len(match) >= 6 {
			start := match[4] // Start of class values
			end := match[5]   // End of class values
			classValues := content[start:end]

			// Split classes by whitespace
			classes := strings.Fields(classValues)

			// Filter for Tailwind classes (simple heuristic)
			tailwindClasses := []string{}
			for _, class := range classes {
				if p.isTailwindClass(class) {
					tailwindClasses = append(tailwindClasses, class)
				}
			}

			if len(tailwindClasses) > 0 {
				// Extract element info for context
				elementStart := p.findElementStart(content, match[0])
				element := p.extractElementName(content, elementStart)

				doc.ClassRefs = append(doc.ClassRefs, ClassRef{
					Classes: tailwindClasses,
					Start:   start,
					End:     end,
					Element: element,
				})
			}
		}
	}

	return doc, nil
}

func (p *HTMLParser) isTailwindClass(class string) bool {
	// Common Tailwind prefixes and patterns
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

func (p *HTMLParser) findElementStart(content string, classPos int) int {
	// Look backwards for the start of the element
	for i := classPos; i >= 0; i-- {
		if content[i] == '<' {
			return i
		}
	}
	return 0
}

func (p *HTMLParser) extractElementName(content string, start int) string {
	// Extract element name from opening tag
	tagMatch := regexp.MustCompile(`<(\w+(?:\.\w+)*)`).FindStringSubmatch(content[start : start+50])
	if len(tagMatch) > 1 {
		return tagMatch[1]
	}
	return "element"
}
