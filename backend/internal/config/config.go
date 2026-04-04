package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string
}

type DBConfig struct {
	Path string
}

type JWTConfig struct {
	Secret      string
	ExpireHours int
}

type Config struct {
	App AppConfig
	DB  DBConfig
	JWT JWTConfig
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, reading from environment")
	}

	expireHours, err := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOURS"))
	if err != nil {
		expireHours = 72
	}

	return &Config{
		App: AppConfig{
			Port: os.Getenv("APP_PORT"),
		},
		DB: DBConfig{
			Path: os.Getenv("DB_PATH"),
		},
		JWT: JWTConfig{
			Secret:      os.Getenv("JWT_SECRET"),
			ExpireHours: expireHours,
		},
	}

}
