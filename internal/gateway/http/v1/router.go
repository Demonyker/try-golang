// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "fairseller-backend/docs"
	"fairseller-backend/internal/entity"
)

// NewRouter -.
// Swagger spec:
// @title       Fairseller API
// @description Api for fairseller APP
// @version     1.0
// @host        localhost:3000
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l entity.Logger, authUseCase AuthUseCase) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(func(c *gin.Context) {
		if err := l.ServerRequestInfo(c); err != nil {
			errorResponse(c, http.StatusInternalServerError, err.Error())
		}
	})
	handler.Use(l.ServerResponseInfo)

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	v1Handler := handler.Group("/v1")

	// RESPONSE EWE
	newAuthRoutes(v1Handler, authUseCase, l)
}
