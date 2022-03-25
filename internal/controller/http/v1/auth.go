package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fairseller-backend/internal/entity"
	"fairseller-backend/internal/usecase"
	"fairseller-backend/pkg/logger"
)

type authRoutes struct {
	authUseCase usecase.Auth
	logger      logger.Interface
}

func newAuthRoutes(handler *gin.RouterGroup, authUseCase usecase.Auth, l logger.Interface) {
	routes := &authRoutes{authUseCase, l}

	authHandler := handler.Group("/auth")
	authHandler.POST("/sign-up-request", routes.signUpRequest)
}

// @Summary     Sign up
// @Description Sign up first step with send code to mobile
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       data body SignUpRequestDto true "Phone for getting code"
// @Success     200
// @Failure     500 {object} response
// @Router      /auth/sign-up-request [post]
func (r *authRoutes) signUpRequest(c *gin.Context) {
	body := SignUpRequestDto{}
	if err := c.BindJSON(&body); err != nil {
		r.logger.Error(err, "http - v1 - auth - signUpRequest")
		validationErrorResponse(err, c)

		return
	}

	dto := entity.SignUpRequest{
		Phone: body.Phone,
	}

	if err := r.authUseCase.SignUpRequest(c.Request.Context(), dto); err != nil {
		r.logger.Error(err, "http - v1 - auth - signUpRequest")
		errorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.Status(http.StatusOK)
}
