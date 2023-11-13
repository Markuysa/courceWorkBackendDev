package usecase

import (
	"github.com/Markuysa/courceWorkBackendDev/config"
	"go.opentelemetry.io/otel/trace"
)

type UC struct {
	cfg    *config.Config
	tracer trace.Tracer
}

func New(
	cfg *config.Config,
	tracer trace.Tracer,
) Usecase {
	return &UC{
		cfg:    cfg,
		tracer: tracer,
	}
}
