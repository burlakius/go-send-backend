package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env"`
	Database   `yaml:"database"`
	HTTPServer `yaml:"http_server"`
}

type Database struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Port        string        `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func MustLoad() *Config {
	configPath := flag.String("-c", "./config/dev.yaml", "PATH TO CONFIG FILE")
	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatalf("Cannot load configuration file: File does not exist\nPath: %s\n", *configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(*configPath, &cfg); err != nil {
		log.Fatalf("Cannot read configuration file: %s\n", err)
	}

	return &cfg
}
