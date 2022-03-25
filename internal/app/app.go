// Package app configures and runs application.
package app

import (
	"fairseller-backend/pkg/httpServer"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"fairseller-backend/config"
	v1 "fairseller-backend/internal/controller/http/v1"
	"fairseller-backend/internal/repository"
	"fairseller-backend/internal/useCase"
	"fairseller-backend/pkg/logger"
	"fairseller-backend/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	repositoryContainer := repository.New(pg)
	// Use case
	authUseCase := useCase.NewAuthUseCase(
		repositoryContainer.UserRepository,
		l,
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, authUseCase)
	server := httpServer.New(handler, httpServer.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-server.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = server.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
