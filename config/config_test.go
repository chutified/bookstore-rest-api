package config

import "testing"

func TestGetConfig(t *testing.T) {
	cfg, err := GetConfig()
	if err != nil {
		t.Errorf("invalid config.yaml file")
	}
	if cfg.Log == nil {
		t.Errorf("expected cfg.Log")
	}
	if cfg.Srv == nil {
		t.Errorf("expected cfg.Srv")
	}
	if cfg.DB == nil {
		t.Errorf("expected cfg.DB")
	}
}
