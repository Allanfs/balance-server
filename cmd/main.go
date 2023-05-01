package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Running")
	})

	api := app.Group("api")
	api_v1 := api.Group("v1")

	lancamentosGroup := api_v1.Group("lancamentos")
	lancamentosGroup.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("Rota cadastrar lançamentos")
	})
	lancamentosGroup.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Rota obters lançamentos")
	})

	lancamentosGroup.Get("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Rota de obter lançamento por id: " + c.Params("id"))
	})

	app.Listen(":3000")
}
