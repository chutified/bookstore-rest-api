package handlers

import (
	"errors"
	"testing"
	"tommychu/workdir/026_api-example-v2/app/models"

	"github.com/google/go-cmp/cmp"
)

func TestHandleErrs(t *testing.T) {

	tests := []struct {
		title    string
		input    []error
		expected models.AppErrors
	}{
		{
			"nil",
			nil,
			models.AppErrors{},
		},
		{
			"one",
			[]error{errors.New("test0")},
			models.AppErrors{Errors: []string{"test0"}},
		},
		{
			"two",
			[]error{errors.New("test0"), errors.New("test1")},
			models.AppErrors{Errors: []string{"test0", "test1"}},
		},
		{
			"five",
			[]error{errors.New("test0"), errors.New("test1"), errors.New("test2"), errors.New("test3"), errors.New("test4")},
			models.AppErrors{Errors: []string{"test0", "test1", "test2", "test3", "test4"}},
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {

			got := HandleErrs(test.input...)
			if !(cmp.Equal(got, test.expected)) {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}
