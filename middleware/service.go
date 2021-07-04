package libmw

import "github.com/labstack/echo"

// DefaultMiddleware will always be applied on new instance
type DefaultMiddleware interface {
	CORS() echo.MiddlewareFunc
	AccessLog() echo.MiddlewareFunc
}

type AuthMiddleware interface {
	// applied with epmsrv.WithBasicAuth
	StdBasicAuthValidator() echo.MiddlewareFunc
}

type CommonMiddleware interface {
	DefaultMiddleware
	AuthMiddleware
}
