package config

import "testing"

func TestGetConfig(t *testing.T) {
	cfg := GetConfig()
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
