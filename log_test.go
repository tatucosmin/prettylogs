package prettylogs

import (
	"bytes"
	"testing"
)

func TestLogger(t *testing.T) {

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
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			logger := NewConfigurableLogger(&buf, test.logLevel, false)
			_, err := logger.LogWithLevel(test.wantedLogLevel, test.in)

			handleError(t, test.err, err)

			got := buf.String()

			if got != test.want {
				t.Fatalf("expected output %s but got %s", buf.String(), test.in)
			}
		})
	}

}

func handleError(t testing.TB, got, want error) {
	if got != want {
		t.Fatalf("expected error %v but got %v", got, want)
	}
}
