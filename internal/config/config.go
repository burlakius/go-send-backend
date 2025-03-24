package config

import (
	"log"
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

func MustLoad(configFilePath *string) *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(*configFilePath, &cfg); err != nil {
		log.Fatalf("Cannot read configuration file: %s\n", err)
	}

	return &cfg
}
