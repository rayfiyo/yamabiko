// internal/config/config.go

package config

import (
	"fmt"
	"os"
	"strconv"
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
func Load() *Config {
	port := getEnvAsInt("PORT", 8080)

	return &Config{
		Port:         port,
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnvAsInt("DB_PORT", 5432),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", "password"),
		DBName:       getEnv("DB_NAME", "myapp"),
		GoogleAPIKey: getEnv("GOOGLE_API_KEY", ""),
	}
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

func getEnvAsInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultVal
	}
	return val
}
