package config

import "io"

// LogConfig controlls logs destination.
type LogConfig struct {
	Output io.Writer
}
