package transport

import (
	"github.com/labstack/echo/v4"
	"github.com/nomad-kzn/template/internal/models"
	"net/http"

	"github.com/nomad-kzn/template/pkg/logger"

	d "github.com/nomad-kzn/template/internal/deps"
	i "github.com/nomad-kzn/template/internal/interfaces"
)

type Handler struct {
	lg      logger.Logger
	useCase i.Usecase
}

func NewHandler(useCase i.Usecase, deps d.Deps) *Handler {
	return &Handler{
		lg:      deps.Logger,
		useCase: useCase,
	}
}

func (h *Handler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, models.CurrentHealth())
}
