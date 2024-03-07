package shorturl

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"shrtify.de/src/types"
)

func GetMiddleware(db *gorm.DB) Middleware {
	return Middleware{
		getUrl: func(ctx *fiber.Ctx) error {
			code := ctx.Params("code")
			if len(code) < 1 {
				panic(errors.New("redirect code not found"))
			}
			model := &types.ShortUrl{
				Code: code,
			}
			result := db.Where("code = ?", code).First(model)
			if result.Error != nil {
				panic(errors.New("redirect url not found"))
			}
			redirectUrl, err := url.Parse(model.Url)
			if err != nil {
				panic(errors.New("redirect url parse error"))
			}
			if redirectUrl.Host == ctx.Hostname() {
				panic(errors.New("redirect url same hostname"))
			}
			return ctx.Redirect(redirectUrl.String(), http.StatusMovedPermanently)
		},
	}
}

type Middleware struct {
	getUrl fiber.Handler
}
