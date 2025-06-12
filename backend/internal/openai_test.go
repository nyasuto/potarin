package internal

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mockRoundTripper not needed because we use httptest server to set url and client

func TestChatSuccess(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"choices":[{"message":{"role":"assistant","content":"hello"}}]}`)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	c := &Client{
		apiKey:     "test",
		baseURL:    server.URL,
		httpClient: server.Client(),
	}

	content, err := c.Chat(context.Background(), ChatRequest{})
	if err != nil {
		t.Fatalf("Chat returned error: %v", err)
	}
	if content != "hello" {
		t.Fatalf("expected 'hello', got %q", content)
	}
}

func TestChatError(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "bad request")
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	c := &Client{
		apiKey:     "test",
		baseURL:    server.URL,
		httpClient: server.Client(),
	}

	_, err := c.Chat(context.Background(), ChatRequest{})
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "bad request" {
		t.Fatalf("unexpected error: %v", err)
	}
}
func TestNewClientEmptyKey(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "")
	if _, err := NewClient(); err == nil {
		t.Fatal("expected error when API key is empty")

	}
}

func TestGetModelFromEnv(t *testing.T) {
	t.Setenv("OPENAI_MODEL", "test-model")
	if got := GetModel(); got != "test-model" {
		t.Fatalf("expected 'test-model', got %q", got)
	}
}

func TestGetModelFallback(t *testing.T) {
	t.Setenv("OPENAI_MODEL", "")
	if got := GetModel(); got != "gpt-4o-mini" {
		t.Fatalf("expected fallback 'gpt-4o-mini', got %q", got)
	}
}
