package error

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"pinterest/internal/pkg/response"
)

const(
	NoType = ErrorType(iota)
	BadRequest
	PinNotFound
	UserNotFound
	BadLogin
	BadPassword
	BadEmail
	LoginIsExist
	EmailIsExist
	Unauthorized
	TooMuchSize
)
type ErrorType uint

type Error struct {
	errorType ErrorType
	originalError error
	message string
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

func (e ErrorType) Newf(msg string, args ...interface{}) error {
	err := fmt.Errorf(msg, args...)
	return Error{errorType: e, originalError: err}
}

func (e ErrorType) Wrap(err error, msg string) error {
	return e.Wrapf(err, msg)
}

func (e ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	newErr := errors.Wrapf(err, msg, args)

	return Error{errorType: e, originalError: newErr}
}

func New(msg string) error {
	return Error{errorType: NoType, originalError: errors.New(msg)}
}

func Newf(msg string, args ...interface{}) error {
	return Error{errorType: NoType, originalError: New(fmt.Sprintf(msg, args...))}
}

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
			message: customErr.message,
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

func WithMessage(err error, message string) error {
	if customErr, ok := err.(Error); ok {
		customErr.message = message
		return customErr
	}

	return Error{errorType: NoType, originalError: err, message: message}
}

func ErrorHandler(w http.ResponseWriter, err error) {
	var status int
	var message string

	e, _ := err.(Error)
	switch GetType(e) {
	case BadRequest:
		status = http.StatusBadRequest
		message = e.message
	case UserNotFound:
		status = http.StatusNotFound
		message = "User is not found"
	case PinNotFound:
		status = http.StatusNotFound
		message = "Pin is not found"
	case BadLogin, BadPassword:
		status = http.StatusUnauthorized
		message = "Login or password is incorrect"
	case LoginIsExist:
		status = http.StatusUnauthorized
		message = "Change your login, login is already exist"
	case EmailIsExist:
		status = http.StatusUnauthorized
		message = "Change your email, email is already exist"
	case Unauthorized:
		status = http.StatusUnauthorized
		message = "User is unauthorized"
	case TooMuchSize:
		status = http.StatusUnauthorized
		message = "Image size should be less than 10 MB"
	default:
		status = http.StatusInternalServerError
		message = "Internal server error"
	}

	log.Println(err.Error())

	response.Respond(w, status, map[string]string {
		"error": message,
	})
}