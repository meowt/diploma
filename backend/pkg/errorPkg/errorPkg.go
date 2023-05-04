package errorPkg

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ErrorGatewayImpl struct {
}

type ErrorGatewayModule struct {
}

func SetupErrorGateway() ErrorGatewayModule {
	return ErrorGatewayModule{
		Gateway: &ErrorGatewayImpl{DatabaseClient: databaseClient, ErrCreator: errCreator},
	}
}

type HttpError struct {
	Err  error
	Code int
}

type ErrCreator interface {
	New(oldErr error) (newErr error)
}

func New(oldErr error) (newErr error) {

	return fmt.Errorf("")
}

type ErrProcessor interface {
	ProcessError(c *gin.Context, errToProcess error)
}

func (e *HttpError) ProcessError(c *gin.Context, errToProcess error) (err error) {
	return
}
