package config

import "os"

type Config struct {
	DBUrl string
}

func Load() *Config {
	return &Config{
		DBUrl: getEnv("DATABASE_URL", "postgres://postgres:password@localhost:5432/app?sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}