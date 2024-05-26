package deps

import (
	"github.com/nomad-kzn/template/internal/configs"
	"github.com/nomad-kzn/template/pkg/logger"
)

type Deps struct {
	Logger logger.Logger
	Configs configs.ServiceCfg
}

func NewDeps(l logger.Logger, cfg configs.ServiceCfg) *Deps {
	return &Deps{
		Logger: l,
		Configs: cfg,
	}
}