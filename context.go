/*
providing the option to safely persist and retrieve the logger instance from our context.
while its best not to treat our context as a dumping ground for singletons,
the logger is so commonly used that passing it around will pollute our method signatures.
*/
package levelog

import (
	"context"
	"log"
	"os"

	"github.com/matansh/levelog/loglevel"
)

// https://staticcheck.io/docs/checks/#SA1029
type loggerCtxKey string

const contextKey loggerCtxKey = "logger"

func ContextWithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, contextKey, logger)
}

func FromContext(ctx context.Context) Logger {
	if logger, ok := ctx.Value(contextKey).(Logger); ok {
		return logger
	}
	// while not ideal, its better to avoid failing just because we failed to find a logger on the provided context
	// instead providing a reasonable default and printing a warn message
	logger := NewLogger(
		loglevel.Debug,
		log.New(os.Stdout, "", 0),
		log.New(os.Stderr, "", 0),
	)
	logger.Warn("the provided context did not contain a logger")

	return logger
}
