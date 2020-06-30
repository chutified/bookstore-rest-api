package handlers

import (
	"tommychu/workdir/027_api-example-v2/app/models"
)

// handleErrs returns errors in AppErrors.
func handleErrs(errs ...error) models.AppErrors {

	var result models.AppErrors
	for _, err := range errs {
		result.Errors = append(result.Errors, err.Error())
	}
	return result
}
