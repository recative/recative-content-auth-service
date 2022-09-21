package http_err

import (
	"fmt"
	"net/http"
)

type ResponseErrorType struct {
	Code int
	Name string
}

func (responseErrorType ResponseErrorType) Error() string {
	return fmt.Sprintf("%d: %s", responseErrorType.Code, responseErrorType.Name)
}

func (responseErrorType ResponseErrorType) Wrap(err error) error {
	if _, ok := err.(*ResponseError); ok {
		return err
	}
	return &ResponseError{
		Id:      "",
		Code:    responseErrorType.Code,
		Name:    responseErrorType.Name,
		Message: "",
		Payload: err,
	}
}

func (responseErrorType ResponseErrorType) WrapAndSetId(err error, id string) error {
	if _, ok := err.(*ResponseError); ok {
		return err
	}
	return &ResponseError{
		Id:      id,
		Code:    responseErrorType.Code,
		Name:    responseErrorType.Name,
		Message: "",
		Payload: err,
	}
}

func (responseErrorType ResponseErrorType) ForceWrap(err error) error {
	return &ResponseError{
		Id:      "",
		Code:    responseErrorType.Code,
		Name:    responseErrorType.Name,
		Message: "",
		Payload: err,
	}
}

func (responseErrorType ResponseErrorType) New(message ...string) error {
	var realMessage = ""
	if len(message) != 0 {
		realMessage = message[0]
	}

	return &ResponseError{
		Id:      "",
		Code:    responseErrorType.Code,
		Name:    responseErrorType.Name,
		Message: realMessage,
		Payload: nil,
	}
}

func (responseErrorType ResponseErrorType) NewWithPayload(message string, payload any) error {
	return &ResponseError{
		Id:      "",
		Code:    responseErrorType.Code,
		Name:    responseErrorType.Name,
		Message: message,
		Payload: payload,
	}
}

func (responseErrorType ResponseErrorType) Is(err error) bool {
	v, ok := err.(*ResponseErrorType)
	if !ok {
		return false
	}
	return v.Code == responseErrorType.Code
}

type ResponseError struct {
	Id      string `json:"id"`
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
	Payload any    `json:"payload"`
}

func (responseError *ResponseError) Error() string {
	return fmt.Sprintf("%s: %d: %s: %s: %v",
		responseError.Id,
		responseError.Code,
		responseError.Name,
		responseError.Message,
		responseError.Payload)
}

func (responseError *ResponseError) ResponseStatusCode() int {
	if responseError == nil {
		return http.StatusOK
	}
	return responseError.Code / 1000
}
