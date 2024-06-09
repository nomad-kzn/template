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

var _ interfaces.Usecase = (*UCImpl)(nil)

type UCImpl struct {
	r   interfaces.Repository
	s   interfaces.Srvc
	cfg configs.ServiceCfg
	l   logger.Logger
}

func NewUCImpl(d *deps.Deps, db *sql.DB) (*UCImpl, error) {
	if d == nil {
		return nil, fmt.Errorf("usecase init err: deps is nil")
	}

	service, err := srvc.NewServiceImpl(d)
	if err != nil {
		return nil, err
	}

	return &UCImpl{
		r:   repository.NewUserRepoImpl(db),
		s:   service,
		cfg: d.Configs,
		l:   d.Logger,
	}, nil
}

func (u *UCImpl) GetUser(ctx context.Context, id string) (*models.User, error) {
	u.l.Info("it is ok!")
	result, err := u.r.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return result.Convert(), nil
}
