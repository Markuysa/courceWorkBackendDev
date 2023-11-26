package server

import (
	adminDelivery "github.com/Markuysa/courceWorkBackendDev/internal/admin/delivery/http"
	adminRepo "github.com/Markuysa/courceWorkBackendDev/internal/admin/repository"
	adminUC "github.com/Markuysa/courceWorkBackendDev/internal/admin/usecase"
	"github.com/Markuysa/courceWorkBackendDev/internal/auth/cache"
	authDelivery "github.com/Markuysa/courceWorkBackendDev/internal/auth/delivery/http"
	authRepo "github.com/Markuysa/courceWorkBackendDev/internal/auth/repository"
	authUC "github.com/Markuysa/courceWorkBackendDev/internal/auth/usecase"
	clientDelivery "github.com/Markuysa/courceWorkBackendDev/internal/client/delivery/http"
	clientRepo "github.com/Markuysa/courceWorkBackendDev/internal/client/repository"
	clientUC "github.com/Markuysa/courceWorkBackendDev/internal/client/usecase"
	"github.com/Markuysa/courceWorkBackendDev/internal/middleware"
	"github.com/Markuysa/courceWorkBackendDev/utils/pgconn"
	"github.com/Markuysa/courceWorkBackendDev/utils/redisconnector"
)

func (a App) MapHandlers() error {
	pgRepo := pgconn.New(a.cfg.Postgres)
	redisConn := redisconnector.New(a.cfg.Redis)

	clientRepos := clientRepo.New(pgRepo)
	adminRepos := adminRepo.New(pgRepo)
	authRepos := authRepo.New(pgRepo)

	sessionCache := cache.New(redisConn, a.cfg)

	clientUseCase := clientUC.New(a.cfg, clientRepos)
	adminUseCase := adminUC.New(a.cfg, adminRepos)
	authUseCase := authUC.New(a.cfg, sessionCache, authRepos)

	mw := middleware.New(sessionCache)

	clientHandlers := clientDelivery.New(clientUseCase)
	adminHandlers := adminDelivery.New(adminUseCase)
	authHandlers := authDelivery.New(authUseCase)

	clientGroup := a.app.Group("/client")
	clientDelivery.MapClientRoutes(clientGroup, mw, clientHandlers)

	adminGroup := a.app.Group("/admin")
	adminDelivery.MapAdminRoutes(adminGroup, mw, adminHandlers)

	authGroup := a.app.Group("/auth")
	authDelivery.MapAuthRoutes(authGroup, mw, authHandlers)

	return nil
}
