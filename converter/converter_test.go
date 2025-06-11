package converter

import (
	"tailwind-v4-to-css-converter/internal/parser"
	"testing"
)

func TestGenerateSemanticName(t *testing.T) {
	tests := []struct {
		name         string
		filename     string
		element      string
		classes      []parser.ExtractedClass
		expectedName string
	}{
		{
			name:     "OTP component from filename - root div",
			filename: "qwik-otp.tsx",
			element:  "div",
			classes: []parser.ExtractedClass{
				{Name: "flex", Category: "display"},
				{Name: "min-h-screen", Category: "sizing"},
				{Name: "justify-center", Category: "alignment"},
			},
			expectedName: "otp-root",
		},
		{
			name:     "OTP component from filename - input element",
			filename: "qwik-otp.tsx",
			element:  "input",
			classes: []parser.ExtractedClass{
				{Name: "w-12", Category: "sizing"},
				{Name: "h-12", Category: "sizing"},
				{Name: "border-2", Category: "visual"},
			},
			expectedName: "otp-input",
		},
		{
			name:     "OTP component from filename - button element",
			filename: "qwik-otp.tsx",
			element:  "button",
			classes: []parser.ExtractedClass{
				{Name: "bg-blue-600", Category: "visual"},
				{Name: "hover:bg-blue-700", Category: "responsive"},
				{Name: "px-4", Category: "spacing"},
			},
			expectedName: "otp-button",
		},
		{
			name:     "Modal component from filename",
			filename: "react-modal.tsx",
			element:  "div",
			classes: []parser.ExtractedClass{
				{Name: "fixed", Category: "positioning"},
				{Name: "inset-0", Category: "positioning"},
				{Name: "bg-black", Category: "visual"},
			},
			expectedName: "modal-overlay",
		},
		{
			name:         "Component syntax - Otp.Root",
			filename:     "",
			element:      "Otp.Root",
			classes:      []parser.ExtractedClass{},
			expectedName: "otp-root",
		},
		{
			name:         "Component syntax - Accordion.Trigger",
			filename:     "",
			element:      "Accordion.Trigger",
			classes:      []parser.ExtractedClass{},
			expectedName: "accordion-trigger",
		},
		{
			name:     "Generic div without context",
			filename: "",
			element:  "div",
			classes: []parser.ExtractedClass{
				{Name: "flex", Category: "display"},
				{Name: "p-4", Category: "spacing"},
			},
			expectedName: "layout-container",
		},
		{
			name:     "Hero section pattern",
			filename: "hero-section.tsx",
			element:  "div",
			classes: []parser.ExtractedClass{
				{Name: "min-h-screen", Category: "sizing"},
				{Name: "flex", Category: "display"},
				{Name: "items-center", Category: "alignment"},
				{Name: "justify-center", Category: "alignment"},
			},
			expectedName: "section-hero",
		},
		{
			name:     "Card pattern with component context",
			filename: "product-card.tsx",
			element:  "div",
			classes: []parser.ExtractedClass{
				{Name: "rounded", Category: "effects"},
				{Name: "shadow", Category: "visual"},
				{Name: "bg-white", Category: "visual"},
				{Name: "p-6", Category: "spacing"},
			},
			expectedName: "card-card",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conv := NewConverter()
			conv.componentContext = conv.extractComponentFromFilename(tt.filename)

			result := conv.generateSemanticName(tt.element, tt.classes)
			if result != tt.expectedName {
				t.Errorf("generateSemanticName() = %v, want %v", result, tt.expectedName)
			}
		})
	}
}

func TestExtractComponentFromFilename(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected string
	}{
		{
			name:     "Qwik OTP component",
			filename: "qwik-otp.tsx",
			expected: "otp",
		},
		{
			name:     "React Modal component",
			filename: "react-modal.jsx",
			expected: "modal",
		},
		{
			name:     "Vue Accordion component",
			filename: "vue-accordion.vue",
			expected: "accordion",
		},
		{
			name:     "Simple component name",
			filename: "dropdown.tsx",
			expected: "dropdown",
		},
		{
			name:     "Hyphenated component without framework",
			filename: "product-card.tsx",
			expected: "card",
		},
		{
			name:     "Underscore separated",
			filename: "otp_input.tsx",
			expected: "otp",
		},
		{
			name:     "Full path",
			filename: "/src/components/qwik-otp.tsx",
			expected: "otp",
		},
		{
			name:     "Windows path",
			filename: "C:\\src\\components\\react-modal.tsx",
			expected: "modal",
		},
		{
			name:     "Empty filename",
			filename: "",
			expected: "",
		},
		{
			name:     "Complex component name",
			filename: "advanced-data-table.tsx",
			expected: "table",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conv := NewConverter()
			result := conv.extractComponentFromFilename(tt.filename)
			if result != tt.expected {
				t.Errorf("extractComponentFromFilename() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestConvertWithContext(t *testing.T) {
	conv := NewConverter()

	classes := []parser.ExtractedClass{
		{Name: "flex", Category: "display", Context: "div"},
		{Name: "justify-center", Category: "alignment", Context: "div"},
		{Name: "bg-blue-600", Category: "visual", Context: "button"},
		{Name: "hover:bg-blue-700", Category: "responsive", Context: "button"},
	}

	cssRules, semanticMappings := conv.ConvertWithContext(classes, "qwik-otp.tsx")

	// Should generate rules with OTP-prefixed names
	expectedSemanticNames := []string{"otp-root", "otp-button"}
	actualSemanticNames := make([]string, len(semanticMappings))
	for i, mapping := range semanticMappings {
		actualSemanticNames[i] = mapping.SemanticName
	}

	for _, expectedName := range expectedSemanticNames {
		found := false
		for _, actualName := range actualSemanticNames {
			if actualName == expectedName {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected semantic name %s not found in %v", expectedName, actualSemanticNames)
		}
	}

	// Should generate corresponding CSS rules
	if len(cssRules) != len(semanticMappings) {
		t.Errorf("Number of CSS rules (%d) should match semantic mappings (%d)", len(cssRules), len(semanticMappings))
	}
}

func TestDetectSpecificUIPattern(t *testing.T) {
	conv := NewConverter()

	tests := []struct {
		name          string
		allClasses    string
		element       string
		componentName string
		expected      string
	}{
		{
			name:          "Modal overlay pattern",
			allClasses:    "fixed inset-0 bg-black opacity-50",
			element:       "div",
			componentName: "modal",
			expected:      "modal-overlay",
		},
		{
			name:          "Dropdown menu pattern",
			allClasses:    "absolute top-8 shadow bg-white rounded",
			element:       "div",
			componentName: "dropdown",
			expected:      "dropdown-menu",
		},
		{
			name:          "Card pattern",
			allClasses:    "rounded shadow bg-white p-6",
			element:       "div",
			componentName: "product",
			expected:      "product-card",
		},
		{
			name:          "Hero section pattern",
			allClasses:    "min-h-screen flex items-center justify-center",
			element:       "div",
			componentName: "landing",
			expected:      "landing-hero",
		},
		{
			name:          "No pattern match",
			allClasses:    "text-sm text-gray-600",
			element:       "p",
			componentName: "article",
			expected:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := conv.detectSpecificUIPattern(tt.allClasses, tt.element, tt.componentName)
			if result != tt.expected {
				t.Errorf("detectSpecificUIPattern() = %v, want %v", result, tt.expected)
			}
		})
	}
}
