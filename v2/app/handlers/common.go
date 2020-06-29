package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// errToJSON returns errors in JSON
func errToJSON(errs ...error) gin.H {
	if len(errs) == 1 {
		return gin.H{
			"error": errs[0].Error(),
		}
	}

	m := make(gin.H)
	for i, err := range errs {
		m[fmt.Sprintf("error-%d", i)] = err.Error()
	}
	return m
}
