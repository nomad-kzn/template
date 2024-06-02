package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nomad-kzn/template/internal/configs"
	"github.com/nomad-kzn/template/internal/deps"
	"github.com/nomad-kzn/template/internal/interfaces"
	"github.com/nomad-kzn/template/internal/models"
	"github.com/nomad-kzn/template/internal/repository"
	"github.com/nomad-kzn/template/internal/srvc"
	"github.com/nomad-kzn/template/pkg/logger"
)

var _ interfaces.Usecase = (*UsecaseImpl)(nil)

type UsecaseImpl struct {
	r   interfaces.Repository
	s interfaces.Srvc
	cfg configs.ServiceCfg
	l   logger.Logger
}

func NewUCImpl(d *deps.Deps, db *sql.DB) (*UsecaseImpl, error) {
	if d == nil {
		return nil, fmt.Errorf("usecase init err: deps is nil")
	}

	srvc, err := srvc.NewSrvcImpl(d)
	if err != nil {
		return nil, err
	}

	return &UsecaseImpl{
		r:   repository.NewUserRepoImpl(db),
		s: srvc,
		cfg: d.Configs,
		l:   d.Logger,
	}, nil
}

func (u *UsecaseImpl) GetUser(ctx context.Context, id string) (*models.User, error) {
	u.l.Info("it is ok!")
	result, err := u.r.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return result.Convert(), nil
}
