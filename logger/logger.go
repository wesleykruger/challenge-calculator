package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelError LogLevel = "error"
)

func init() {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	Logger = log.Output(output)
}

func SetLogLevel(level LogLevel) {
	switch level {
	case LogLevelDebug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case LogLevelInfo:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case LogLevelError:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func UserMsg(message string) {
	_, err := os.Stdout.WriteString(message + "\n")
	if err != nil {
		Logger.Error().Msg(fmt.Sprintf("Error writing to stdout: %v", err))
	}
}

func Debug(message string) {
	Logger.Debug().Msg(message)
}

func Info(message string) {
	Logger.Info().Msg(message)
}

func Error(message string) {
	Logger.Error().Msg(message)
}
