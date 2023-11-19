package http

import "github.com/Markuysa/courceWorkBackendDev/internal/client/usecase"

type Handlers interface {
}

type ClientHandlers struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) Handlers {
	return &ClientHandlers{
		uc: uc,
	}
}
