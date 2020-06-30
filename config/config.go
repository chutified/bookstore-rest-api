package config

import (
	"net/http"
	"os"
	"time"
)

// Config holds all configuration values.
type Config struct {
	Log *LogConfig
	Srv *ServerConfig
	DB  *DBConfig
}

var cfg = &Config{
	Log: &LogConfig{
		Output: os.Stdout,
	},

	Srv: &ServerConfig{
		Addr:              ":8081",
		ReadTimeout:       3 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       10 * time.Second,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	},

	DB: &DBConfig{
		Host:     "localhost",
		Port:     5432,
		DBName:   "project_01",
		User:     "user_00",
		Password: "159258",
	},
}

// GetConfig returns app configuration.
func GetConfig() *Config {
	return cfg
}
