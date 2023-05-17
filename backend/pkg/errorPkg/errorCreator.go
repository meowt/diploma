package errorPkg

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

type ErrorManager struct {
	*ErrorCreator
	*ErrorProcessor
}

func InitErrorManager(logger *log.Logger) *ErrorManager {
	return &ErrorManager{
		ErrorCreator:   SetupErrorCreator(logger),
		ErrorProcessor: SetupErrorProcessor(logger),
	}
}

type ErrorCreator struct {
	*log.Logger
}

func SetupErrorCreator(Logger *log.Logger) *ErrorCreator {
	return &ErrorCreator{Logger}
}

var (
	ErrAccessDenied      = errors.New("access denied")
	ErrWrongPassword     = errors.New("wrong password")
	ErrParsingToken      = errors.New("error while parsing token")
	ErrEmptyAuthHeader   = errors.New("empty authentication header")
	ErrStandardLibrary   = errors.New("internal server error related to standard golang library")
	ErrSQLNoRows         = errors.New("nothing was found in database")
	ErrEmptyCookie       = errors.New("client's refresh_token cookie is empty")
	ErrBusyEmail         = errors.New("this email is already used by another user")
	ErrInformation       = fmt.Errorf("filename: %v, line: %v", filename, line)
	_, filename, line, _ = runtime.Caller(1)
)

type ErrCreator interface {
	New(oldErr error) (newErr error)
	NewErrWrongPassword() (err error)
	NewErrParsingToken() (err error)
	NewErrEmptyAuthHeader() (err error)
	NewErrStandardLibrary() (err error)
	NewErrEmptyCookie() (err error)
	NewErrBusyEmail() (err error)
}

// New method is just a stub
func (e *ErrorCreator) New(oldErr error) (newErr error) {
	return oldErr
}

func (e *ErrorCreator) NewErrSQLNoRows() (err error) {
	return ErrSQLNoRows
}

func (e *ErrorCreator) NewErrWrongPassword() (err error) {
	return errors.Join(ErrWrongPassword, ErrInformation)
}

func (e *ErrorCreator) NewErrParsingToken() (err error) {
	return ErrParsingToken
}

func (e *ErrorCreator) NewErrEmptyAuthHeader() (err error) {
	return ErrEmptyAuthHeader
}

func (e *ErrorCreator) NewErrStandardLibrary() (err error) {
	return ErrStandardLibrary
}

func (e *ErrorCreator) NewErrEmptyCookie() (err error) {
	return ErrEmptyCookie
}
func (e *ErrorCreator) NewErrBusyEmail() (err error) {
	return ErrBusyEmail
}

func (e *ErrorCreator) NewErrAccessDenied() (err error) {
	return ErrAccessDenied
}
