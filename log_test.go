package prettylogs

import (
	"bytes"
	"testing"
)

func TestLogger(t *testing.T) {

	disableTimestamps := true
	disablePrefixes := false

	tc := []struct {
		name           string
		in             string
		logLevel       LoggerLevel
		want           string
		wantedLogLevel LoggerLevel
		err            error
	}{
		{"a debug level test", "test", LogDebugLevel, "[DEBUG] test", LogDebugLevel, nil},
		{"a info level test", "test", LogInfoLevel, "[INFO] test", LogInfoLevel, nil},
		{"a warning level test", "test", LogWarningLevel, "[WARN] test", LogWarningLevel, nil},
		{"an error level test", "test", LogErrorLevel, "[ERROR] test", LogErrorLevel, nil},
		{"a fatal level test", "test", LogFatalLevel, "[FATAL] test", LogFatalLevel, nil},
		{"a warning level message which should not be written", "test", LogErrorLevel, "", LogWarningLevel, ErrUnderLoggerLevel},
		{"a timestamp test", "test", LogFatalLevel, "[FATAL] test", LogFatalLevel, nil},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			logger := NewConfigurable(&buf, test.logLevel, disablePrefixes, disableTimestamps)
			_, err := logger.LogWithLevel(test.wantedLogLevel, test.in)

			assertError(t, test.err, err)

			got := buf.String()
			assertStrings(t, got, test.want)
		})
	}

}

func assertError(t testing.TB, got, want error) {
	if got != want {
		t.Fatalf("expected error %v but got %v", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	if got != want {
		t.Fatalf("expected output %s but got %s", got, want)
	}
}
