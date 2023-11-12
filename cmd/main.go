package main

import (
	"context"
	"log"

	"github.com/Markuysa/courceWorkBackendDev/internal/app"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := app.NewApp(config)

	if err = app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
