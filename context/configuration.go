package context

import (
	"os"
	"time"
)

// Config is a struct that stores the configuration of the app
type Config struct {
	AppName string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret   string
	JWTExpireIn time.Duration

	DebugMode bool
	LogFormat string
}

// LoadConfig loads the configuration from environment variable
func LoadConfig() *Config {
	return &Config{
		AppName:     "Chainz",
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		JWTExpireIn: 40000,
		DebugMode:   true,
		LogFormat:   "%{color}%{time:2006/01/02 15:04:05 -07:00 MST} [%{level:.6s}] %{shortfile} : %{color:reset}%{message}",
	}
}
