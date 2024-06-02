package srvc

import (
	"context"
	"fmt"

	"github.com/nomad-kzn/template/internal/configs"
	"github.com/nomad-kzn/template/internal/deps"
	"github.com/nomad-kzn/template/internal/interfaces"
	"github.com/nomad-kzn/template/internal/models"
	"github.com/nomad-kzn/template/pkg/logger"
)

var _ interfaces.Srvc = (*SrvcImpl)(nil)

type SrvcImpl struct {
	l   logger.Logger
	cfg configs.ServiceCfg
}

func NewSrvcImpl(d *deps.Deps) (*SrvcImpl, error){
	if d == nil {
		return nil, fmt.Errorf("srvc init err: deps is nil")
	}

	return &SrvcImpl{
		cfg: d.Configs,
		l:   d.Logger,
	}, nil
}

func (s *SrvcImpl) GetUsersStats(ctx context.Context) ([]models.UserStat, error) {
	return []models.UserStat{}, nil
}