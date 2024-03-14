package admin

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jaevor/go-nanoid"
	"gorm.io/gorm"
	"shrtify.de/src/shrtifyerror"
	"shrtify.de/src/types"
)

func GetMiddleware(db *gorm.DB) Middleware {
	canonicID, _ := nanoid.Standard(10)
	return Middleware{
		createShortcode: func(ctx *fiber.Ctx) error {
			request := &types.CreateShortcodeRequestPayload{}

			err := json.Unmarshal(ctx.Body(), &request)
			if err != nil {
				panic(err)
			}
			code := ""
			if request.Request.Code != "" {
				code = request.Request.Code
			} else {
				code = canonicID()
			}
			var count int64
			db.Model(&types.ShortUrl{}).Where("code = ?", code).Count(&count)
			if count > 0 {
				panic(shrtifyerror.ResponseShorturlCodeExists(code))
			}
			var host string
			if request.Request.Host == "" {
				host = ctx.Hostname()
			} else {
				host = request.Request.Host
			}
			result := db.Create(&types.ShortUrl{
				Code: code,
				Url:  request.Request.Url,
				Host: host,
			})

			if result.Error != nil {
				panic(errors.Join(errors.New("failed to create shorturl"), result.Error))
			}

			return ctx.JSON(types.CreateShortcodeResponsePayload{
				Meta: types.Meta{
					Status: fiber.StatusCreated,
					Error:  nil,
				},
				Response: types.CreateShortcodeResponse{
					ShortUrl: ctx.Protocol() + "://" + host + "/p/" + code,
				},
			})
		},
		getShortcodes: func(ctx *fiber.Ctx) error {

			page := ctx.QueryInt("page", 1)
			items := ctx.QueryInt("items", 25)

			models := &[]types.ShortUrl{}
			result := db.Model(types.ShortUrl{}).Limit(items).Offset((page - 1) * items).Find(models)

			if result.Error != nil {
				panic(errors.Join(errors.New("failed to fetch shorturls"), result.Error))
			}
			var count int64
			db.Model(types.ShortUrl{}).Count(&count)
			payload := types.GetShortcodesResponsePayload{
				Meta: types.Meta{
					Status: fiber.StatusOK,
					Error:  nil,
					More: types.GetShortcodesMeta{
						Pages: count / int64(items),
					},
				},
				Response: make([]types.GetShortcodesResponse, 0),
			}

			for _, model := range *models {
				payload.Response = append(payload.Response, types.GetShortcodesResponse{
					ShortUrl:    ctx.Protocol() + "://" + ctx.Hostname() + "/p/" + model.Code,
					OriginalUrl: model.Url,
				})
			}

			return ctx.JSON(payload)
		},
	}
}

type Middleware struct {
	createShortcode fiber.Handler
	getShortcodes   fiber.Handler
}
