package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

type Logger struct {
	logger *logrus.Logger
}

func New(file io.Writer) *Logger {
	logger := logrus.New()
	logger.SetFormatter(&ecslogrus.Formatter{})
	logger.SetOutput(file)

	return &Logger{
		logger,
	}
}

func (l *Logger) Debug(message string, args ...interface{}) {
	l.logger.Debug(message, args)
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.logger.Info(message, args)
}

func (l *Logger) Error(message string, err error) {
	l.logger.WithError(err).Error(message)
}

func (l *Logger) GatewayError(err error) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	isCanShowCaller := ok && details != nil

	if isCanShowCaller {
		l.logger.WithField("caller", details.Name()).WithField("layer", "presenter").WithError(err).Error("Error in gateway")
	} else {
		l.Error("Error in gateway", err)
	}
}

func (l *Logger) DatabaseError(err error) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	isCanShowCaller := ok && details != nil

	if isCanShowCaller {
		l.logger.WithField("caller", details.Name()).WithField("layer", "repository").WithError(err).Error("Databse error")
	} else {
		l.Error("Databse error", err)
	}
}

func (l *Logger) UseCaseError(err error) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	isCanShowCaller := ok && details != nil

	if isCanShowCaller {
		l.logger.WithField("caller", details.Name()).WithField("layer", "useCase").WithError(err).Error("Error in useCase")
	} else {
		l.Error("Error in useCase", err)
	}
}

func (l *Logger) ServerRequestInfo(c *gin.Context) error {
	method := c.Request.Method
	loggerEntry := l.logger.
		WithField("method", method).
		WithField("url", c.Request.URL.Path).
		WithField("client-ip", c.ClientIP()).
		WithField("content-type", c.ContentType()).
		WithField("user-agent", c.Request.UserAgent())

	if method == "GET" {
		query := c.Request.URL.Query()

		loggerEntry = loggerEntry.WithField("query", query)
	} else {
		body, err := c.GetRawData()
		data := string(body)

		if err != nil {
			l.logger.Error(err)

			return err
		}

		parsedData, err := getFieldsFromJson(data)

		if err != nil {
			l.logger.Error(err)

			return err
		}

		loggerEntry = loggerEntry.WithField("data", parsedData)

		c.Request.Body = io.NopCloser(bytes.NewReader(body))
	}

	loggerEntry.Info("REQUEST")

	return nil
}

func getFieldsFromJson(jsonMessage string) (logrus.Fields, error) {
	fields := logrus.Fields{}

	err := json.Unmarshal([]byte(jsonMessage), &fields)

	if err != nil {
		return fields, err
	}

	return fields, nil
}
