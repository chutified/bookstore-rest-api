package handlers

import (
	"fmt"

	"github.com/chutified/bookstore-api-example/app/models"
)

// HandleErrs returns errors in AppErrors.
func HandleErrs(errs ...error) models.AppErrors {

	fmt.Println(errs)

	var result models.AppErrors
	for _, err := range errs {
		result.Errors = append(result.Errors, err.Error())
	}
	return result
}
