// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"fairseller-backend/config"
	v1 "fairseller-backend/internal/controller/http/v1"
	"fairseller-backend/internal/repository"
	"fairseller-backend/internal/usecase"
	"fairseller-backend/pkg/httpserver"
	"fairseller-backend/pkg/logger"
	"fairseller-backend/pkg/postgres"

	"github.com/gin-gonic/gin"
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
	authUseCase := usecase.NewAuthUseCase(
		repositoryContainer.UserRepository,
		l,
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, authUseCase)
	server := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-server.Notify():
		l.Error(fmt.Errorf("app - Run - httpserver.Notify: %w", err))
	}

	// Shutdown
	err = server.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpserver.Shutdown: %w", err))
	}
}
