package main

import (
	"fmt"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/fibererror"
	"golang.org/x/text/language"
	"net/http"
)

type CustomError struct {
	fibererror.Body
}

// Error implements error.
func (c *CustomError) Error() string {
	return c.Message
}

func NewCustomError() error {
	return &CustomError{
		Body: fibererror.Body{
			Code: "CUS001",
		},
	}
}

type customResponse struct {
}

// Response implements response.Custom.
func (c *customResponse) Response(ctx *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *CustomError:
		return ctx.Status(http.StatusBadRequest).JSON(e)
	}
	return nil
}

func NewCustomResponse() fibererror.Custom {
	return &customResponse{}
}

func main() {
	app := fiber.New()
	app.Use(fiberi18n.New(&fiberi18n.Config{
		RootPath:        "./localize",
		AcceptLanguages: []language.Tag{language.Thai, language.English},
		DefaultLanguage: language.English,
	}))

	customResp := NewCustomResponse()
	response := fibererror.New(&fibererror.Config{
		Custom: &customResp,
		I18n: &fibererror.I18n{
			Enabled: true,
			Localize: func(ctx *fiber.Ctx, code string) (string, error) {
				return fiberi18n.Localize(ctx, code)
			},
		},
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return response.With(c).Response(NewCustomError())
	})

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("err:")
	}
}
