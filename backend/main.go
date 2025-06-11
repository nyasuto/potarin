package main

import (
	"context"
	"encoding/json"
	"fmt"

	"potarin-backend/internal"
	shared "potarin-shared"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type UserProfile struct {
	Name        string   `json:"name"`
	Location    string   `json:"location"`
	BikeType    string   `json:"bikeType"`
	Level       string   `json:"level"`
	Preferences []string `json:"preferences"`
}

// ユーザープロフィールのサンプルデータ
var userProfile = UserProfile{
	Name:        "サンプルユーザー",
	Location:    "藤沢市",
	BikeType:    "クロスバイク",
	Level:       "初心者",
	Preferences: []string{"海沿い", "自然", "カフェ"},
}

func main() {
	internal.LoadEnv()
	ai, err := internal.NewClient()
	if err != nil {
		fmt.Printf("failed to create OpenAI client: %v\n", err)
		return
	}
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/api/v1/suggestions", func(c *fiber.Ctx) error {
		prompt := c.Query("prompt")
		suggestions, err := fetchSuggestions(c.Context(), ai, prompt)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(suggestions)
	})

	app.Post("/api/v1/details", func(c *fiber.Ctx) error {
		var s shared.Suggestion
		if err := c.BodyParser(&s); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		detail, err := fetchDetail(c.Context(), ai, s)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(detail)
	})

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}

type AIClient interface {
	Chat(ctx context.Context, req internal.ChatRequest) (string, error)
}

func fetchSuggestions(ctx context.Context, ai AIClient, userPrompt string) ([]shared.Suggestion, error) {
	userProf, err := json.Marshal(userProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user profile: %w", err)
	}
	schema := `{"type":"object","properties":{"suggestions":{"type":"array","items":{"type":"object","properties":{"id":{"type":"string"},"title":{"type":"string"},"description":{"type":"string"}},"required":["id","title","description"]}}},"required":["suggestions"]}`
	if userPrompt == "" {
		userPrompt = "今日は天気が良いので、3つの異なるサイクリングコースを提案してください。日付とその季節を考慮してください 本日は六月です"
	}
	req := internal.ChatRequest{
		Model: "gpt-4o",
		Messages: []internal.Message{
			{
				Role: "system",
				Content: "Return course suggestions as JSON.\n" + schema +
					"\nあなたは親しみやすく、情報に詳しいサイクリングアドバイザーです。ユーザーのプロフィールは以下の通りです：" + string(userProf),
			},
			{Role: "user", Content: userPrompt},
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

func fetchDetail(ctx context.Context, ai AIClient, suggestion shared.Suggestion) (shared.Detail, error) {
	schema := `{"type":"object","properties":{"summary":{"type":"string"},"routes":{"type":"array","items":{"type":"object","properties":{"title":{"type":"string"},"description":{"type":"string"},"position":{"type":"object","properties":{"lat":{"type":"number"},"lng":{"type":"number"}},"required":["lat","lng"]}},"required":["title","description","position"]}}},"required":["summary","routes"]}`
	userPrompt := fmt.Sprintf("タイトル: %s\n説明: %s\nこのコースの詳細を教えてください。", suggestion.Title, suggestion.Description)
	req := internal.ChatRequest{
		Model: "gpt-4o",
		Messages: []internal.Message{
			{Role: "system", Content: "Return course detail as JSON.\n" + schema},
			{Role: "user", Content: userPrompt},
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
