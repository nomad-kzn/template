package app

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nomad-kzn/template/internal/configs"
	"github.com/nomad-kzn/template/internal/deps"
	"github.com/nomad-kzn/template/internal/transport"
	"github.com/nomad-kzn/template/internal/usecase"
	"github.com/nomad-kzn/template/pkg/db"
	"github.com/nomad-kzn/template/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"time"

	"context"
)

type App struct {
	e *echo.Echo
}

func NewApp(cfg *configs.Cfg) (*App, error) {
	if cfg == nil {
		return nil, fmt.Errorf("app init err: config is nil")
	}

	l := logger.NewLoggerImpl(cfg.LoggerCfg)
	d := deps.NewDeps(l, *cfg.ServiceCfg)

	database := db.NewPostgresClient(cfg.DatabaseConfig)

	if err := database.Connect(context.Background()); err != nil {
		l.Error("NewApp init err: failed to connect database", err)
		return nil, err
	}

	uc, err := usecase.NewUCImpl(d, database.Database())
	if err != nil {
		return nil, err
	}

	e := echo.New()
	transport.BindRoutes(uc, *d, e)

	return &App{e: e}, nil
}

func (a *App) Run() {
	go func() {
		if err := a.e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	a.e.Logger.Info("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.e.Shutdown(ctx); err != nil {
		a.e.Logger.Fatal(err)
	}
}
