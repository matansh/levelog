/*
this package wraps the built in `log` lib,
exposing a log level based interface instead of the standard "printer" interface
*/

package levelog

import (
	"fmt"
	"log"

	"github.com/matansh/levelog/loglevel"
)

type Logger interface {
	Log(loglevel.Level, string)
	Debug(...any)
	Debugf(string, ...any)
	Info(...any)
	Infof(string, ...any)
	Warn(...any)
	Warnf(string, ...any)
	Error(error)
	PanicRecovery()
}

type logger struct {
	minLevel loglevel.Level
	stdout   *log.Logger
	stderr   *log.Logger
}

func NewLogger(level loglevel.Level, stdout, stderr *log.Logger) Logger {
	return &logger{
		minLevel: level,
		stdout:   stdout,
		stderr:   stderr,
	}
}

func (l logger) Log(lvl loglevel.Level, msg string) {
	// filtering out output that is below the wanted log level
	if lvl < l.minLevel {
		return
	}

	msg = fmt.Sprintf("[%s] %s", lvl, msg)
	if lvl == loglevel.Error {
		l.stderr.Print(msg)

		return
	}

	l.stdout.Print(msg)
}

func (l logger) Debug(v ...any) {
	l.Log(loglevel.Debug, fmt.Sprint(v...))
}

func (l logger) Debugf(format string, v ...any) {
	l.Log(loglevel.Debug, fmt.Sprintf(format, v...))
}

func (l logger) Info(v ...any) {
	l.Log(loglevel.Info, fmt.Sprint(v...))
}

func (l logger) Infof(format string, v ...any) {
	l.Log(loglevel.Info, fmt.Sprintf(format, v...))
}

func (l logger) Warn(v ...any) {
	l.Log(loglevel.Warn, fmt.Sprint(v...))
}

func (l logger) Warnf(format string, v ...any) {
	l.Log(loglevel.Warn, fmt.Sprintf(format, v...))
}

func (l logger) Error(err error) {
	l.Log(loglevel.Error, fmt.Sprintf("%+v", err))
}

// PanicRecovery is intended to be deferred by the programs entrypoint.
func (l logger) PanicRecovery() {
	captured := recover()
	if err, isErr := captured.(error); isErr {
		l.Error(err)
	}
	// re-panicking in order to maintain expected behavior
	panic(captured)
}
