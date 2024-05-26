package configs

import (
	"github.com/nomad-kzn/template/pkg/logger"
	"github.com/spf13/viper"
)

const (
	envTimeoutSeconds                 = "TIMEOUT_SECONDS"
	envLogLvl = "LOG_LVL"

	defaultTimeoutSeconds = 30
	defaultLogLvl = "INFO"
)

type (
	Cfg struct {
		Env string
		LoggerCfg *logger.LoggerCfg
		DBCfg string
		ServiceCfg *ServiceCfg
	}

	ServiceCfg struct {
		TimeoutSeconds int
	}
)

func NewServiceCfg(v *viper.Viper) *ServiceCfg {
	v.SetDefault(envTimeoutSeconds, defaultTimeoutSeconds)

	return &ServiceCfg{
		TimeoutSeconds: v.GetInt(envTimeoutSeconds),
	}
}

func NewLoggerCfg(v *viper.Viper) *logger.LoggerCfg {
	v.SetDefault(envLogLvl, defaultLogLvl)

	return &logger.LoggerCfg{
		Lvl: v.GetString(envLogLvl),
	}
}