package levelog

import (
	"fmt"
	"testing"

	"github.com/matansh/levelog/loglevel"
)

type testLogger struct {
	t *testing.T
}

// TestLogger is an implementation if the level logger interface that pipes logs into the test output.
func NewTestLogger(t *testing.T) Logger {
	t.Helper()

	return &testLogger{
		t: t,
	}
}

func (tl testLogger) Log(lvl loglevel.Level, msg string) {
	tl.t.Helper()
	tl.t.Logf("[%s] %s", lvl, msg)
}

func (tl testLogger) Debug(msgs ...any) {
	tl.t.Helper()
	tl.Log(loglevel.Debug, fmt.Sprint(msgs...))
}

func (tl testLogger) Debugf(format string, msgs ...any) {
	tl.t.Helper()
	tl.Log(loglevel.Debug, fmt.Sprintf(format, msgs...))
}

func (tl testLogger) Info(msgs ...any) {
	tl.t.Helper()
	tl.Log(loglevel.Info, fmt.Sprint(msgs...))
}

func (tl testLogger) Infof(format string, msgs ...any) {
	tl.t.Helper()
	tl.Log(loglevel.Info, fmt.Sprintf(format, msgs...))
}

func (tl testLogger) Warn(msgs ...any) {
	tl.t.Helper()
	tl.Log(loglevel.Warn, fmt.Sprint(msgs...))
}

func (tl testLogger) Warnf(format string, msgs ...any) {
	tl.t.Helper()
	tl.Log(loglevel.Warn, fmt.Sprintf(format, msgs...))
}

func (tl testLogger) Error(err error) {
	tl.t.Helper()
	tl.Log(loglevel.Error, fmt.Sprintf("%+v", err))
}

func (tl testLogger) PanicRecovery() {} // irrelevant in a test context
