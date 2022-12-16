//go:build !generate

// Generate by error_generator.go, never change this file manually.
package definition_error

import (
	"github.com/recative/recative-service-sdk/pkg/http_engine/http_err"
)

var (
	TimeFormatNotSupported = http_err.ResponseErrorType{
		Code: 401001,
		Name: "time_format_not_supported",
	}

	PermissionNotExist = http_err.ResponseErrorType{
		Code: 400001,
		Name: "permission_not_exist",
	}
)
