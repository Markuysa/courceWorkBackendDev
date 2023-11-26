package server

import (
	"context"
	"log"

	"github.com/Markuysa/courceWorkBackendDev/config"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type App struct {
	cfg    config.Config
	tracer trace.Tracer
	logger *zap.Logger
	app    *fiber.App
}

func New(
	cfg config.Config,
	tracer trace.Tracer,
	logger *zap.Logger,
) *App {
	return &App{
		app:    fiber.New(),
		cfg:    cfg,
		tracer: tracer,
		logger: logger,
	}
}

func (a App) Start(ctx context.Context) error {

	err := a.MapHandlers()
	if err != nil {
		return err
	}

	if err = a.app.Listen(a.cfg.HTTP.URI); err != nil {
		log.Fatal(err)
	}

	return nil
}
