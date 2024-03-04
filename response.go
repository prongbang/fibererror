package fibererror

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/goerror"
	"net/http"
)

type Config struct {
	Custom *Custom
	I18n   *I18n
}

type I18n struct {
	Enabled  bool
	Localize func(c *fiber.Ctx, code string) (string, error)
}

type Custom interface {
	Response(ctx *fiber.Ctx, err error) error
}

type Response interface {
	With(c *fiber.Ctx) HttpResponse
}

type HttpResponse interface {
	Response(err error) error
}

type response struct {
	Cus  *Custom
	I18n *I18n
}

type httpResponse struct {
	Ctx  *fiber.Ctx
	Cus  *Custom
	I18n *I18n
}

// With implements Response.
func (r *response) With(c *fiber.Ctx) HttpResponse {
	return &httpResponse{
		Ctx:  c,
		Cus:  r.Cus,
		I18n: r.I18n,
	}
}

// Response implements Response.
func (s *httpResponse) Response(err error) error {
	switch e := err.(type) {
	// Information
	case *goerror.Continue:
		return s.Ctx.Status(http.StatusContinue).JSON(e)
	case *goerror.SwitchingProtocols:
		return s.Ctx.Status(http.StatusSwitchingProtocols).JSON(e)
	case *goerror.Processing:
		return s.Ctx.Status(http.StatusProcessing).JSON(e)
	case *goerror.EarlyHints:
		return s.Ctx.Status(http.StatusEarlyHints).JSON(e)

	// Successful
	case *goerror.OK:
		return s.Ctx.Status(http.StatusOK).JSON(e)
	case *goerror.Created:
		return s.Ctx.Status(http.StatusCreated).JSON(e)
	case *goerror.Accepted:
		return s.Ctx.Status(http.StatusAccepted).JSON(e)
	case *goerror.NonAuthoritativeInformation:
		return s.Ctx.Status(http.StatusNonAuthoritativeInfo).JSON(e)
	case *goerror.NoContent:
		return s.Ctx.Status(http.StatusNoContent).JSON(e)
	case *goerror.ResetContent:
		return s.Ctx.Status(http.StatusResetContent).JSON(e)
	case *goerror.PartialContent:
		return s.Ctx.Status(http.StatusPartialContent).JSON(e)
	case *goerror.MultiStatus:
		return s.Ctx.Status(http.StatusMultiStatus).JSON(e)
	case *goerror.AlreadyReported:
		return s.Ctx.Status(http.StatusAlreadyReported).JSON(e)
	case *goerror.IMUsed:
		return s.Ctx.Status(http.StatusIMUsed).JSON(e)

	// Redirection
	case *goerror.MultipleChoices:
		return s.Ctx.Status(http.StatusMultipleChoices).JSON(e)
	case *goerror.MovedPermanently:
		return s.Ctx.Status(http.StatusMovedPermanently).JSON(e)
	case *goerror.Found:
		return s.Ctx.Status(http.StatusFound).JSON(e)
	case *goerror.SeeOther:
		return s.Ctx.Status(http.StatusSeeOther).JSON(e)
	case *goerror.NotModified:
		return s.Ctx.Status(http.StatusNotModified).JSON(e)
	case *goerror.UseProxy:
		return s.Ctx.Status(http.StatusUseProxy).JSON(e)
	case *goerror.TemporaryRedirect:
		return s.Ctx.Status(http.StatusTemporaryRedirect).JSON(e)
	case *goerror.PermanentRedirect:
		return s.Ctx.Status(http.StatusPermanentRedirect).JSON(e)

	// Client error
	case *goerror.BadRequest:
		return s.Ctx.Status(http.StatusBadRequest).JSON(e)
	case *goerror.Unauthorized:
		return s.Ctx.Status(http.StatusUnauthorized).JSON(e)
	case *goerror.PaymentRequired:
		return s.Ctx.Status(http.StatusPaymentRequired).JSON(e)
	case *goerror.Forbidden:
		return s.Ctx.Status(http.StatusForbidden).JSON(e)
	case *goerror.NotFound:
		return s.Ctx.Status(http.StatusNotFound).JSON(e)
	case *goerror.MethodNotAllowed:
		return s.Ctx.Status(http.StatusMethodNotAllowed).JSON(e)
	case *goerror.NotAcceptable:
		return s.Ctx.Status(http.StatusNotAcceptable).JSON(e)
	case *goerror.ProxyAuthRequired:
		return s.Ctx.Status(http.StatusProxyAuthRequired).JSON(e)
	case *goerror.RequestTimeout:
		return s.Ctx.Status(http.StatusRequestTimeout).JSON(e)
	case *goerror.Conflict:
		return s.Ctx.Status(http.StatusConflict).JSON(e)
	case *goerror.Gone:
		return s.Ctx.Status(http.StatusGone).JSON(e)
	case *goerror.LengthRequired:
		return s.Ctx.Status(http.StatusLengthRequired).JSON(e)
	case *goerror.PreconditionFailed:
		return s.Ctx.Status(http.StatusPreconditionFailed).JSON(e)
	case *goerror.RequestEntityTooLarge:
		return s.Ctx.Status(http.StatusRequestEntityTooLarge).JSON(e)
	case *goerror.RequestURITooLong:
		return s.Ctx.Status(http.StatusRequestURITooLong).JSON(e)
	case *goerror.UnsupportedMediaType:
		return s.Ctx.Status(http.StatusUnsupportedMediaType).JSON(e)
	case *goerror.RequestedRangeNotSatisfiable:
		return s.Ctx.Status(http.StatusRequestedRangeNotSatisfiable).JSON(e)
	case *goerror.ExpectationFailed:
		return s.Ctx.Status(http.StatusExpectationFailed).JSON(e)
	case *goerror.Teapot:
		return s.Ctx.Status(http.StatusTeapot).JSON(e)
	case *goerror.MisdirectedRequest:
		return s.Ctx.Status(http.StatusMisdirectedRequest).JSON(e)
	case *goerror.UnprocessableEntity:
		return s.Ctx.Status(http.StatusUnprocessableEntity).JSON(e)
	case *goerror.Locked:
		return s.Ctx.Status(http.StatusLocked).JSON(e)
	case *goerror.FailedDependency:
		return s.Ctx.Status(http.StatusFailedDependency).JSON(e)
	case *goerror.TooEarly:
		return s.Ctx.Status(http.StatusTooEarly).JSON(e)
	case *goerror.UpgradeRequired:
		return s.Ctx.Status(http.StatusUpgradeRequired).JSON(e)
	case *goerror.PreconditionRequired:
		return s.Ctx.Status(http.StatusPreconditionRequired).JSON(e)
	case *goerror.TooManyRequests:
		return s.Ctx.Status(http.StatusTooManyRequests).JSON(e)
	case *goerror.RequestHeaderFieldsTooLarge:
		return s.Ctx.Status(http.StatusRequestHeaderFieldsTooLarge).JSON(e)
	case *goerror.UnavailableForLegalReasons:
		return s.Ctx.Status(http.StatusUnavailableForLegalReasons).JSON(e)

	// Server error
	case *goerror.InternalServerError:
		return s.Ctx.Status(http.StatusInternalServerError).JSON(e)
	case *goerror.NotImplemented:
		return s.Ctx.Status(http.StatusNotImplemented).JSON(e)
	case *goerror.BadGateway:
		return s.Ctx.Status(http.StatusBadGateway).JSON(e)
	case *goerror.ServiceUnavailable:
		return s.Ctx.Status(http.StatusServiceUnavailable).JSON(e)
	case *goerror.GatewayTimeout:
		return s.Ctx.Status(http.StatusGatewayTimeout).JSON(e)
	case *goerror.HTTPVersionNotSupported:
		return s.Ctx.Status(http.StatusHTTPVersionNotSupported).JSON(e)
	case *goerror.VariantAlsoNegotiates:
		return s.Ctx.Status(http.StatusVariantAlsoNegotiates).JSON(e)
	case *goerror.InsufficientStorage:
		return s.Ctx.Status(http.StatusInsufficientStorage).JSON(e)
	case *goerror.LoopDetected:
		return s.Ctx.Status(http.StatusLoopDetected).JSON(e)
	case *goerror.NotExtended:
		return s.Ctx.Status(http.StatusNotExtended).JSON(e)
	case *goerror.NetworkAuthenticationRequired:
		return s.Ctx.Status(http.StatusNetworkAuthenticationRequired).JSON(e)

	// Other
	default:
		if s.Cus != nil {
			if s.I18n != nil && s.I18n.Enabled && s.I18n.Localize != nil {
				body, e1 := goerror.GetBody(err)
				if e1 == nil && body.Code != "" && body.Message == "" {
					if localize, e2 := s.I18n.Localize(s.Ctx, body.Code); e2 == nil {
						goerror.SetMessage(err, localize)
					}
				}
			}
			return (*s.Cus).Response(s.Ctx, err)
		}
		// Default response
		return s.Ctx.Status(http.StatusBadRequest).JSON(goerror.NewBadRequest())
	}
}

func New(config ...*Config) Response {
	resp := &response{}
	if len(config) > 0 {
		cfg := config[0]
		resp.Cus = cfg.Custom
		resp.I18n = cfg.I18n
	}
	return resp
}
