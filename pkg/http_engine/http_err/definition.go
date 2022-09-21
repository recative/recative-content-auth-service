package http_err

var (
	InternalServerError = ResponseErrorType{
		Code: 500000,
		Name: "Internal Server Error",
	}

	UnexpectedInternalServerError = ResponseErrorType{
		Code: 500001,
		Name: "Unexpected Internal Server Error",
	}

	InvalidArgument = ResponseErrorType{
		Code: 400000,
		Name: "Invalid Argument",
	}

	Unauthorized = ResponseErrorType{
		Code: 401000,
		Name: "Unauthorized",
	}

	Forbidden = ResponseErrorType{
		Code: 403000,
		Name: "Forbidden",
	}

	NotFound = ResponseErrorType{
		Code: 404000,
		Name: "Not Found",
	}

	NotAcceptable = ResponseErrorType{
		Code: 406000,
		Name: "Not Acceptable",
	}

	Conflict = ResponseErrorType{
		Code: 409000,
		Name: "Conflict",
	}

	ServiceUnavailable = ResponseErrorType{
		Code: 503000,
		Name: "Service Unavailable",
	}
)
