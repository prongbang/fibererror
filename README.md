# fibererror 🚨

[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/fibererror)](https://goreportcard.com/report/github.com/prongbang/fibererror)
[![Go Reference](https://pkg.go.dev/badge/github.com/prongbang/fibererror.svg)](https://pkg.go.dev/github.com/prongbang/fibererror)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/prongbang/fibererror.svg)](https://golang.org)

> Elegant error response handling for Fiber applications with support for custom errors and internationalization.

## ✨ Features

- 🎯 **Simple Integration** - Easy to integrate with existing Fiber applications
- 🔧 **Custom Error Types** - Create your own error types with custom codes
- 🌍 **i18n Support** - Built-in internationalization for error messages
- ⚡ **High Performance** - Optimized for speed and efficiency
- 🎨 **Flexible Configuration** - Customizable response formats
- 📦 **Standard HTTP Errors** - Pre-defined standard HTTP error responses

## 📊 Performance

```shell
BenchmarkFiberErrorResponse_Response-10    	  186163	      7115 ns/op
BenchmarkBuildInErrorResponse_Response-10      	  143802	      7479 ns/op
```

## 📦 Installation

```shell
go get github.com/prongbang/fibererror
```

## 🚀 Quick Start

### Basic Usage with Standard HTTP Status

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

## 🛠️ Advanced Usage

### Custom Error Types

Create your own error types with custom codes:

```go
package main

import (
    "github.com/prongbang/goerror"
    "github.com/prongbang/fibererror"
    "github.com/gofiber/fiber/v2"
    "net/http"
)

// Define custom error type
type CustomError struct {
    goerror.Body
}

func (c *CustomError) Error() string {
    return c.Message
}

func NewCustomError() error {
    return &CustomError{
        Body: goerror.Body{
            Code:    "CUS001",
            Message: "Custom error occurred",
        },
    }
}

// Define custom response handler
type customResponse struct{}

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
    
    // Configure custom response
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

### 🌍 Internationalization Support

Localize error messages based on `Accept-Language` header:

#### 1. Create localization files

`localize/en.yaml`:
```yaml
CUS001: Custom error 001
```

`localize/th.yaml`:
```yaml
CUS001: ข้อผิดพลาดแบบกำหนดเอง 001
```

#### 2. Configure i18n with fibererror

```go
package main

import (
    "fmt"
    "github.com/gofiber/contrib/fiberi18n/v2"
    "github.com/gofiber/fiber/v2"
    "github.com/prongbang/fibererror"
    "golang.org/x/text/language"
)

func main() {
    app := fiber.New()
    
    // Configure i18n middleware
    app.Use(fiberi18n.New(&fiberi18n.Config{
        RootPath:        "./localize",
        AcceptLanguages: []language.Tag{language.Thai, language.English},
        DefaultLanguage: language.English,
    }))
    
    // Configure fibererror with i18n
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
        fmt.Println("Error starting server:", err)
    }
}
```

## 📝 Configuration Options

### fibererror.Config

| Option | Type | Description |
|--------|------|-------------|
| `Custom` | `*Custom` | Custom error response handler |
| `I18n` | `*I18n` | Internationalization configuration |

### fibererror.I18n

| Option | Type | Description |
|--------|------|-------------|
| `Enabled` | `bool` | Enable/disable i18n support |
| `Localize` | `func(*fiber.Ctx, string) (string, error)` | Localization function |

## 🔍 Examples

### Handling Multiple Error Types

```go
func (c *customResponse) Response(ctx *fiber.Ctx, err error) error {
    switch resp := err.(type) {
    case *CustomError:
        return ctx.Status(http.StatusBadRequest).JSON(resp)
    case *AuthError:
        return ctx.Status(http.StatusUnauthorized).JSON(resp)
    case *ValidationError:
        return ctx.Status(http.StatusUnprocessableEntity).JSON(resp)
    default:
        return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "code": "INTERNAL_ERROR",
            "message": "An unexpected error occurred",
        })
    }
}
```

### Error Response Format

Standard error response structure:

```json
{
    "code": "CUS001",
    "message": "Custom error message"
}
```

With additional fields:

```json
{
    "code": "VAL001",
    "message": "Validation failed",
    "details": {
        "field": "email",
        "reason": "invalid format"
    }
}
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 💖 Support the Project

If you find this package helpful, please consider supporting it:

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/prongbang)

## 🔗 Related Projects

- [goerror](https://github.com/prongbang/goerror) - Error handling utilities for Go
- [Fiber](https://github.com/gofiber/fiber) - Express-inspired web framework
- [fiberi18n](https://github.com/gofiber/contrib/tree/main/fiberi18n) - i18n middleware for Fiber

---
