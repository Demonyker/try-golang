package v1

import (
	"net/http"

	"fairseller-backend/pkg/logger"

	"github.com/gin-gonic/gin"
)

type authRoutes struct {
	authUseCase AuthUseCase
	logger      logger.Interface
}

func newAuthRoutes(handler *gin.RouterGroup, authUseCase AuthUseCase, l logger.Interface) {
	routes := &authRoutes{authUseCase, l}

	authHandler := handler.Group("/auth")
	authHandler.POST("/sign-up-request", routes.signUpRequest)
}

// @Summary     Sign up
// @Description Sign up first step with send code to mobile
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       data body signUpRequestBody true "Phone for getting code"
// @Success     200
// @Failure     500 {object} response
// @Router      /auth/sign-up-request [post]
func (r *authRoutes) signUpRequest(c *gin.Context) {
	body := signUpRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		r.logger.Error(err, "http - v1 - auth - signUpRequest")
		validationErrorResponse(err, c)

		return
	}

	if err := r.authUseCase.SignUpRequest(c.Request.Context(), body.Phone); err != nil {
		r.logger.Error(err, "http - v1 - auth - signUpRequest")
		errorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.Status(http.StatusOK)
}
