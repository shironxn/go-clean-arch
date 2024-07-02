package config

import (
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

func New() (*App, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	app := App{
		Server: ServerConfig{
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
		},
		Database: DatabaseConfig{
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
		},
	}

	return &app, nil
}
