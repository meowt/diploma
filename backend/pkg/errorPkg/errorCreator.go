package errorPkg

import (
	"errors"
	"log"
)

type ErrorManager struct {
	*ErrorCreator
	*ErrorProcessor
}

func InitErrorManager(logger *log.Logger) *ErrorManager {
	return &ErrorManager{
		ErrorCreator:   InitErrorCreator(logger),
		ErrorProcessor: InitErrorProcessor(logger),
	}
}

type ErrorCreator struct {
	*log.Logger
}

func InitErrorCreator(Logger *log.Logger) *ErrorCreator {
	return &ErrorCreator{Logger}
}

var (
	ErrWrongPassword = errors.New("wrong password")
)

type ErrCreator interface {
	New(oldErr error) (newErr error)
	NewErrWrongPassword() (err error)
}

// New method is just a stub
func (e *ErrorCreator) New(oldErr error) (newErr error) {
	return ErrWrongPassword
}

func (e *ErrorCreator) NewErrWrongPassword() (err error) {
	return ErrWrongPassword
}
