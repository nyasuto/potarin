package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// JSONSchemaConfig represents a schema configuration for structured outputs.
type JSONSchemaConfig struct {
	Name   string          `json:"name,omitempty"`
	Schema json.RawMessage `json:"schema,omitempty"`
	Strict bool            `json:"strict,omitempty"`
}

// ResponseFormat specifies how the model should format its response.
type ResponseFormat struct {
	Type       string            `json:"type"`
	JSONSchema *JSONSchemaConfig `json:"json_schema,omitempty"`
}

type ChatRequest struct {
	Model          string         `json:"model"`
	Messages       []Message      `json:"messages"`
	ResponseFormat ResponseFormat `json:"response_format"`
}

type chatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

const fallbackModel = "gpt-4o-mini"

// GetModel returns the chat model defined in the environment or a default.
func GetModel() string {
	m := os.Getenv("OPENAI_MODEL")
	if m == "" {
		return fallbackModel
	}
	return m
}

// LoadEnv loads environment variables from .env.local if present.
func LoadEnv() {
	if err := godotenv.Load(".env.local"); err != nil {
		if !os.IsNotExist(err) {
			// Log or print the error if it's not a "file not found" error
			fmt.Printf("Error loading .env.local: %v\n", err)
		}
	}
}

// NewClient creates a new OpenAI client using environment variables.
// It returns an error if the API key is empty.
func NewClient() (*Client, error) {
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY is not set")
	}
	return &Client{
		apiKey:     key,
		baseURL:    "https://api.openai.com",
		httpClient: &http.Client{},
	}, nil
}

// Chat sends a chat completion request and returns the first message content.
func (c *Client) Chat(ctx context.Context, req ChatRequest) (string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/v1/chat/completions", bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(string(body))
	}
	var res chatResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return "", err
	}
	if len(res.Choices) == 0 {
		return "", errors.New("no choices returned")
	}
	return res.Choices[0].Message.Content, nil
}
