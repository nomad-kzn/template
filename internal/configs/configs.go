package configs

import (
	"github.com/nomad-kzn/template/pkg/db"
	"github.com/nomad-kzn/template/pkg/logger"
	"github.com/spf13/viper"
)

const (
	envTimeoutSeconds = "TIMEOUT_SECONDS"
	envLogLvl         = "LOG_LVL"
	envEnvironment    = "ENV"

	defaultTimeoutSeconds = 30
	defaultLogLvl         = "INFO"
	defaultEnv            = "test"
)

type (
	Cfg struct {
		Env            string
		LoggerCfg      *logger.LoggerCfg
		DBCfg          string
		ServiceCfg     *ServiceCfg
		DatabaseConfig *db.DatabaseConfig
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

func NewCfg() *Cfg {
	v := viper.New()
	v.SetDefault(envEnvironment, defaultEnv)

	return &Cfg{
		Env:            v.GetString(envEnvironment),
		LoggerCfg:      NewLoggerCfg(v),
		DBCfg:          "",
		ServiceCfg:     NewServiceCfg(v),
		DatabaseConfig: db.NewDatabaseConfig(v),
	}
}
