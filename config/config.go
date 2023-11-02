package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config - configuration structure for api application.
type Config struct {
	App     App     `yaml:"app"`
	HTTP    HTTP    `yaml:"http"`
	Storage Storage `yaml:"storage"`
	Log     Log     `yaml:"log"`
}

// App - name and version of application.
type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

// HTTP - port for listen and serve application.
type HTTP struct {
	Port string `yaml:"port" env-default:"0.0.0.0:8080"`
}

// Storage - configuration for storage of application.
type Storage struct {
	URL string `yaml:"url" env:"URL"`
	// Pool int - TODO, when choosing PostgreSQL
}

// Log - level of logging aplication.
type Log struct {
	LogLevel string `yaml:"log_level" env-default:"info"`
}

// MustConfig - create configuration from config file.
func MustConfig() *Config {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		log.Fatalf("error read config: %s", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalf("error read env: %s", err)
	}

	return cfg
}
