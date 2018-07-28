package context

import (
	"os"
	"time"
)

// Config is a struct that stores the configuration of the app
type Config struct {
	AppName string

	DBURL string

	JWTSecret   string
	JWTExpireIn time.Duration

	DebugMode bool
	LogFormat string
}

// LoadConfig loads the configuration from environment variable
func LoadConfig() *Config {
	return &Config{
		AppName:     "Chainz",
		DBURL:       os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		JWTExpireIn: 40000,
		DebugMode:   true,
		LogFormat:   "%{color}%{time:2006/01/02 15:04:05 -07:00 MST} [%{level:.6s}] %{shortfile} : %{color:reset}%{message}",
	}
}
