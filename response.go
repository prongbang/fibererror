package fibererror

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Custom interface {
	Response(Ctx *fiber.Ctx, err error) error
}

type Response interface {
	Response(err error) error
	Custom(cus Custom) Response
}

type response struct {
	Ctx *fiber.Ctx
	Cus Custom
}

// Custom implements Response.
func (r *response) Custom(cus Custom) Response {
	r.Cus = cus
	return r
}

// Response implements Response.
func (s *response) Response(err error) error {
	switch e := err.(type) {
	// Information
	case *Continue:
		return s.Ctx.Status(http.StatusContinue).JSON(e)
	case *SwitchingProtocols:
		return s.Ctx.Status(http.StatusSwitchingProtocols).JSON(e)
	case *Processing:
		return s.Ctx.Status(http.StatusProcessing).JSON(e)

	// Successful
	case *OK:
		return s.Ctx.Status(http.StatusOK).JSON(e)
	case *Created:
		return s.Ctx.Status(http.StatusCreated).JSON(e)
	case *Accepted:
		return s.Ctx.Status(http.StatusAccepted).JSON(e)
	case *NonAuthoritativeInformation:
		return s.Ctx.Status(http.StatusNonAuthoritativeInfo).JSON(e)
	case *NoContent:
		return s.Ctx.Status(http.StatusNoContent).JSON(e)
	case *ResetContent:
		return s.Ctx.Status(http.StatusResetContent).JSON(e)
	case *PartialContent:
		return s.Ctx.Status(http.StatusPartialContent).JSON(e)
	case *MultiStatus:
		return s.Ctx.Status(http.StatusMultiStatus).JSON(e)
	case *AlreadyReported:
		return s.Ctx.Status(http.StatusAlreadyReported).JSON(e)
	case *IMUsed:
		return s.Ctx.Status(http.StatusIMUsed).JSON(e)

	// Redirection
	case *MultipleChoices:
		return s.Ctx.Status(http.StatusMultipleChoices).JSON(e)
	case *MovedPermanently:
		return s.Ctx.Status(http.StatusMovedPermanently).JSON(e)
	case *Found:
		return s.Ctx.Status(http.StatusFound).JSON(e)
	case *SeeOther:
		return s.Ctx.Status(http.StatusSeeOther).JSON(e)
	case *NotModified:
		return s.Ctx.Status(http.StatusNotModified).JSON(e)
	case *UseProxy:
		return s.Ctx.Status(http.StatusUseProxy).JSON(e)
	case *TemporaryRedirect:
		return s.Ctx.Status(http.StatusTemporaryRedirect).JSON(e)
	case *PermanentRedirect:
		return s.Ctx.Status(http.StatusPermanentRedirect).JSON(e)

	// Client error
	case *BadRequest:
		return s.Ctx.Status(http.StatusBadRequest).JSON(e)
	case *Unauthorized:
		return s.Ctx.Status(http.StatusUnauthorized).JSON(e)
	case *PaymentRequired:
		return s.Ctx.Status(http.StatusPaymentRequired).JSON(e)
	case *Forbidden:
		return s.Ctx.Status(http.StatusForbidden).JSON(e)
	case *MethodNotAllowed:
		return s.Ctx.Status(http.StatusMethodNotAllowed).JSON(e)
	case *NotAcceptable:
		return s.Ctx.Status(http.StatusNotAcceptable).JSON(e)
	case *ProxyAuthRequired:
		return s.Ctx.Status(http.StatusProxyAuthRequired).JSON(e)
	case *RequestTimeout:
		return s.Ctx.Status(http.StatusRequestTimeout).JSON(e)
	case *Conflict:
		return s.Ctx.Status(http.StatusConflict).JSON(e)
	case *Gone:
		return s.Ctx.Status(http.StatusGone).JSON(e)
	case *LengthRequired:
		return s.Ctx.Status(http.StatusLengthRequired).JSON(e)
	case *PreconditionFailed:
		return s.Ctx.Status(http.StatusPreconditionFailed).JSON(e)
	case *RequestEntityTooLarge:
		return s.Ctx.Status(http.StatusRequestEntityTooLarge).JSON(e)
	case *RequestURITooLong:
		return s.Ctx.Status(http.StatusRequestURITooLong).JSON(e)
	case *UnsupportedMediaType:
		return s.Ctx.Status(http.StatusUnsupportedMediaType).JSON(e)
	case *RequestedRangeNotSatisfiable:
		return s.Ctx.Status(http.StatusRequestedRangeNotSatisfiable).JSON(e)
	case *ExpectationFailed:
		return s.Ctx.Status(http.StatusExpectationFailed).JSON(e)
	case *Teapot:
		return s.Ctx.Status(http.StatusTeapot).JSON(e)
	case *MisdirectedRequest:
		return s.Ctx.Status(http.StatusMisdirectedRequest).JSON(e)
	case *UnprocessableEntity:
		return s.Ctx.Status(http.StatusUnprocessableEntity).JSON(e)
	case *Locked:
		return s.Ctx.Status(http.StatusLocked).JSON(e)
	case *FailedDependency:
		return s.Ctx.Status(http.StatusFailedDependency).JSON(e)
	case *TooEarly:
		return s.Ctx.Status(http.StatusTooEarly).JSON(e)
	case *UpgradeRequired:
		return s.Ctx.Status(http.StatusUpgradeRequired).JSON(e)
	case *PreconditionRequired:
		return s.Ctx.Status(http.StatusPreconditionRequired).JSON(e)
	case *TooManyRequests:
		return s.Ctx.Status(http.StatusTooManyRequests).JSON(e)
	case *RequestHeaderFieldsTooLarge:
		return s.Ctx.Status(http.StatusRequestHeaderFieldsTooLarge).JSON(e)
	case *UnavailableForLegalReasons:
		return s.Ctx.Status(http.StatusUnavailableForLegalReasons).JSON(e)

	// Server error
	case *InternalServerError:
		return s.Ctx.Status(http.StatusInternalServerError).JSON(e)
	case *NotImplemented:
		return s.Ctx.Status(http.StatusNotImplemented).JSON(e)
	case *BadGateway:
		return s.Ctx.Status(http.StatusBadGateway).JSON(e)
	case *ServiceUnavailable:
		return s.Ctx.Status(http.StatusServiceUnavailable).JSON(e)
	case *GatewayTimeout:
		return s.Ctx.Status(http.StatusGatewayTimeout).JSON(e)
	case *HTTPVersionNotSupported:
		return s.Ctx.Status(http.StatusHTTPVersionNotSupported).JSON(e)
	case *VariantAlsoNegotiates:
		return s.Ctx.Status(http.StatusVariantAlsoNegotiates).JSON(e)
	case *InsufficientStorage:
		return s.Ctx.Status(http.StatusInsufficientStorage).JSON(e)
	case *LoopDetected:
		return s.Ctx.Status(http.StatusLoopDetected).JSON(e)
	case *NotExtended:
		return s.Ctx.Status(http.StatusNotExtended).JSON(e)
	case *NetworkAuthenticationRequired:
		return s.Ctx.Status(http.StatusNetworkAuthenticationRequired).JSON(e)

	// Other
	default:
		if s.Cus != nil {
			return s.Cus.Response(s.Ctx, err)
		}
		// Default response
		return s.Ctx.Status(http.StatusBadRequest).JSON(NewBadRequest())
	}
}

func New(c *fiber.Ctx) Response {
	return &response{
		Ctx: c,
	}
}
