package main

import (
	"context"
	"log"

	"github.com/Markuysa/courceWorkBackendDev/config"
	"github.com/Markuysa/courceWorkBackendDev/internal/server"
	"github.com/Markuysa/courceWorkBackendDev/utils/logger"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
	"github.com/joho/godotenv"
)

func main() {
	lg := logger.InitLogger()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	tracer, err := oteltrace.InitTracer(cfg.Trace.URL, cfg.Trace.ServiceName)
	if err != nil {
		log.Fatal(err)
	}

	app := server.New(cfg, tracer, lg)

	if err = app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
