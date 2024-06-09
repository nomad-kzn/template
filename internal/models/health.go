package models

import (
	"os"
	"time"
)

type Health struct {
	Data    string `json:"data"`
	Version string `json:"version"`
}

func CurrentHealth() Health {
	return Health{
		Data:    time.Now().Format("02.01.2006 15:04:05"),
		Version: os.Getenv("SERVICE_VERSION"),
	}
}
