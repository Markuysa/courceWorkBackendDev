package app

import (
	"context"

	"github.com/Markuysa/courceWorkBackendDev/config"
	"github.com/Markuysa/courceWorkBackendDev/internal/usecase"
)

type App struct {
	cfg config.Config
}

func New(cfg config.Config) *App {
	return &App{cfg: cfg}
}

func (a App) Start(ctx context.Context) error {
	uc := usecase.New(nil, nil)

	return nil
}
