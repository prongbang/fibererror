package fibererror_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/fibererror"
	"github.com/prongbang/goerror"
	"net/http"
	"net/http/httptest"
	"testing"
)

var response = fibererror.New()

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

func TestNewCustomError(t *testing.T) {
	app := fiber.New()

	customResp := NewCustomResponse()
	res := fibererror.New(&fibererror.Config{
		Custom: &customResp,
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return res.With(c).Response(NewCustomError())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusBadRequest {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewUseProxy(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewUseProxy())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusUseProxy {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewUnauthorized(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewUnauthorized())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusUnauthorized {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewTemporaryRedirect(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewTemporaryRedirect())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusTemporaryRedirect {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewNotFound(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewNotFound())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusNotFound {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewSwitchingProtocols(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewSwitchingProtocols())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusSwitchingProtocols {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewSeeOther(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewSeeOther())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusSeeOther {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewResetContent(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewResetContent())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusResetContent {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewRequestTimeout(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewRequestTimeout())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusRequestTimeout {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewProxyAuthRequired(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewProxyAuthRequired())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusProxyAuthRequired {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewProcessing(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewProcessing())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusProcessing {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewPermanentRedirect(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewPermanentRedirect())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusPermanentRedirect {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewPaymentRequired(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewPaymentRequired())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusPaymentRequired {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewPartialContent(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewPartialContent())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusPartialContent {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewOK(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewOK(nil))
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusOK {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewNotModified(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewNotModified())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusNotModified {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewNotAcceptable(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewNotAcceptable())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusNotAcceptable {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewNonAuthoritativeInformation(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewNonAuthoritativeInformation())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusNonAuthoritativeInfo {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewNoContent(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewNoContent())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusNoContent {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewMultipleChoices(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewMultipleChoices())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusMultipleChoices {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewMultiStatus(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewMultiStatus())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusMultiStatus {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewMovedPermanently(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewMovedPermanently())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusMovedPermanently {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewMethodNotAllowed(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewMethodNotAllowed())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewIMUsed(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewIMUsed())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusIMUsed {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewFound(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewFound())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusFound {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewForbidden(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewForbidden())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusForbidden {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewEarlyHints(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewEarlyHints())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusEarlyHints {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewCreated(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewCreated(nil))
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusCreated {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewContinue(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewContinue())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusContinue {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewConflict(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewConflict())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusConflict {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewBadRequest(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewBadRequest())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusBadRequest {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewAlreadyReported(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewAlreadyReported())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusAlreadyReported {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewAccepted(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewAccepted())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusAccepted {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewGone(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewGone())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusGone {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewLengthRequired(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewLengthRequired())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusLengthRequired {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewPreconditionFailed(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewPreconditionFailed())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusPreconditionFailed {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewRequestEntityTooLarge(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewRequestEntityTooLarge())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusRequestEntityTooLarge {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewRequestURITooLong(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewRequestURITooLong())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusRequestURITooLong {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewUnsupportedMediaType(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewUnsupportedMediaType())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusUnsupportedMediaType {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewRequestedRangeNotSatisfiable(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewRequestedRangeNotSatisfiable())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusRequestedRangeNotSatisfiable {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewExpectationFailed(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewExpectationFailed())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusExpectationFailed {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewTeapot(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewTeapot())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusTeapot {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewMisdirectedRequest(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewMisdirectedRequest())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusMisdirectedRequest {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewUnprocessableEntity(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewUnprocessableEntity())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewLocked(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewLocked())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusLocked {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewFailedDependency(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewFailedDependency())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusFailedDependency {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewTooEarly(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewTooEarly())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusTooEarly {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewUpgradeRequired(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewUpgradeRequired())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusUpgradeRequired {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewPreconditionRequired(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewPreconditionRequired())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusPreconditionRequired {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewTooManyRequests(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewTooManyRequests())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusTooManyRequests {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewRequestHeaderFieldsTooLarge(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewRequestHeaderFieldsTooLarge())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusRequestHeaderFieldsTooLarge {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewUnavailableForLegalReasons(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewUnavailableForLegalReasons())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusUnavailableForLegalReasons {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewInternalServerError(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewInternalServerError())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusInternalServerError {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewNotImplemented(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewNotImplemented())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusNotImplemented {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewBadGateway(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewBadGateway())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusBadGateway {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewServiceUnavailable(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewServiceUnavailable())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewGatewayTimeout(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewGatewayTimeout())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusGatewayTimeout {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewHTTPVersionNotSupported(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewHTTPVersionNotSupported())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusHTTPVersionNotSupported {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewVariantAlsoNegotiates(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewVariantAlsoNegotiates())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusVariantAlsoNegotiates {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewInsufficientStorage(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewInsufficientStorage())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusInsufficientStorage {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewLoopDetected(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewLoopDetected())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusLoopDetected {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewNotExtended(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewNotExtended())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusNotExtended {
		t.Error("Error", resp.StatusCode)
	}
}

func TestNewNetworkAuthenticationRequired(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewNetworkAuthenticationRequired())
	})

	resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

	if resp.StatusCode != http.StatusNetworkAuthenticationRequired {
		t.Error("Error", resp.StatusCode)
	}
}

func BenchmarkFiberErrorResponse_Response(b *testing.B) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return response.With(c).Response(goerror.NewNetworkAuthenticationRequired())
	})

	for i := 0; i < b.N; i++ {
		resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

		if resp.StatusCode != http.StatusNetworkAuthenticationRequired {
			b.Error("Error", resp.StatusCode)
		}
	}
}

func BenchmarkBuildInErrorResponse_Response(b *testing.B) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNetworkAuthenticationRequired).
			JSON(goerror.NetworkAuthenticationRequired{
				Body: goerror.Body{
					Code:    goerror.CodeNetworkAuthenticationRequired,
					Message: http.StatusText(http.StatusNetworkAuthenticationRequired),
				},
			})
	})

	for i := 0; i < b.N; i++ {
		resp, _ := app.Test(httptest.NewRequest("GET", "/test", nil))

		if resp.StatusCode != http.StatusNetworkAuthenticationRequired {
			b.Error("Error", resp.StatusCode)
		}
	}
}
