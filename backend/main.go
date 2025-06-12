package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"potarin-backend/internal"
	shared "potarin-shared"
	schemas "potarin-shared/schemas"

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
	log.Println("environment variables loaded")
	ai, err := internal.NewClient()
	if err != nil {
		log.Fatalf("failed to create OpenAI client: %v", err)
	}
	log.Println("OpenAI client initialized")
	app := fiber.New()
	log.Println("fiber app initialized")
	app.Use(cors.New())

	app.Get("/api/v1/suggestions", func(c *fiber.Ctx) error {
		prompt := c.Query("prompt")
		log.Printf("received suggestions request: %s", prompt)
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
		log.Printf("received detail request for: %s", s.Title)
		detail, err := fetchDetail(c.Context(), ai, s)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(detail)
	})

	log.Println("server listening on :8080")
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
	if userPrompt == "" {
		userPrompt = "今日は天気が良いので、3つの異なるサイクリングコースを提案してください。日付とその季節を考慮してください 本日は六月です"
	}
	log.Printf("fetching suggestions with prompt: %s", userPrompt)
	req := internal.ChatRequest{
		Model: internal.GetModel(),
		Messages: []internal.Message{
			{
				Role: "system",
				Content: "Return course suggestions as JSON.\n" +
					"あなたは親しみやすく、情報に詳しいサイクリングアドバイザーです。ユーザーのプロフィールは以下の通りです：" + string(userProf),
			},
			{Role: "user", Content: userPrompt},
		},
		ResponseFormat: internal.ResponseFormat{
			Type: "json_schema",
			JSONSchema: &internal.JSONSchemaConfig{
				Name:   "Suggestions",
				Schema: json.RawMessage(schemas.SuggestionsSchema),
				Strict: true,
			},
		},
	}
	log.Print(req)
	content, err := ai.Chat(ctx, req)
	log.Print(content, err)

	if err != nil {
		return nil, err
	}

	var parsed struct {
		Suggestions []shared.Suggestion `json:"suggestions"`
	}
	if err := json.Unmarshal([]byte(content), &parsed); err != nil {
		return nil, err
	}
	if parsed.Suggestions == nil {
		return []shared.Suggestion{}, nil
	}
	log.Printf("parsed %d suggestions", len(parsed.Suggestions))
	return parsed.Suggestions, nil
}

func fetchDetail(ctx context.Context, ai AIClient, suggestion shared.Suggestion) (shared.Detail, error) {
	userPrompt := fmt.Sprintf("タイトル: %s\n説明: %s\nこのコースの詳細を教えてください。", suggestion.Title, suggestion.Description)
	log.Printf("fetching detail for: %s", suggestion.Title)
	req := internal.ChatRequest{
		Model: internal.GetModel(),
		Messages: []internal.Message{
			{Role: "system", Content: "Return course detail as JSON."},
			{Role: "user", Content: userPrompt},
		},
		ResponseFormat: internal.ResponseFormat{
			Type: "json_schema",
			JSONSchema: &internal.JSONSchemaConfig{
				Name:   "Detail",
				Schema: json.RawMessage(schemas.DetailSchema),
				Strict: true,
			},
		},
	}
	content, err := ai.Chat(ctx, req)
	if err != nil {
		return shared.Detail{}, err
	}
	var detail shared.Detail
	if err := json.Unmarshal([]byte(content), &detail); err != nil {
		return shared.Detail{}, err
	}
	log.Println("detail parsed successfully")
	return detail, nil
}
