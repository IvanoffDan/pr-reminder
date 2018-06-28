package utils

import (
	"sync"

	"github.com/sirupsen/logrus"
)

// Logger is an instance of logrus.Logger
var Logger *logrus.Logger
var once sync.Once

// GetLogger returns a new instance of a logger
func GetLogger() *logrus.Logger {
	once.Do(func() {
		Logger = NewLogger()
	})
	return Logger
}

// NewLogger creates a new instance of Logger
func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		ForceColors: true,
	}

	return logger
}
