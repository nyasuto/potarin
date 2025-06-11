package internal

import "testing"

func TestNewClientEmptyKey(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "")
	if _, err := NewClient(); err == nil {
		t.Fatal("expected error when API key is empty")
	}
}
