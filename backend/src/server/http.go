package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
	"shrtify.de/src/admin"
	"shrtify.de/src/shorturl"
)

func NewHttpServer(db *gorm.DB, errorHandler func(*fiber.Ctx, error) error) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	app.Use(logger.New())

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	api := app.Group("/api")

	v1 := api.Group("/v1")

	admin.Route(v1.Group("/admin"), admin.GetMiddleware(db))
	shorturl.Route(app.Group("/p"), shorturl.GetMiddleware(db))

	return app
}
