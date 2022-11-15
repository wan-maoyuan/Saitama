package errors

import "net/http"

func BadRequest(reason, message string) *Error {
	return New(http.StatusBadRequest, reason, message)
}

func Unauthorized(reason, message string) *Error {
	return New(http.StatusUnauthorized, reason, message)
}

func Forbidden(reason, message string) *Error {
	return New(http.StatusForbidden, reason, message)
}

func NotFound(reason, message string) *Error {
	return New(http.StatusNotFound, reason, message)
}

func Conflict(reason, message string) *Error {
	return New(http.StatusConflict, reason, message)
}

func InternalServer(reason, message string) *Error {
	return New(http.StatusInternalServerError, reason, message)
}

func ServiceUnavailable(reason, message string) *Error {
	return New(http.StatusServiceUnavailable, reason, message)
}

func GatewayTimeout(reason, message string) *Error {
	return New(http.StatusGatewayTimeout, reason, message)
}
