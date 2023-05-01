package main

import (
	"errors"
	"log"

	"github.com/allanfs/balance-server/internal/app/inbound"
	"github.com/allanfs/balance-server/internal/app/infra"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := infra.StartPostgres()
	FinishHim(err)

	defer db.Close()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Running")
	})

	api := app.Group("api")
	api_v1 := api.Group("v1")

	inbound.SetLancamentosRoutes(api_v1.Group("lancamentos"))

	app.Listen(":3000")
}

func FinishHim(err error) {
	if err != nil {

		ie := new(infra.InfraError)
		if errors.As(err, &ie) {

			log.Fatalf("Falha ao iniciar o %s: %s, causado por: %v ", ie.Name, ie.Message, ie.Error())
		} else {
			log.Fatalf("Falha ao iniciar: %v ", err)
		}
	}

}
