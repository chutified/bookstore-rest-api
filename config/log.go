package config

import (
	"io"
)

// LogConfig is a logging config struct.
type LogConfig struct {
	Output io.Writer
	Dest   string `yaml:"destination"`
}
