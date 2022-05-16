package logger

import (
	"io"
	"runtime"

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
