package libserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	libcommon "github.com/machmum/lib-simple/common"
	libctx "github.com/machmum/lib-simple/context"
	liblog "github.com/machmum/lib-simple/log"
	libmw "github.com/machmum/lib-simple/middleware"
)

// NewStd return new server with default log
func NewStd(opt libcommon.Options) Server {
	liblog.NewStd()
	return New(opt)
}

// New returns new server with default middleware
func New(opt libcommon.Options) Server {
	s := &server{
		e:   echo.New(),
		opt: opt,
		mw:  libmw.NewMiddleware(opt),
	}
	s.e.HideBanner = true
	s.e.Use(
		// Recover from panics,
		// see: https://echo.labstack.com/middleware/recover
		middleware.Recover(),

		// Request ID middleware generates a unique id for a request
		middleware.RequestID(),

		// CORS handler
		s.mw.CORS(),

		// AccessLog print information about each HTTP request
		s.mw.AccessLog(),
	)
	return s
}

type server struct {
	e   *echo.Echo
	opt libcommon.Options
	mw  libmw.CommonMiddleware
}

func (s *server) Echo() *echo.Echo {
	return s.e
}

func (s *server) Start() {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", s.opt.ServerAddr),
	}

	// Start server
	go func() {
		err := s.e.StartServer(srv)
		if err != nil {
			fmt.Printf("[server] Failed to start: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	deadline := 5 * time.Second
	fmt.Println("[server] Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), deadline)
	defer cancel()
	err := s.e.Shutdown(ctx)
	if err != nil {
		s.e.Logger.Fatal(err)
	}
}

func (s *server) WithBasicAuth() echo.MiddlewareFunc {
	return s.mw.StdBasicAuthValidator()
}

// NewHttpContext will add epmctx.RequestID to context from echo.HeaderXRequestID
func NewHttpContext(c echo.Context) context.Context {
	reqID := c.Response().Header().Get(echo.HeaderXRequestID)
	return libctx.NewRequestIDContext(
		c.Request().Context(), libcommon.NewRequestID(reqID),
	)
}
