package logger

import (
	"io"
	"os"
	"time"

	"github.com/aeilang/backend/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func StartLogger() {
	zerolog.TimeFieldFormat = time.RFC3339Nano

	cfg := config.GetConfig()
	logLevel := convertLogLevel(cfg.LogLevel)

	var output io.Writer = zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	if cfg.Environment == "production" {
		fileLogger := &lumberjack.Logger{
			Filename:   cfg.LogFileName,
			MaxSize:    cfg.LogMaxSize,
			MaxBackups: cfg.LogMaxBackups,
			MaxAge:     cfg.LogMaxAge,
			Compress:   cfg.LogCompress,
		}

		output = zerolog.MultiLevelWriter(os.Stderr, fileLogger)
	}

	log.Logger = zerolog.New(output).
		Level(logLevel).With().Timestamp().Logger()

	log.Info().Msg("logger set up sucessful")
}

func convertLogLevel(level string) zerolog.Level {
	switch level {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		return zerolog.TraceLevel
	}
}
