package app

import (
	"fmt"
	"testing"
	"time"

	"github.com/chutified/bookstore-api-example/config"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/go-playground/assert.v1"
)

type nilWriter struct{}

func (w nilWriter) Write([]byte) (int, error) {
	return 0, nil
}

func TestNew(t *testing.T) {
	a := New()
	assert.Equal(t, fmt.Sprintf("%T", a), "*app.App")
}

func TestInitialize(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	cfg := config.GetConfig()
	cfg.Log.Output = nilWriter{}
	a := New()

	// tests table
	tests := []struct {
		name   string
		action func(*config.Config)
		dbUp   bool
		noErr  bool
	}{
		{
			name:   "ok",
			action: func(cfg *config.Config) {},
			dbUp:   true,
			noErr:  true,
		},
		{
			name:   "invalid config",
			action: func(cfg *config.Config) { cfg.DB.Password = "###_infalid_passwd_###" },
			dbUp:   false,
			noErr:  false,
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.action(cfg)

			// set
			err := a.Initialize(cfg)
			if err == nil {
				defer a.Close()
			}

			// check
			if err == nil {
				assert.Equal(t, test.dbUp, (a.DB.DB().Ping() == nil))
			}
			assert.Equal(t, test.noErr, err == nil)
		})
	}
}

func TestRun(t *testing.T) {
	cfg := config.GetConfig()
	a := New()
	err1 := a.Initialize(cfg)
	assert.Equal(t, err1, nil)
	defer func() {
		errs := a.Close()
		assert.Equal(t, errs[0], nil)
	}()

	go func() {
		err2 := a.Run()
		assert.Equal(t, err2, fmt.Errorf("http: Server closed"))
	}()

	time.Sleep(500 * time.Millisecond)

	err3 := a.Srv.Close()
	assert.Equal(t, err3, nil)
}
