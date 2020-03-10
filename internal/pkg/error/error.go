package error

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"pinterest/internal/pkg/response"
)

const(
	NoType = ErrorType(iota)
	BadRequest
	NotFound
	BadCookie
	BadToken
	UserNotFound
	BadLogin
	LoginIsExist
	BadEmail
	EmailIsExist
	BadPassword
	Unauthorized
	TooMuchSize
	BadPin
	DBError

	//add any type you want
)
type ErrorType uint

type Error struct {
	errorType ErrorType
	originalError error
}
// Error returns the mssage of a customError
func (error Error) Error() string {
	return error.originalError.Error()
}
// New creates a new customError
func (e ErrorType) New(msg string) error {
	return Error{errorType: e,
		originalError: New(msg),
	}
}
// New creates a new customError with formatted message
func (e ErrorType) Newf(msg string, args ...interface{}) error {
	err := fmt.Errorf(msg, args...)
	return Error{errorType: e, originalError: err}
}

// Wrap creates a new wrapped error
func (e ErrorType) Wrap(err error, msg string) error {
	return e.Wrapf(err, msg)
}

// Wrap creates a new wrapped error with formatted message
func (e ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	newErr := e.Wrapf(err, msg, args)

	return Error{errorType: e, originalError: newErr}
}

func New(msg string) error {
	return Error{errorType: NoType, originalError: New(msg)}
}

// Newf creates a no type error with formatted message
func Newf(msg string, args ...interface{}) error {
	return Error{errorType: NoType, originalError: New(fmt.Sprintf(msg, args...))}
}

// Wrap wrans an error with a string
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// Wrapf wraps an error with format string
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(Error); ok {
		return Error{
			errorType: customErr.errorType,
			originalError: wrappedError,
		}
	}

	return Error{errorType: NoType, originalError: wrappedError}
}

func GetType(err error) ErrorType {
	if customErr, ok := err.(Error); ok {
		return customErr.errorType
	}

	return NoType
}

func ErrorHandler(w http.ResponseWriter, err error) {
	var status int
	var message string
	switch err.(Error).errorType {
	case BadRequest:
		status = http.StatusBadRequest
		message = err.Error()
	case NotFound:
		status = http.StatusNotFound
		message = err.Error()
	/*case UserNotFound:
		status = http.StatusUnauthorized
		message = err.Error()*/
	case BadLogin:
		status = http.StatusUnauthorized
		message = err.Error()
	case BadPassword:
		status = http.StatusUnauthorized
		message = err.Error()
	case BadEmail:
		status = http.StatusUnauthorized
		message = err.Error()
	case LoginIsExist:
		status = http.StatusUnauthorized
		message = err.Error()
	case EmailIsExist:
		status = http.StatusUnauthorized
		message = err.Error()
	case Unauthorized:
		status = http.StatusUnauthorized
		message = err.Error()
	default:
		status = http.StatusInternalServerError
		message = "Internal server error"
	}

	response.Respond(w, status, map[string]string {
		"error": message,
	})
}