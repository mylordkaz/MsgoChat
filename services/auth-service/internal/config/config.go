package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAdress		string
	DatabaseURL			string
	JWTSecret			string
	JWTExpiryHours		int
	Environment			string
}

func Load() (*Config, error){
	// load .env file
	godotenv.Load()

	config := &Config{
		ServerAdress: getEnv("SERVER_ADDRESS", ":8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", ""),
		JWTExpiryHours: getEnvAsInt("JWT_EXPIRY_HOURS", 24),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
	return config, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	strValue := getEnv(key, "")
	if value, err := strconv.Atoi(strValue); err == nil {
		return value
	}
	return fallback
}
