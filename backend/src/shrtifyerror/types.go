package shrtifyerror

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type ResponseError struct {
	error
	PublicMessage string
	Code          int
	More          any
}

func ResponseShorturlCodeExists(shorturlCode string) ResponseError {
	return ResponseError{
		error:         errors.New("failed to create shorturl bebause code already exists"),
		PublicMessage: "admin_middleware_shorturl_code_exists",
		Code:          fiber.StatusBadRequest,
		More:          shorturlCode,
	}
}

func (e *ResponseError) Error() string {
	return e.PublicMessage
}
