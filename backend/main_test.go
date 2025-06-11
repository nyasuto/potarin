package main

import (
	"context"
	"testing"

	"potarin-backend/internal"
	shared "potarin-shared"
)

type mockClient struct {
	response string
	err      error
}

func (m mockClient) Chat(ctx context.Context, req internal.ChatRequest) (string, error) {
	return m.response, m.err
}

func TestFetchSuggestionsSuccess(t *testing.T) {
	mc := mockClient{response: `{"suggestions":[{"id":"1","title":"Course","description":"Desc"}]}`}
	got, err := fetchSuggestions(context.Background(), mc, "")
	if err != nil {
		t.Fatalf("fetchSuggestions returned error: %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("expected 1 suggestion, got %d", len(got))
	}
	if got[0].ID != "1" || got[0].Title != "Course" {
		t.Fatalf("unexpected suggestion: %+v", got[0])
	}
}

func TestFetchSuggestionsInvalidJSON(t *testing.T) {
	mc := mockClient{response: `invalid`}
	if _, err := fetchSuggestions(context.Background(), mc, ""); err == nil {
		t.Fatal("expected error from invalid JSON")
	}
}

func TestFetchDetailSuccess(t *testing.T) {
	resp := `{"summary":"sum","routes":[{"title":"t","description":"d","position":{"lat":1,"lng":2}}]}`
	mc := mockClient{response: resp}
	sug := shared.Suggestion{ID: "1", Title: "t", Description: "d"}
	got, err := fetchDetail(context.Background(), mc, sug)
	if err != nil {
		t.Fatalf("fetchDetail returned error: %v", err)
	}
	if got.Summary != "sum" || len(got.Routes) != 1 {
		t.Fatalf("unexpected detail: %+v", got)
	}
	if got.Routes[0].Position.Lat != 1 || got.Routes[0].Position.Lng != 2 {
		t.Fatalf("unexpected position: %+v", got.Routes[0].Position)
	}
}

func TestFetchDetailInvalidJSON(t *testing.T) {
	mc := mockClient{response: `invalid`}
	sug := shared.Suggestion{}
	if _, err := fetchDetail(context.Background(), mc, sug); err == nil {
		t.Fatal("expected error from invalid JSON")
	}
}
