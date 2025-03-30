package logger

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestSetLogLevel(t *testing.T) {
	tests := []struct {
		name     string
		level    LogLevel
		expected zerolog.Level
	}{
		{
			name:     "debug level",
			level:    LogLevelDebug,
			expected: zerolog.DebugLevel,
		},
		{
			name:     "info level",
			level:    LogLevelInfo,
			expected: zerolog.InfoLevel,
		},
		{
			name:     "error level",
			level:    LogLevelError,
			expected: zerolog.ErrorLevel,
		},
		{
			name:     "invalid level defaults to info",
			level:    "invalid",
			expected: zerolog.InfoLevel,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			SetLogLevel(test.level)
			assert.Equal(t, test.expected, zerolog.GlobalLevel())
		})
	}
}

func TestLogFunctions(t *testing.T) {
	// Test that each log function can be called without panicking
	Debug("test debug message")
	Info("test info message")
	Error("test error message")
	UserMsg("test user message")
}
