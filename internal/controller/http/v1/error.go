package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type response struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{msg})
}

func getValidationErrorMsg(fe validator.FieldError) string {
	value := fmt.Sprintf("%v", fe.Value())
	tag := fe.Tag()

	switch tag {
	case "required":
		return fe.Field() + " field is required"
	case "email":
		return fe.Field() + " must be valid email"
	case "e164":
		return value + " is invalid phone number"
	}

	return fe.Field() + " is unknown error"
}

func validationErrorResponse(err error, context *gin.Context) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]string, len(ve))
		for i, fe := range ve {
			out[i] = getValidationErrorMsg(fe)
		}

		errorResponse(context, http.StatusBadRequest, out[0])
	}
}
