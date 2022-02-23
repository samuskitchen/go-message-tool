package errores

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	HttpCode     int
	Err          error
	ErrorMessage string
}

func newErr(err error, message string, httpcode int) error {
	return &CustomError{
		Err:          err,
		ErrorMessage: message,
		HttpCode:     httpcode,
	}
}

func (e *CustomError) GetError() error {
	return e.Err
}

func (e *CustomError) Message() string {
	return e.ErrorMessage
}

func (e *CustomError) Error() string {
	if e.Err == nil {
		return "error trivial"
	}

	return e.Err.Error()
}

func NewBadRequest(err error, format string, a ...interface{}) error {
	return newErr(err, fmt.Sprintf(format, a...), http.StatusBadRequest)
}

func NewInternal(err error, format string, a ...interface{}) error {
	return newErr(err, fmt.Sprintf(format, a...), http.StatusInternalServerError)
}

func NewUnsupported(err error, format string, a ...interface{}) error {
	return newErr(err, fmt.Sprintf(format, a...), http.StatusUnsupportedMediaType)
}

func NewUnauthorized(err error, format string, a ...interface{}) error {
	return newErr(err, fmt.Sprintf(format, a...), http.StatusUnauthorized)
}

func NewForbidden(err error, format string, a ...interface{}) error {
	return newErr(err, fmt.Sprintf(format, a...), http.StatusForbidden)
}

func NewNotFound(err error, format string, a ...interface{}) error {
	return newErr(err, fmt.Sprintf(format, a...), http.StatusNotFound)
}
