package db

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	SSLMode  string
}

const (
	dbDefaultUsername = "postuser"
	dbDefaultPassword = "postgres"
	dbDefaultHost     = "localhost"
	dbDefaultPort     = "5432"
	dbDefaultDatabase = "portal"
	dbDefaultSslMode  = "disable"

	envDatabaseUsername = "DATABASE_USERNAME"
	envDatabasePassword = "DATABASE_PASSWORD"
	envDatabaseHost     = "DATABASE_HOST"
	envDatabasePort     = "DATABASE_PORT"
	envDatabaseName     = "DATABASE_NAME"
	envDatabaseSSLMode  = "DATABASE_SSLMODE"
)

func NewDatabaseConfig(v *viper.Viper) *DatabaseConfig {
	v.SetDefault(envDatabaseUsername, dbDefaultUsername)
	v.SetDefault(envDatabasePassword, dbDefaultPassword)
	v.SetDefault(envDatabaseHost, dbDefaultHost)
	v.SetDefault(envDatabasePort, dbDefaultPort)
	v.SetDefault(envDatabaseName, dbDefaultDatabase)
	v.SetDefault(envDatabaseSSLMode, dbDefaultSslMode)

	return &DatabaseConfig{
		Username: v.GetString(envDatabaseUsername),
		Password: v.GetString(envDatabasePassword),
		Host:     v.GetString(envDatabaseHost),
		Port:     v.GetString(envDatabasePort),
		Database: v.GetString(envDatabaseName),
		SSLMode:  v.GetString(envDatabaseSSLMode),
	}
}

func (d DatabaseConfig) URI() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Database,
		d.SSLMode,
	)
}
