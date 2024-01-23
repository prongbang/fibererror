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
	"github.com/prongbang/fibererror"
	"github.com/gofiber/fiber/v2"
)

func login() error {
	return fibererror.NewUnauthorized()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		err := login()
		return fibererror.New(c).Response(err)
	})

	app.Listen(":3000")
}
```

- Use by Custom Code

```go
package main

import (
	"github.com/prongbang/fibererror"
	"github.com/gofiber/fiber/v2"
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
			Code:    "CUS001",
			Message: "Custom 001",
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

func custom() error {
	return NewCustomError()
}

func main() {
	app := fiber.New()

	customResp := NewCustomResponse()

	app.Get("/", func(c *fiber.Ctx) error {
		err := custom()
		return fibererror.New(c).Custom(customResp).Response(err)
	})

	app.Listen(":3000")
}
```
