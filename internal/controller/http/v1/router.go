// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"fairseller-backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "fairseller-backend/docs"
	"fairseller-backend/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Fairseller API
// @description Api for fairseller APP
// @version     1.0
// @host        localhost:3000
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, authUseCase usecase.Auth) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	v1Handler := handler.Group("/v1")
	newAuthRoutes(v1Handler, authUseCase, l)
}
