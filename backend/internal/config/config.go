package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	JWT  JWTConfig
	DB   DBConfig
}

type JWTConfig struct {
	Secret string
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	SSLMode  string
	TimeZone string
}

var Cfg *Config

func Load() {
	godotenv.Load()

	Cfg = &Config{
		Port: ":" + getEnv("APP_PORT", "8080"),
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "secret"),
		},
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "db"),
			User:     getEnv("DB_USER", "user"),
			Password: getEnv("DB_PASSWORD", "password"),
			Name:     getEnv("DB_NAME", "moxer"),
			Port:     getEnv("DB_PORT", "5432"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
			TimeZone: getEnv("DB_TIMEZONE", "UTC"),
		},
	}
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		c.Host, c.User, c.Password, c.Name, c.Port, c.SSLMode, c.TimeZone,
	)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
