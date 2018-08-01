package context

import (
	"os"
)

// Config is a struct that stores the configuration of the app
type Config struct {
	AppName   string
	DBURL     string
	Port      string
	DebugMode bool
	LogFormat string
}

// LoadConfig loads the configuration from environment variable
func LoadConfig() *Config {
	return &Config{
		AppName:   "Chainz",
		DBURL:     os.Getenv("DATABASE_URL"),
		Port:      os.Getenv("PORT"),
		DebugMode: true,
		LogFormat: "%{color}%{time:2006/01/02 15:04:05 -07:00 MST} [%{level:.6s}] %{shortfile} : %{color:reset}%{message}",
	}
}
