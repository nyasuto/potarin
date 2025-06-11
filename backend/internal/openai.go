package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
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

type ResponseFormat struct {
	Type string `json:"type"`
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

// LoadEnv loads environment variables from .env.local if present.
func LoadEnv() {
	_ = godotenv.Load(".env.local")
}

// NewClient creates a new OpenAI client using environment variables.
func NewClient() *Client {
	return &Client{
		apiKey:     os.Getenv("OPENAI_API_KEY"),
		baseURL:    "https://api.openai.com",
		httpClient: &http.Client{},
	}
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
