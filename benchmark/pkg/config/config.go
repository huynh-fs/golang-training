package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Database
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string

	// JWT
	JWTAccessSecret  string
	JWTRefreshSecret string

	// Server
	GoPort string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, loading config from OS environment variables")
	}

	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),

		JWTAccessSecret:  os.Getenv("JWT_ACCESS_SECRET"),
		JWTRefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),

		GoPort: os.Getenv("GO_PORT"),
	}

	if cfg.DBHost == "" || cfg.JWTAccessSecret == "" {
		return nil, fmt.Errorf("một hoặc nhiều biến môi trường quan trọng chưa được thiết lập")
	}
	
	if cfg.DBSSLMode == "" {
		cfg.DBSSLMode = "disable" 
	}
	
	if cfg.GoPort == "" {
		cfg.GoPort = "8080" 
	}

	return cfg, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.DBHost,
		c.DBUser,
		c.DBPassword,
		c.DBName,
		c.DBPort,
		c.DBSSLMode,
	)
}