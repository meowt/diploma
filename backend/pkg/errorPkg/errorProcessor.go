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

func InitErrorProcessor(Logger *log.Logger) *ErrorProcessor {
	return &ErrorProcessor{Logger}
}

func (processor *ErrorProcessor) ProcessError(c *gin.Context, errToProcess error) {
	log.Println("Error processed:", errToProcess)
	switch errToProcess {
	//TODO: implement more error handlers
	case ErrWrongPassword:
		c.AbortWithStatusJSON(http.StatusUnauthorized, errToProcess)
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, errToProcess)
	}
}
