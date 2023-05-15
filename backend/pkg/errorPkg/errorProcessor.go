package errorPkg

import (
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

type ErrorProcessor struct {
	*log.Logger
}

type ErrProcessor interface {
	ProcessError(c *gin.Context, errToProcess error)
}

func SetupErrorProcessor(Logger *log.Logger) *ErrorProcessor {
	return &ErrorProcessor{Logger}
}

func (processor *ErrorProcessor) ProcessError(c *gin.Context, errToProcess error) {
	_, filename, line, _ := runtime.Caller(1)
	log.Printf("[error]: %v, \nfilename: %v, line: %v\n", errToProcess, filename, line)

	switch errToProcess {
	//http.StatusBadRequest (code 400) section
	case ErrEmptyCookie:
		c.AbortWithStatusJSON(http.StatusBadRequest, errToProcess)
	case ErrEmptyAuthHeader:
		c.AbortWithStatusJSON(http.StatusBadRequest, errToProcess)

	//http.StatusUnauthorized (code 401) section
	case ErrParsingToken:
		c.AbortWithStatusJSON(http.StatusUnauthorized, errToProcess)
	case ErrWrongPassword:
		c.AbortWithStatusJSON(http.StatusUnauthorized, errToProcess)

	//http.StatusInternalServerError (code 500) section
	case ErrStandardLibrary:
		c.AbortWithStatusJSON(http.StatusInternalServerError, errToProcess)
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, errToProcess)
	}
}
