package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "fairseller-backend/internal/entity"
	"fairseller-backend/internal/useCase"
	"fairseller-backend/pkg/logger"
)

type authRoutes struct {
	authUseCase useCase.Auth
	logger      logger.Interface
}

func newAuthRoutes(handler *gin.RouterGroup, authUseCase useCase.Auth, logger logger.Interface) {
	routes := &authRoutes{authUseCase, logger}

	authHandler := handler.Group("/auth")
	authHandler.POST("/sign-up", routes.signUp)
}

// @Summary     Sign up
// @Description Sign up first step with send code to mobile
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       request body entity.SignUpRequestDto true "Phone for getting code"
// @Success     200
// @Failure     500 {object} response
// @Router      /auth/sign-up [post]
func (r *authRoutes) signUp(c *gin.Context) {
	// translations, err := r.userUseCase.GetSignUpCode(c.Request.Context())
	// if err != nil {
	// 	r.logger.Error(err, "http - v1 - auth")
	// 	errorResponse(c, http.StatusInternalServerError, "database problems")

	// 	return
	// }

	c.JSON(http.StatusOK, 200)
}
