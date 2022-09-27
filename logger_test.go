package levelog_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/matansh/levelog"
	"github.com/matansh/levelog/loglevel"
)

func TestLogger(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		minLevel       loglevel.Level
		writeLevel     loglevel.Level
		expectedStdout string
		expectedStderr string
	}{
		{
			name:           "debug in stdout",
			minLevel:       loglevel.Debug,
			writeLevel:     loglevel.Debug,
			expectedStdout: "[DEBUG] test\n",
		},
		{
			name:           "info in stdout",
			minLevel:       loglevel.Info,
			writeLevel:     loglevel.Info,
			expectedStdout: "[INFO] test\n",
		},
		{
			name:           "warn in stdout",
			minLevel:       loglevel.Warn,
			writeLevel:     loglevel.Warn,
			expectedStdout: "[WARN] test\n",
		},
		{
			name:           "error in stderr",
			minLevel:       loglevel.Debug,
			writeLevel:     loglevel.Error,
			expectedStderr: "[ERROR] test\n",
		},
		{
			name:       "debug filtered by min info",
			minLevel:   loglevel.Info,
			writeLevel: loglevel.Debug,
		},
		{
			name:       "info filtered by min warn",
			minLevel:   loglevel.Warn,
			writeLevel: loglevel.Info,
		},
		{
			name:       "warn filtered by min error",
			minLevel:   loglevel.Error,
			writeLevel: loglevel.Warn,
		},
		{
			name:           "error outputted in min error",
			minLevel:       loglevel.Error,
			writeLevel:     loglevel.Error,
			expectedStderr: "[ERROR] test\n",
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			stdout := bytes.Buffer{}
			stderr := bytes.Buffer{}
			testLogger := levelog.NewLogger(tc.minLevel, log.New(&stdout, "", 0), log.New(&stderr, "", 0))
			testLogger.Log(tc.writeLevel, "test")

			if tc.expectedStdout != stdout.String() {
				t.Fail()
			}
			if tc.expectedStderr != stderr.String() {
				t.Fail()
			}
		})
	}
}
