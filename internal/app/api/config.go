package api

import (
	"github.com/admin-nsk/StandartWebServer/storage"
	"os"

	"github.com/joho/godotenv"
)

//General instance for API server of REST application
type Config struct {
	//Port
	BindAddr    string `toml:"bind_addr"`
	LoggerLevel string `toml:"logger_level"`
	Storage     *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}

func GetEnvConfig(fileName string, config *Config) error {
	config.BindAddr = os.Getenv("bind_addr")
	config.LoggerLevel = os.Getenv("logger_level")
	return godotenv.Load(fileName)
}
