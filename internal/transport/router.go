package transport

import (
	"github.com/labstack/echo/v4"
	"github.com/nomad-kzn/template/internal/deps"
	"github.com/nomad-kzn/template/internal/interfaces"
	"net/http"

	echoMw "github.com/labstack/echo/v4/middleware"
)

const (
	defaultEchoRecoverStackSize  = 4
	defaultEchoRecoverStackShift = 10
	defaultEchoRecoverLogLvl     = 4
)

func BindRoutes(useCase interfaces.Usecase, deps deps.Deps, e *echo.Echo) {
	h := NewHandler(useCase, deps)

	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.Use(echoMw.RecoverWithConfig(echoMw.RecoverConfig{
		Skipper:   echoMw.DefaultSkipper,
		StackSize: defaultEchoRecoverStackSize << defaultEchoRecoverStackShift,
		LogLevel:  defaultEchoRecoverLogLvl,
	}))

	e.GET("/ping", h.Ping)
	e.GET("/routes", func(c echo.Context) error {
		return c.JSON(http.StatusOK, e.Routes())
	})
}
