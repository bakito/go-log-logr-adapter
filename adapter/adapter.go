package adapter

import (
	"log"

	"github.com/go-logr/logr"
)

// ToStd return a standard logger forwarding the messages to the given logr
func ToStd(l logr.Logger) *log.Logger {
	return ToStdBuilder(l).Build()
}

// ToStdBuilder return a standard logger builder
func ToStdBuilder(l logr.Logger) LoggerBuilder {
	return &loggerBuilder{l: l}
}

// SetDefault sets the given logr as writer for the default logger
func SetDefault(l logr.Logger) {
	ToStdBuilder(l).ApplyToDefault()
}

type logrWriter struct {
	l  logr.Logger
	ml MessageLogger
}

func (w *logrWriter) Write(msg []byte) (n int, err error) {
	if w.ml == nil {
		w.ml = &defaultMessageLogger{}
	}
	w.ml.Log(w.l, msg)
	return 0, nil
}
