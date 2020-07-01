package handlers

import (
	"errors"
	"testing"
	"tommychu/workdir/026_api-example-v2/app/models"

	"github.com/google/go-cmp/cmp"
)

func TestHandleErrs(t *testing.T) {

	// tests table
	tests := []struct {
		name     string
		input    []error
		expected models.AppErrors
	}{
		{
			name:     "nil",
			input:    nil,
			expected: models.AppErrors{},
		},
		{
			name:     "one",
			input:    []error{errors.New("test0")},
			expected: models.AppErrors{Errors: []string{"test0"}},
		},
		{
			name:     "two",
			input:    []error{errors.New("test0"), errors.New("test1")},
			expected: models.AppErrors{Errors: []string{"test0", "test1"}},
		},
		{
			name:     "five",
			input:    []error{errors.New("test0"), errors.New("test1"), errors.New("test2"), errors.New("test3"), errors.New("test4")},
			expected: models.AppErrors{Errors: []string{"test0", "test1", "test2", "test3", "test4"}},
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			got := HandleErrs(test.input...)
			if !(cmp.Equal(got, test.expected)) {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}
