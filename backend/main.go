package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"potarin-backend/internal"
	shared "potarin-shared"
)

func main() {
	internal.LoadEnv()
	ai := internal.NewClient()
	app := fiber.New()

	app.Get("/api/v1/suggestions", func(c *fiber.Ctx) error {
		suggestions, err := fetchSuggestions(c.Context(), ai)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(suggestions)
	})

	app.Get("/api/v1/details", func(c *fiber.Ctx) error {
		id := c.Query("id")
		if id == "" {
			return fiber.NewError(fiber.StatusBadRequest, "id required")
		}
		detail, err := fetchDetail(c.Context(), ai, id)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(detail)
	})

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}

func fetchSuggestions(ctx context.Context, ai *internal.Client) ([]shared.Suggestion, error) {
	schema := `{"type":"object","properties":{"suggestions":{"type":"array","items":{"type":"object","properties":{"id":{"type":"string"},"title":{"type":"string"},"description":{"type":"string"}},"required":["id","title","description"]}}},"required":["suggestions"]}`
	req := internal.ChatRequest{
		Model: "gpt-4o",
		Messages: []internal.Message{
			{Role: "system", Content: "Return course suggestions as JSON." + schema},
			{Role: "user", Content: "Please suggest some courses."},
		},
		ResponseFormat: internal.ResponseFormat{Type: "json_object"},
	}
	content, err := ai.Chat(ctx, req)
	if err != nil {
		return nil, err
	}
	var parsed struct {
		Suggestions []shared.Suggestion `json:"suggestions"`
	}
	if err := json.Unmarshal([]byte(content), &parsed); err != nil {
		return nil, err
	}
	return parsed.Suggestions, nil
}

func fetchDetail(ctx context.Context, ai *internal.Client, id string) (shared.Detail, error) {
	schema := `{"type":"object","properties":{"summary":{"type":"string"},"routes":{"type":"array","items":{"type":"object","properties":{"title":{"type":"string"},"description":{"type":"string"},"position":{"type":"object","properties":{"lat":{"type":"number"},"lng":{"type":"number"}},"required":["lat","lng"]}},"required":["title","description","position"]}}},"required":["summary","routes"]}`
	req := internal.ChatRequest{
		Model: "gpt-4o",
		Messages: []internal.Message{
			{Role: "system", Content: "Return course detail as JSON." + schema},
			{Role: "user", Content: fmt.Sprintf("detail for course %s", id)},
		},
		ResponseFormat: internal.ResponseFormat{Type: "json_object"},
	}
	content, err := ai.Chat(ctx, req)
	if err != nil {
		return shared.Detail{}, err
	}
	var detail shared.Detail
	if err := json.Unmarshal([]byte(content), &detail); err != nil {
		return shared.Detail{}, err
	}
	return detail, nil
}
