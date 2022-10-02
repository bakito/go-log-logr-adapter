package adapter

import (
	"log"

	"github.com/go-logr/logr"
)

type LoggerBuilder interface {
	Build() *log.Logger
	ApplyToDefault()
	WithMessageLogger(ml MessageLogger) LoggerBuilder
	WithPrefix(prefix string) LoggerBuilder
	WithFlags(flags int) LoggerBuilder
}

type loggerBuilder struct {
	l      logr.Logger
	ml     MessageLogger
	prefix string
	flags  int
}

func (b *loggerBuilder) Build() *log.Logger {
	return log.New(&logrWriter{l: b.l, ml: b.ml}, b.prefix, b.flags)
}

func (b *loggerBuilder) ApplyToDefault() {
	log.SetOutput(&logrWriter{l: b.l, ml: b.ml})
	log.SetPrefix(b.prefix)
	log.SetFlags(b.flags)
}

func (b *loggerBuilder) WithMessageLogger(ml MessageLogger) LoggerBuilder {
	b.ml = ml
	return b
}

func (b *loggerBuilder) WithPrefix(prefix string) LoggerBuilder {
	b.prefix = prefix
	return b
}

func (b *loggerBuilder) WithFlags(flags int) LoggerBuilder {
	b.flags = flags
	return b
}
