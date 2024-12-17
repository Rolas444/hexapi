package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New(
		// add buffer size to chrome
		fiber.Config{
			ReadBufferSize: 1024 * 64,
		},
	)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Corriendo API!")
	})

	app.Listen(":3000")
}
