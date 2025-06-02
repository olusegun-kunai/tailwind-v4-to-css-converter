package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"tailwind-v4-to-css-converter/converter"
)

type AIClient struct {
	apiKey  string
	baseURL string
	model   string
	cache   map[string][]converter.CSSProperty
}

type AIRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AIResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

func NewAIClient(apiKey string) *AIClient {
	return &AIClient{
		apiKey:  apiKey,
		baseURL: "https://api.openai.com/v1/chat/completions", // Default to OpenAI
		model:   "gpt-3.5-turbo",
		cache:   make(map[string][]converter.CSSProperty),
	}
}

func (ai *AIClient) ConvertUnknownClass(className string) ([]converter.CSSProperty, error) {
	// Check cache first
	if cached, exists := ai.cache[className]; exists {
		return cached, nil
	}

	// If no API key provided, return a placeholder
	if ai.apiKey == "" {
		placeholder := []converter.CSSProperty{
			{Name: "/* Unknown Tailwind class */", Value: className},
			{Name: "/* Add manual conversion */", Value: ""},
		}
		ai.cache[className] = placeholder
		return placeholder, nil
	}

	// Query AI for conversion
	properties, err := ai.queryAI(className)
	if err != nil {
		// Fallback to placeholder on error
		placeholder := []converter.CSSProperty{
			{Name: "/* AI conversion failed */", Value: className},
			{Name: "/* Error */", Value: err.Error()},
		}
		ai.cache[className] = placeholder
		return placeholder, nil
	}

	// Cache successful result
	ai.cache[className] = properties
	return properties, nil
}

func (ai *AIClient) queryAI(className string) ([]converter.CSSProperty, error) {
	prompt := fmt.Sprintf(`Convert the Tailwind CSS class "%s" to vanilla CSS properties.

Return the result as a JSON object with this structure:
{
  "properties": [
    {"name": "css-property-name", "value": "css-value"},
    {"name": "another-property", "value": "another-value"}
  ]
}

Only return the JSON object, no other text.`, className)

	request := AIRequest{
		Model: ai.model,
		Messages: []Message{
			{
				Role:    "system",
				Content: "You are a CSS expert that converts Tailwind CSS classes to vanilla CSS properties. Always respond with valid JSON only.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		MaxTokens: 200,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", ai.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ai.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var aiResponse AIResponse
	if err := json.Unmarshal(body, &aiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if len(aiResponse.Choices) == 0 {
		return nil, fmt.Errorf("no response from AI")
	}

	return ai.parseAIResponse(aiResponse.Choices[0].Message.Content)
}

func (ai *AIClient) parseAIResponse(content string) ([]converter.CSSProperty, error) {
	// Clean up the response (remove markdown formatting if present)
	content = strings.TrimSpace(content)
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimSuffix(content, "```")
	content = strings.TrimSpace(content)

	var response struct {
		Properties []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"properties"`
	}

	if err := json.Unmarshal([]byte(content), &response); err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %v", err)
	}

	var properties []converter.CSSProperty
	for _, prop := range response.Properties {
		properties = append(properties, converter.CSSProperty{
			Name:  prop.Name,
			Value: prop.Value,
		})
	}

	return properties, nil
}

func (ai *AIClient) BatchConvert(classNames []string) (map[string][]converter.CSSProperty, error) {
	results := make(map[string][]converter.CSSProperty)

	for _, className := range classNames {
		properties, err := ai.ConvertUnknownClass(className)
		if err != nil {
			// Continue with other classes even if one fails
			results[className] = []converter.CSSProperty{
				{Name: "/* Conversion failed */", Value: className},
			}
			continue
		}
		results[className] = properties
	}

	return results, nil
}

func (ai *AIClient) SearchTailwindDocs(className string) (string, error) {
	// This could be extended to search Tailwind documentation
	// For now, return a placeholder
	return fmt.Sprintf("Search Tailwind docs for: %s", className), nil
}

func (ai *AIClient) SetModel(model string) {
	ai.model = model
}

func (ai *AIClient) SetBaseURL(url string) {
	ai.baseURL = url
}

func (ai *AIClient) ClearCache() {
	ai.cache = make(map[string][]converter.CSSProperty)
}

func (ai *AIClient) GetCacheSize() int {
	return len(ai.cache)
}
