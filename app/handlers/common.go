package handlers

import "fmt"

// unWrap unwraps all errors into a string
func unWrap(errs []error) string {
	var result string
	for _, err := range errs {
		result = fmt.Sprintf("%s\n%s", result, err.Error())
	}
	return result
}
