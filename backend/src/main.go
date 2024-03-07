package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"shrtify.de/src/database"
	"shrtify.de/src/server"
	"shrtify.de/src/shrtifyerror"
	"shrtify.de/src/types"
)

type ServerStatus struct {
	Status string
}

func main() {
	types.TypesInitialization()
	fx.New(
		fx.Provide(
			shrtifyerror.NewErrorHandler,
			database.NewPostgres,
			server.NewHttpServer,
		),
		fx.Invoke(func(app *fiber.App) {
			log.Println("Server started.")
			log.Fatal(app.Listen(":3000"))
		}),
	).Run()

}
