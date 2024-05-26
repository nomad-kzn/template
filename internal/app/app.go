package app

import (
	"fmt"
	"github.com/nomad-kzn/template/internal/configs"
	"github.com/nomad-kzn/template/internal/deps"
	"github.com/nomad-kzn/template/internal/interfaces"
	"github.com/nomad-kzn/template/internal/usecase"
	"github.com/nomad-kzn/template/pkg/logger"

	"context"
)

type App struct {
	u interfaces.Usecase
}

func NewApp(cfg *configs.Cfg) (*App, error) {
	if cfg == nil {
		return nil, fmt.Errorf("app init err: config is nil")
	}

	l := logger.NewLoggerImpl(cfg.LoggerCfg)
	d := deps.NewDeps(l, *cfg.ServiceCfg)

	uc, err := usecase.NewUCImpl(d)
	if err != nil {
		return nil, err
	}

	return &App{u: uc}, nil
}

func (a *App) Run(ctx context.Context) error {
	_, err := a.u.GetUser(ctx, "")
	return err
}
