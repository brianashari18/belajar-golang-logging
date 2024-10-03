package golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLevel(T *testing.T) {
	logger := logrus.New()

	logger.Trace("This is trace level")
	logger.Debug("This is debug level")
	logger.Info("This is info level")
	logger.Warn("This is warn level")
	logger.Error("This is error level")
}

func TestLoggingLevel(T *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is trace level")
	logger.Debug("This is debug level")
	logger.Info("This is info level")
	logger.Warn("This is warn level")
	logger.Error("This is error level")
}
