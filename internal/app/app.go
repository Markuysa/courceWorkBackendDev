package app

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/config"
)

type App struct {
	cfg config.Config
}

func New(cfg config.Config) *App {
	return &App{cfg: cfg}
}

func (a App) Start(ctx context.Context) error {

	return nil
}
