package logger

import (
	"io"
	"log/slog"
	"time"

	"github.com/lmittmann/tint"
)

func New(level slog.Level, output io.Writer) *slog.Logger {
	handler := tint.NewTextHandler(output, &tint.Options{Level: level, TimeFormat: time.TimeOnly, NoColor: false})
	return slog.New(handler)
}

func SetupLogger(level slog.Level, output io.Writer) *slog.Logger {
	log := New(level, output)
	slog.SetDefault(log)
	return log
}
