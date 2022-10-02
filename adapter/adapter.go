package adapter

import (
	"log"
	"strings"

	"github.com/go-logr/logr"
)

// ToStd return a standard logger forwarding the messages to the given logr
func ToStd(l logr.Logger) *log.Logger {
	return log.New(&logrWriter{l: l}, "", 0)
}

type logrWriter struct {
	l logr.Logger
}

func (w *logrWriter) Write(p []byte) (n int, err error) {
	w.l.Info(strings.TrimSuffix(string(p), "\n"))
	return 0, nil
}
