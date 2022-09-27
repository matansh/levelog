package log

import (
	"fmt"
	"log"
	"os"

	"github.com/matansh/levelog"
	"github.com/matansh/levelog/loglevel"
)

var defaultLogger = levelog.NewLogger(
	loglevel.Debug,
	log.New(os.Stdout, "", 0),
	log.New(os.Stderr, "", 0),
)

// allowing alteration of the global logger
func SetGlobalLogger(logger levelog.Logger) {
	defaultLogger = logger
}

// "printer" logging interface, identical to the built in lib's exposed interface
// print statements will use log level info, fatal & panic log level error

func Printf(format string, v ...any) {
	defaultLogger.Log(loglevel.Info, fmt.Sprintf(format, v...))
}

func Print(v ...any) {
	defaultLogger.Log(loglevel.Info, fmt.Sprint(v...))
}

func Println(v ...any) {
	defaultLogger.Log(loglevel.Info, fmt.Sprintln(v...))
}

func Fatal(v ...any) {
	defaultLogger.Log(loglevel.Error, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...any) {
	defaultLogger.Log(loglevel.Error, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Fatalln(v ...any) {
	defaultLogger.Log(loglevel.Error, fmt.Sprintln(v...))
	os.Exit(1)
}

func Panic(v ...any) {
	s := fmt.Sprint(v...)
	defaultLogger.Log(loglevel.Error, s)
	panic(s)
}

func Panicf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	defaultLogger.Log(loglevel.Error, s)
	panic(s)
}

func Panicln(v ...any) {
	s := fmt.Sprintln(v...)
	defaultLogger.Log(loglevel.Error, s)
	panic(s)
}

// level based logging interface, as defined by levelog.Logger

func Debug(v ...any) {
	defaultLogger.Log(loglevel.Debug, fmt.Sprint(v...))
}

func Debugf(format string, v ...any) {
	defaultLogger.Log(loglevel.Debug, fmt.Sprintf(format, v...))
}

func Info(v ...any) {
	defaultLogger.Log(loglevel.Info, fmt.Sprint(v...))
}

func Infof(format string, v ...any) {
	defaultLogger.Log(loglevel.Info, fmt.Sprintf(format, v...))
}

func Warn(v ...any) {
	defaultLogger.Log(loglevel.Warn, fmt.Sprint(v...))
}

func Warnf(format string, v ...any) {
	defaultLogger.Log(loglevel.Warn, fmt.Sprintf(format, v...))
}

func Error(err error) {
	defaultLogger.Log(loglevel.Error, fmt.Sprintf("%+v", err))
}

// PanicRecovery is intended to be deferred by the programs entrypoint.
func PanicRecovery() {
	captured := recover()
	if err, isErr := captured.(error); isErr {
		defaultLogger.Error(err)
	}
	// re-panicking in order to maintain expected behavior
	panic(captured)
}
