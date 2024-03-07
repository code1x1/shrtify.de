package admin

import (
	"github.com/gofiber/fiber/v2"
)

func Route(r fiber.Router, mdlw Middleware) {
	r.Post("/createShortcode", mdlw.createShortcode)
	r.Get("/getShortcodes", mdlw.getShortcodes)
}
