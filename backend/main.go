package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/v1/suggestions", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "suggestions placeholder"})
	})

	app.Get("/api/v1/details", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "details placeholder"})
	})

	app.Listen(":8080")
}
