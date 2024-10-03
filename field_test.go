package golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.WithField("username", "brianashari").Info("Hello World")
	logger.WithField("username", "brianashari").
		WithField("name", "brian anashari").
		Info("Hello World")

}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.WithFields(logrus.Fields{
		"username": "brianashari",
		"name":     "brian anashari",
	}).Info("Hello World")
}
