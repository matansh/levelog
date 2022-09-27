package loglevel

import (
	"errors"
	"fmt"
	"strings"
)

type Level int

const (
	// provided that 0 is the zero value of an int opting not to have it as a possible value.
	Debug = Level(iota + 1)
	Info
	Warn
	Error
)

const (
	debugStr = "DEBUG"
	infoStr  = "INFO"
	warnStr  = "WARN"
	errorStr = "ERROR"
)

var ErrInvalidLogLevel = errors.New("invalid log level")

// FromString converts string representations of available log levels into their corresponding enum values.
func FromString(level string) (Level, error) {
	switch strings.ToUpper(level) {
	case debugStr:
		return Debug, nil
	case infoStr:
		return Info, nil
	case warnStr:
		return Warn, nil
	case errorStr:
		return Error, nil
	default:
		return 0, fmt.Errorf("could not parse '%s': %w", level, ErrInvalidLogLevel)
	}
}

func (ll Level) String() string {
	return map[Level]string{
		Debug: debugStr,
		Info:  infoStr,
		Warn:  warnStr,
		Error: errorStr,
	}[ll]
}
