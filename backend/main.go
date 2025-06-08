package main

import (
	"github.com/gofiber/fiber/v2"
	shared "potarin-shared"
)

func main() {
	app := fiber.New()

	app.Get("/api/v1/suggestions", func(c *fiber.Ctx) error {
		suggestions := []shared.Suggestion{
			{ID: "1", Title: "River Side Ride", Description: "Enjoy the river view."},
			{ID: "2", Title: "City Exploration", Description: "Discover downtown."},
		}
		return c.JSON(suggestions)
	})

	app.Get("/api/v1/details", func(c *fiber.Ctx) error {
		details := shared.Detail{
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
		}
		return c.JSON(details)
	})

	app.Listen(":8080")
}
