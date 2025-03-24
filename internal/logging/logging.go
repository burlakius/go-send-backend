package logging

import (
	"log"
	"log/slog"
	"os"
)

const (
	EnvDev    = "dev"
	EnvLocal  = "local"
	EnvRemote = "remote"
)

func SetupLogger(logFilePath, env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case EnvDev:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case EnvLocal, EnvRemote:
		file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
		if err != nil {
			log.Fatalf("Error opening file: %s\n", err)
		}
		defer file.Close()

		logger = slog.New(
			slog.NewTextHandler(file, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return logger
}
