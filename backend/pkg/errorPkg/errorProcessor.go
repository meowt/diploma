package errorPkg

import (
	"log"
	"net/http"

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

type ErrJson struct {
	ErrorString string `json:"Error"`
}

func (processor *ErrorProcessor) ProcessError(c *gin.Context, errToProcess error) {
	log.Printf("[error]: %v, \nfilename: %v, line: %v\n", errToProcess, filename, line)
	errJson := ErrJson{ErrorString: errToProcess.Error()}

	switch errToProcess {
	//http.StatusBadRequest (code 400) section
	case ErrEmptyCookie:
		c.AbortWithStatusJSON(http.StatusBadRequest, errJson)
	case ErrEmptyAuthHeader:
		c.AbortWithStatusJSON(http.StatusBadRequest, errJson)

	//http.StatusUnauthorized (code 401) section
	case ErrParsingToken:
		c.AbortWithStatusJSON(http.StatusUnauthorized, errJson)
	case ErrWrongPassword:
		c.AbortWithStatusJSON(http.StatusUnauthorized, errJson)

	//http.StatusNotFound (code 404) section
	case ErrSQLNoRows:
		c.AbortWithStatusJSON(http.StatusNotFound, errJson)

	//http.StatusInternalServerError (code 500) section
	case ErrStandardLibrary:
		c.AbortWithStatusJSON(http.StatusInternalServerError, errJson)
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, errJson)
	}
}
