package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    DatabaseURL string
    SessionKey  string
    Port        string
}

func Load() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using environment variables")
    }

    return &Config{
        DatabaseURL: getEnv("DATABASE_URL", "postgres://finance_user:secure_password@localhost/finance_db?sslmode=disable"),
        SessionKey:  getEnv("SESSION_KEY", "super-secret-key-change-in-production"),
        Port:        getEnv("PORT", "8080"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
