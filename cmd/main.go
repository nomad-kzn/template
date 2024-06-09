package main

import (
	"log"

	"github.com/nomad-kzn/template/internal/app"
	"github.com/nomad-kzn/template/internal/configs"
)

func main() {
	cfg := configs.NewCfg()
	application, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal("app init err: ", err)
	}

	application.Run()
}
