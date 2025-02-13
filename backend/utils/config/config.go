// internal/config/config.go

package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         int
	DBHost       string
	DBPort       int
	DBUser       string
	DBPassword   string
	DBName       string
	GoogleAPIKey string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load the .env: %v", err)
	}

	port, err := getEnvAsInt("PORT", 8080)
	if err != nil {
		return nil, err
	}

	dbPort, err := getEnvAsInt("DB_PORT", 5432)
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:         port,
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       dbPort,
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", "password"),
		DBName:       getEnv("DB_NAME", "myapp"),
		GoogleAPIKey: getEnv("GOOGLE_API_KEY", ""),
	}, nil
}

func (c *Config) ConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName,
	)
}

func getEnv(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func getEnvAsInt(key string, defaultVal int) (int, error) {
	valStr := os.Getenv(key)
	if valStr == "" {
		return defaultVal, nil
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return val, err
	}
	return val, nil
}
