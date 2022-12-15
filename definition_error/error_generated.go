//go:build !generate

// Generate by error_generator.go, never change this file manually.
package definition_error

import (
	"github.com/recative/recative-service-sdk/pkg/http_engine/http_err"
)

var (
	InternalServerError = http_err.ResponseErrorType{
		Code: 500000,
		Name: "internal_server_error",
	}
)
