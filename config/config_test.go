package config

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetConfig(t *testing.T) {

	file := fmt.Sprintf("%s/settings.yaml", rootDir())
	fileTest := fmt.Sprintf("%s.save", file)

	// tests table
	tests := []struct {
		name         string
		actionBefore func()
		actionAfter  func()
		noErr        bool
	}{
		{
			name:         "ok",
			actionBefore: func() {},
			actionAfter:  func() {},
			noErr:        true,
		},
		{
			name: "no file",
			actionBefore: func() {
				os.Rename(file, fileTest)
			},
			actionAfter: func() {
				os.Rename(fileTest, file)
			},
			noErr: false,
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.actionBefore()
			_, err := GetConfig()
			test.actionAfter()

			assert.Equal(t, test.noErr, (err == nil))
		})
	}
}

func TestFromFile(t *testing.T) {

	// tests table
	tests := []struct {
		name       string
		yaml       string
		errMessage string
	}{
		{
			name:       "ok standart output logging",
			yaml:       "/config/tests/0_settings.yaml",
			errMessage: "",
		},
		{
			// TODO file logging test
			name:       "ok file logging",
			yaml:       "/config/tests/1_settings.yaml",
			errMessage: "",
		},
		{
			name:       "non existing file",
			yaml:       "/non-existing.file",
			errMessage: "no such file or directory",
		},
		{
			name:       "invalid file type",
			yaml:       "/config/tests/invalid.yaml",
			errMessage: "unmarshal errors",
		},
		{
			name:       "invalid read tmout",
			yaml:       "/config/tests/0_invalid.yaml",
			errMessage: "invalid read_timeout: time",
		},
		{
			name:       "invalid read header tmout",
			yaml:       "/config/tests/1_invalid.yaml",
			errMessage: "invalid read_header_timeout",
		},
		{
			name:       "invalid write tmout",
			yaml:       "/config/tests/2_invalid.yaml",
			errMessage: "invalid write_timeout",
		},
		{
			name:       "invalid idle tmout",
			yaml:       "/config/tests/3_invalid.yaml",
			errMessage: "invalid idle_timeout",
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			_, err := fromFile(test.yaml)
			if err != nil {
				expr := regexp.MustCompile(fmt.Sprintf(`(.*%v.*)`, test.errMessage))
				ok := expr.Match([]byte(err.Error()))

				assert.Equal(t, true, ok)
			} else {
				assert.Equal(t, test.errMessage, "")
			}
		})
	}
}
