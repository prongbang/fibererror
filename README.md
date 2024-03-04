# fibererror

[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/fibererror)](https://goreportcard.com/report/github.com/prongbang/fibererror)

Create a custom response using the error interface in Go.

### Install

```shell
go get github.com/prongbang/fibererror
```

### How to use

- Use by Standard HTTP Status Code

```go
package main

import (
	"github.com/prongbang/goerror"
	"github.com/prongbang/fibererror"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	response := fibererror.New()
	
	app.Get("/", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewUnauthorized())
	})

	_ = app.Listen(":3000")
}
```

- Use by Custom Code

```go
package main

import (
	"github.com/prongbang/goerror"
	"github.com/prongbang/fibererror"
	"github.com/gofiber/fiber/v2"
)

type CustomError struct {
	goerror.Body
}

// Error implements error.
func (c *CustomError) Error() string {
	return c.Message
}

func NewCustomError() error {
	return &CustomError{
		Body: goerror.Body{
			Code:    "CUS001",
		},
	}
}

type customResponse struct {
}

// Response implements response.Custom.
func (c *customResponse) Response(ctx *fiber.Ctx, err error) error {
	switch resp := err.(type) {
	case *CustomError:
		return ctx.Status(http.StatusBadRequest).JSON(resp)
	}
	return nil
}

func NewCustomResponse() fibererror.Custom {
	return &customResponse{}
}

func main() {
	app := fiber.New()

	customResp := NewCustomResponse()
	response := fibererror.New(&fibererror.Config{
		Custom: &customResp,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return response.With(c).Response(NewCustomError())
	})

	_ = app.Listen(":3000")
}
```

- Use by Custom Code + Localization header `Accept-Language: en`

localize/en.yaml
```yaml
CUS001: Custom 001
```

localize/th.yaml
```yaml
CUS001: ดัดแปลง 001
```

```go
package main

import (
	"fmt"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/fibererror"
	"golang.org/x/text/language"
	"net/http"
)

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
```