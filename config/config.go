package config

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/pkg/errors"
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

	cfg, err := fromFile("/settings.yaml")
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func fromFile(yamlPath string) (*Config, error) {

	// read file
	configPath := path.Join(rootDir(), yamlPath)
	bs, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// get cfg
	var cfg Config
	err = yaml.UnmarshalStrict(bs, &cfg)
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
		return nil, errors.Wrap(err, "invalid read_timeout")
	}
	cfg.Srv.ReadHeaderTimeout, err = time.ParseDuration(cfg.Srv.ReadHeaderTimeoutX)
	if err != nil {
		return nil, errors.Wrap(err, "invalid read_header_timeout")
	}
	cfg.Srv.WriteTimeout, err = time.ParseDuration(cfg.Srv.WriteTimeoutX)
	if err != nil {
		return nil, errors.Wrap(err, "invalid write_timeout")
	}
	cfg.Srv.IdleTimeout, err = time.ParseDuration(cfg.Srv.IdleTimeoutX)
	if err != nil {
		return nil, errors.Wrap(err, "invalid idle_timeout")
	}

	return &cfg, nil
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
