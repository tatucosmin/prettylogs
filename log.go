package main

import (
	"fmt"
	"io"
	"os"
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
	writer          io.Writer
	level           LoggerLevel
	disablePrefixes bool
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

func NewLogger() *Logger {
	return &Logger{
		writer:          os.Stdout,
		level:           LogInfoLevel,
		disablePrefixes: false,
	}
}

func (logger *Logger) handleLogPrefixFormat(prefix, str string) (int, error) {
	if prefix == "" {
		return fmt.Fprintf(logger.writer, "%v\n", str)
	}
	return fmt.Fprintf(logger.writer, "%s %v\n", prefix, str)
}

func (logger *Logger) Log(str string) (int, error) {
	var prefix string
	if !logger.disablePrefixes {
		if pf, ok := LogLevelPrefixes[logger.level]; ok {
			prefix = pf
		}
	}

	return logger.handleLogPrefixFormat(prefix, str)

}

func (logger *Logger) LogWithLevel(level LoggerLevel, str string) (int, error) {
	if level < logger.level {
		return 0, ErrUnderLoggerLevel
	}

	var prefix string
	if !logger.disablePrefixes {
		if pf, ok := LogLevelPrefixes[level]; ok {
			prefix = pf
		}
	}

	return logger.handleLogPrefixFormat(prefix, str)

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
