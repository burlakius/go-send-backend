package main

import (
	"flag"
	"log"
	"log/slog"
	"os"

	"github.com/burlakius/go-send-backend/internal/config"
	"github.com/burlakius/go-send-backend/internal/logging"
)

var (
	serverConfig *config.Config
	serverLogger *slog.Logger
)

func init() {
	configFilePath, logFilePath := parseFlags()

	serverConfig = config.MustLoad(configFilePath)
	serverLogger = logging.SetupLogger(*logFilePath, serverConfig.Env)
}

func main() {
	// TODO: DB
	// TODO: Auth
	// TODO: API
	// TODO: NP connection
	// TODO: UP connection
}

func parseFlags() (*string, *string) {
	configFilePath := flag.String("config", "./config/dev.yaml", "PATH TO CONFIG FILE")
	if _, err := os.Stat(*configFilePath); os.IsNotExist(err) {
		log.Fatalf("Cannot load configuration file: File does not exist\nPath: %s\n", *configFilePath)
	}

	logFilePath := flag.String("logs", "./logs/logs.log", "PATH TO LOG FILE")
	if _, err := os.Stat(*logFilePath); os.IsNotExist(err) {
		log.Fatalf("Cannot open logs file: File does not exist\nPath: %s\n", *logFilePath)
	}

	flag.Parse()

	return configFilePath, logFilePath
}
