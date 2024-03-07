package shrtifyerror

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"shrtify.de/src/types"
)

func NewErrorHandler() func(ctx *fiber.Ctx, err error) error {
	return func(ctx *fiber.Ctx, err error) error {
		var initialResponseError error
		if e, ok := err.(*fiber.Error); ok {
			initialResponseError = ctx.Status(e.Code).SendFile(fmt.Sprintf("./%d.html", e.Code))

		} else if e, ok := err.(*ResponseError); ok {
			initialResponseError = ctx.Status(e.Code).JSON(types.Meta{
				Status: e.Code,
				Error:  &e.PublicMessage,
			})
		}

		if initialResponseError != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		return nil
	}
}
