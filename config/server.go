package config

import (
	"time"
)

// ServerConfig is a server config struct.
type ServerConfig struct {
	Addr               string `yaml:"addres"`
	ReadTimeout        time.Duration
	ReadHeaderTimeout  time.Duration
	WriteTimeout       time.Duration
	IdleTimeout        time.Duration
	ReadTimeoutX       string `yaml:"read_timeout"`
	ReadHeaderTimeoutX string `yaml:"read_header_timeout"`
	WriteTimeoutX      string `yaml:"write_timeout"`
	IdleTimeoutX       string `yaml:"idle_timeout"`
	MaxHeaderBytes     int    `yaml:"max_header_bytes"`
}
