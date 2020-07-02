package config

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"gopkg.in/yaml.v2"
)

// DebugMode toggles debug/release status
//const DebugMode = false

// Config is a application config struct.
type Config struct {
	Log       *LogConfig    `yaml:"logging"`
	Srv       *ServerConfig `yaml:"server"`
	DB        *DBConfig     `yaml:"database"`
	DebugMode bool          `yaml:"debug_mode"`
}

// GetConfig returns application config struct.
func GetConfig() (*Config, error) {

	cfg, err := fromFile()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func fromFile() (*Config, error) {

	// read file
	configPath := path.Join(rootDir(), "/settings.yaml")
	bs, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// get cfg
	var cfg Config
	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return nil, err
	}

	// set logging
	if cfg.Log.Dest == "STD_OUT" {
		cfg.Log.Output = os.Stdout
	} else {
		// TODO log to file
		cfg.Log.Output = os.Stdout
	}

	// set server
	cfg.Srv.ReadTimeout, err = time.ParseDuration(cfg.Srv.ReadTimeoutX)
	if err != nil {
		return nil, err
	}
	cfg.Srv.ReadHeaderTimeout, err = time.ParseDuration(cfg.Srv.ReadHeaderTimeoutX)
	if err != nil {
		return nil, err
	}
	cfg.Srv.WriteTimeout, err = time.ParseDuration(cfg.Srv.WriteTimeoutX)
	if err != nil {
		return nil, err
	}
	cfg.Srv.IdleTimeout, err = time.ParseDuration(cfg.Srv.IdleTimeoutX)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
