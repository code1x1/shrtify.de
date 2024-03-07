package shorturl

import (
	"github.com/gofiber/fiber/v2"
)

func Route(r fiber.Router, mdlw Middleware) {
	r.Get("/:code", mdlw.getUrl)
}
