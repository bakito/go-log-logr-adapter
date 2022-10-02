package adapter

import (
	"strings"

	"github.com/go-logr/logr"
)

type MessageLogger interface {
	// Log the message with the given logger
	Log(l logr.Logger, msg []byte)
}

var _ MessageLogger = &defaultMessageLogger{}

type defaultMessageLogger struct{}

func (d defaultMessageLogger) Log(l logr.Logger, msg []byte) {
	l.Info(strings.TrimSuffix(string(msg), "\n"))
}
