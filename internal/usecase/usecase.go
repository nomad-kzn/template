package usecase

import (
	"context"
	"fmt"

	"github.com/nomad-kzn/template/internal/configs"
	"github.com/nomad-kzn/template/internal/deps"
	"github.com/nomad-kzn/template/internal/interfaces"
	"github.com/nomad-kzn/template/internal/models"
	"github.com/nomad-kzn/template/pkg/logger"
)

var _ interfaces.Usecase = (*UsecaseImpl)(nil)

type UsecaseImpl struct {
	r   string //place her repository
	cfg configs.ServiceCfg
	l   logger.Logger
}

func NewUCImpl(d *deps.Deps) (*UsecaseImpl, error) {
	if d == nil {
		return nil, fmt.Errorf("usecase init err: deps is nil")
	}

	return &UsecaseImpl{
		r:   "something",
		cfg: d.Configs,
		l:   d.Logger,
	}, nil
}

func (u *UsecaseImpl) GetUser(ctx context.Context, id string) (*models.User, error) {
	u.l.Info("it is ok!")
	return &models.User{}, nil
}
