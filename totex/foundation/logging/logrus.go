package logging

import (
	"github.com/sirupsen/logrus"
)

// LogrusLogger implements the Logger interface using logrus
type LogrusLogger struct {
	logger *logrus.Logger
	name   string
}
