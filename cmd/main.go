package main

import (
	"context"
	"log"

	"github.com/Markuysa/courceWorkBackendDev/config"
	"github.com/Markuysa/courceWorkBackendDev/internal/app"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := app.New(config)

	if err = app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
