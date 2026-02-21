// Package utils provides shared utilities for the GCP security analyzer.
package utils

import (
	"fmt"
	"io"
	"os"
	"time"
)

// LogLevel represents the logging severity.
type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

// Logger is a simple structured logger.
type Logger struct {
	level  LogLevel
	output io.Writer
}

// NewLogger creates a new Logger writing to the given writer at the given level.
func NewLogger(level LogLevel, output io.Writer) *Logger {
	if output == nil {
		output = os.Stderr
	}
	return &Logger{level: level, output: output}
}

// DefaultLogger is the package-level default logger (Info level, stderr).
var DefaultLogger = NewLogger(LevelInfo, os.Stderr)

func (l *Logger) log(level LogLevel, levelStr, format string, args ...interface{}) {
	if level < l.level {
		return
	}
	ts := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(l.output, "%s [%s] %s\n", ts, levelStr, msg)
}

// Debug logs at DEBUG level.
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(LevelDebug, "DEBUG", format, args...)
}

// Info logs at INFO level.
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(LevelInfo, "INFO ", format, args...)
}

// Warn logs at WARN level.
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(LevelWarn, "WARN ", format, args...)
}

// Error logs at ERROR level.
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(LevelError, "ERROR", format, args...)
}

// SetLevel updates the logger's minimum level.
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Package-level convenience functions using the DefaultLogger.

// Debug logs a debug message via the default logger.
func Debug(format string, args ...interface{}) { DefaultLogger.Debug(format, args...) }

// Info logs an info message via the default logger.
func Info(format string, args ...interface{}) { DefaultLogger.Info(format, args...) }

// Warn logs a warning via the default logger.
func Warn(format string, args ...interface{}) { DefaultLogger.Warn(format, args...) }

// Error logs an error via the default logger.
func Error(format string, args ...interface{}) { DefaultLogger.Error(format, args...) }

// SetVerbose configures the default logger for verbose (debug) output.
func SetVerbose() { DefaultLogger.SetLevel(LevelDebug) }

// SetQuiet suppresses everything below ERROR.
func SetQuiet() { DefaultLogger.SetLevel(LevelError) }
