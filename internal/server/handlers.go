package server

import (
	adminDelivery "github.com/Markuysa/courceWorkBackendDev/internal/admin/delivery/http"
	adminRepo "github.com/Markuysa/courceWorkBackendDev/internal/admin/repository"
	adminUC "github.com/Markuysa/courceWorkBackendDev/internal/admin/usecase"
	authDelivery "github.com/Markuysa/courceWorkBackendDev/internal/auth/delivery/http"
	authRepo "github.com/Markuysa/courceWorkBackendDev/internal/auth/repository"
	authUC "github.com/Markuysa/courceWorkBackendDev/internal/auth/usecase"
	clientDelivery "github.com/Markuysa/courceWorkBackendDev/internal/client/delivery/http"
	clientRepo "github.com/Markuysa/courceWorkBackendDev/internal/client/repository"
	clientUC "github.com/Markuysa/courceWorkBackendDev/internal/client/usecase"
	"github.com/Markuysa/courceWorkBackendDev/internal/middleware"
	"github.com/Markuysa/courceWorkBackendDev/utils/pgconnector"
)

func (a App) MapHandlers() error {
	pgRepo := pgconnector.New(a.cfg.Postgres)

	clientRepo := clientRepo.New(pgRepo)
	adminRepo := adminRepo.New(pgRepo)
	authRepo := authRepo.New(pgRepo)

	clientUC := clientUC.New(a.cfg, clientRepo)
	adminUC := adminUC.New(a.cfg, adminRepo)
	authUC := authUC.New(a.cfg, authRepo)

	mw := middleware.New()

	if err := a.MapHandlers(); err != nil {
		return err
	}

	clientHandlers := clientDelivery.New(clientUC)
	adminHandlers := adminDelivery.New(adminUC)
	authHandlers := authDelivery.New(authUC)

	clientGroup := a.app.Group("/client")
	clientDelivery.MapClientRoutes(clientGroup, mw, clientHandlers)

	adminGroup := a.app.Group("/admin")
	adminDelivery.MapAdminRoutes(adminGroup, mw, adminHandlers)

	authGroup := a.app.Group("/auth")
	authDelivery.MapAuthRoutes(authGroup, mw, authHandlers)

	return nil
}
