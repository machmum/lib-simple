package libserver

import "github.com/labstack/echo"

type ServerAuth interface {
	WithBasicAuth() echo.MiddlewareFunc
}

type Server interface {
	Echo() *echo.Echo
	Start()
	ServerAuth
}
