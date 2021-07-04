package libmw

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	libcommon "github.com/machmum/lib-simple/common"
	liblog "github.com/machmum/lib-simple/log"
)

func NewMiddleware(opt libcommon.Options) CommonMiddleware {
	return &mw{opt: opt}
}

// mw middleware contains jwt pub key
type mw struct {
	opt libcommon.Options
}

func (m *mw) CORS() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(libcommon.HeaderAllowOrigin, "*")
			c.Response().Header().Set(libcommon.HeaderAllowMethod, strings.Join(libcommon.AllowedMethods, ","))
			c.Response().Header().Set(libcommon.HeaderAllowHeaders, strings.Join(libcommon.AllowedHeaders, ","))

			if c.Request().Method == http.MethodOptions {
				_, _ = c.Response().Write([]byte("ok"))
			}
			return next(c)
		}
	}
}

// AccessLog: log session's email & basic request properties data
func (m *mw) AccessLog() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqID := c.Response().Header().Get(echo.HeaderXRequestID)
			defer func() {
				rq := c.Request()
				rs := c.Response()
				liblog.Access(libcommon.NewRequestID(reqID), rq, rs)
			}()
			return next(c)
		}
	}
}

func (m *mw) StdBasicAuthValidator() echo.MiddlewareFunc {
	return middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: func(username, password string, c echo.Context) (valid bool, err error) {
			if username == m.opt.Basic.Username &&
				password == m.opt.Basic.Password {
				valid = true
			}
			return
		},
	})
}
