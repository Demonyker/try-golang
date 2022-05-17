package entity

import "github.com/gin-gonic/gin"

type Logger interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Error(message string, err error)
	GatewayError(err error)
	DatabaseError(err error)
	UseCaseError(err error)
	ServerRequestInfo(c *gin.Context) error
	ServerResponseInfo(c *gin.Context)
}
