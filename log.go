package prettylogs

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type ConfigurableLogger struct {
}

type LoggerLevel int
type LoggerError string

func (e LoggerError) Error() string {
	return string(e)
}

const (
	LogDebugLevel LoggerLevel = iota
	LogInfoLevel
	LogWarningLevel
	LogErrorLevel
	LogFatalLevel
)

type Logger struct {
	writer            io.Writer
	level             LoggerLevel
	disablePrefixes   bool
	disableTimestamps bool
}

const (
	ErrUnderLoggerLevel = LoggerError("passed LoggerLevel is under the current Loggerlevel")
)

var LogLevelPrefixes = map[LoggerLevel]string{
	LogDebugLevel:   "[DEBUG]",
	LogInfoLevel:    "[INFO]",
	LogWarningLevel: "[WARN]",
	LogErrorLevel:   "[ERROR]",
	LogFatalLevel:   "[FATAL]",
}

func NewConfigurable(w io.Writer, level LoggerLevel, disablePrexies, disableTimestamps bool) *Logger {
	return &Logger{
		w,
		level,
		disablePrexies,
		disableTimestamps,
	}
}

func New() *Logger {
	return &Logger{
		writer:            os.Stdout,
		level:             LogInfoLevel,
		disablePrefixes:   false,
		disableTimestamps: false,
	}
}

func (logger *Logger) SetLoggerLevel(level LoggerLevel) {
	logger.level = level
}

func (logger *Logger) handleLogFormat(str string, level LoggerLevel) (int, error) {
	var components []string

	// !!! The order of these comparisons is important as the expected way would be time, level, msg
	// * Instead of appending maybe I should index into the array to make sure order is always respected

	if !logger.disableTimestamps {
		timestamp := time.Now().Format(time.DateTime)
		components = append(components, timestamp)
	}

	if !logger.disablePrefixes {
		if pf, ok := LogLevelPrefixes[level]; ok {
			components = append(components, pf)
		}
	}

	components = append(components, fmt.Sprintf("%v", str))

	format := strings.Join(components, " ")

	return fmt.Fprintf(logger.writer, "%s", format)
}

func (logger *Logger) Log(str string) (int, error) {
	return logger.handleLogFormat(str, logger.level)
}

func (logger *Logger) LogWithLevel(level LoggerLevel, str string) (int, error) {
	if level < logger.level {
		return 0, ErrUnderLoggerLevel
	}

	return logger.handleLogFormat(str, level)

}

func (logger *Logger) Debug(str string) {
	logger.LogWithLevel(LogDebugLevel, str)
}

func (logger *Logger) Info(str string) {
	logger.LogWithLevel(LogInfoLevel, str)
}

func (logger *Logger) Warn(str string) {
	logger.LogWithLevel(LogWarningLevel, str)
}

func (logger *Logger) Error(str string) {
	logger.LogWithLevel(LogErrorLevel, str)
}

func (logger *Logger) Fatal(str string) {
	logger.LogWithLevel(LogFatalLevel, str)
}
