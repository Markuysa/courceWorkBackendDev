package server

import (
	"context"
	"log"

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
	tgRepo "github.com/Markuysa/courceWorkBackendDev/internal/telegram/repository"
	tgUC "github.com/Markuysa/courceWorkBackendDev/internal/telegram/usecase"
	"github.com/Markuysa/courceWorkBackendDev/utils/pgconn"
	"github.com/Markuysa/courceWorkBackendDev/utils/redisconnector"
)

func (a App) MapHandlers() error {
	pgRepo := pgconn.New(a.cfg.Postgres)
	redisConn := redisconnector.New(a.cfg.Redis)

	clientRepos := clientRepo.New(pgRepo)
	adminRepos := adminRepo.New(pgRepo)
	authRepos := authRepo.New(pgRepo)
	tgRepos := tgRepo.New(pgRepo)

	sessionCache := cache.New(redisConn, a.cfg)

	clientUseCase := clientUC.New(a.cfg, clientRepos)
	adminUseCase := adminUC.New(a.cfg, adminRepos)
	authUseCase := authUC.New(a.cfg, sessionCache, authRepos)
	tgUSecase := tgUC.New(a.cfg, tgRepos)

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

	go func() {
		tgUSecase.StartWorker(context.Background())
	}()

	go func() {
		if err := tgUSecase.ListenMessages(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}
