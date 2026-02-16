package observability

import (
	"log/slog"
	"os"
)

func NewLogger(env string) *slog.Logger {
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	return slog.New(h).With("service", "token-service", "env", env)
}
