package v1

import (
	"net/http"

	"fairseller-backend/internal/entity"

	"github.com/gin-gonic/gin"
)

type authRoutes struct {
	authUseCase AuthUseCase
	logger      entity.Logger
}

func newAuthRoutes(handler *gin.RouterGroup, authUseCase AuthUseCase, l entity.Logger) {
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
		r.logger.GatewayError(err)
		validationErrorResponse(err, c)

		return
	}

	if err := r.authUseCase.SignUpRequest(body.Phone); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.Status(http.StatusOK)
}
