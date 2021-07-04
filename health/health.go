package libhealth

import (
	"net/http"

	"github.com/labstack/echo"
)

func NewServer() *Server {
	return &Server{}
}

type Server struct {}

func (s *Server) HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}

func (s *Server) ReadinessHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}