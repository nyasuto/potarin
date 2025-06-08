package main

import (
	"github.com/gofiber/fiber/v2"
	shared "potarin-shared"
)

var suggestions = []shared.Suggestion{
	{ID: "1", Title: "River Side Ride", Description: "Enjoy the river view."},
	{ID: "2", Title: "City Exploration", Description: "Discover downtown."},
}

var detailMap = map[string]shared.Detail{
	"1": {
		Summary: "River Side Ride details",
		Routes: []shared.Route{
			{
				Title:       "Start Point",
				Description: "Park",
				Position:    shared.Position{Lat: 35.0, Lng: 135.0},
			},
			{
				Title:       "End Point",
				Description: "Bridge",
				Position:    shared.Position{Lat: 35.1, Lng: 135.1},
			},
		},
	},
	"2": {
		Summary: "City Exploration details",
		Routes: []shared.Route{
			{
				Title:       "Central Plaza",
				Description: "Meeting spot",
				Position:    shared.Position{Lat: 35.2, Lng: 135.2},
			},
			{
				Title:       "Museum",
				Description: "Local history",
				Position:    shared.Position{Lat: 35.3, Lng: 135.3},
			},
		},
	},
}

func main() {
	app := fiber.New()

	app.Get("/api/v1/suggestions", func(c *fiber.Ctx) error {
		return c.JSON(suggestions)
	})

	app.Get("/api/v1/details", func(c *fiber.Ctx) error {
		id := c.Query("id")
		detail, ok := detailMap[id]
		if !ok {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
		}
		return c.JSON(detail)
	})

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
