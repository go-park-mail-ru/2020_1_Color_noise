package error

import (
	"2020_1_Color_noise/internal/pkg/metric"
	"2020_1_Color_noise/internal/pkg/response"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

const (
	NoType = ErrorType(iota)
	BadRequest
	PinNotFound
	BoardNotFound
	UserNotFound
	CommentNotFound
	BadLogin
	BadPassword
	BadEmail
	FollowingIsAlreadyDone
	FollowingIsNotYetDone
	LoginIsExist
	EmailIsExist
	Unauthorized
	TooMuchSize
	SearchNotFound
	StupidUser
)

type ErrorType uint

var array = [...]ErrorType{NoType, BadRequest, PinNotFound, BoardNotFound, UserNotFound, CommentNotFound,
	BadLogin, BadPassword, BadEmail, FollowingIsAlreadyDone, FollowingIsNotYetDone, LoginIsExist,
	EmailIsExist, Unauthorized, TooMuchSize, SearchNotFound}

func Cast(d int) ErrorType {
	if d >= len(array) || d < 0 {
		return NoType
	}

	return array[d]
}

type Error struct {
	errorType     ErrorType
	originalError error
	message       string
}

// Error returns the mssage of a customError
func (error Error) Error() string {
	return error.originalError.Error()
}

// New creates a new customError
func (e ErrorType) New(msg string) error {
	return Error{errorType: e,
		originalError: fmt.Errorf(msg),
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
	newErr := errors.Wrapf(err, msg, args...)

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
			errorType:     customErr.errorType,
			originalError: wrappedError,
			message:       customErr.message,
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

func ErrorHandler(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger, reqId interface{}, err error) {
	var status int
	var message string

	metric.IncreaseError()

	switch GetType(err) {
	case BadRequest:
		status = http.StatusBadRequest
		message = "Bad request"
	case SearchNotFound:
		status = http.StatusNotFound
		message = "Not found"
	case UserNotFound:
		status = http.StatusNotFound
		message = "User is not found"
	case PinNotFound:
		status = http.StatusNotFound
		message = "Pin is not found"
	case CommentNotFound:
		status = http.StatusNotFound
		message = "Comment is not found"
	case BoardNotFound:
		status = http.StatusNotFound
		message = "Board is not found"
	case BadLogin, BadPassword:
		status = http.StatusUnauthorized
		message = "Login or password is incorrect"
	case LoginIsExist:
		status = http.StatusUnauthorized
		message = "login"
	case EmailIsExist:
		status = http.StatusUnauthorized
		message = "email"
	case Unauthorized:
		status = http.StatusUnauthorized
		message = "User is unauthorized"
	case TooMuchSize:
		status = http.StatusRequestEntityTooLarge
		message = "Image size should be less than 7 MB"
	case FollowingIsAlreadyDone:
		status = http.StatusBadRequest
		message = "Following is already done"
	case FollowingIsNotYetDone:
		status = http.StatusBadRequest
		message = "Following is not yet done"
	case StupidUser:
		status = http.StatusTeapot
		message = "Why are you even trying to do this?"
	default:
		status = http.StatusInternalServerError
		message = "Internal server error"
	}
	/*
		logger.Error(
			zap.String("reqId:", fmt.Sprintf("%v", reqId)),
			zap.String("error:", err.Error()),
		)
	*/
	response.Respond(w, status, map[string]string{
		"error": message,
	})
}
