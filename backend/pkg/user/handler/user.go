package handler

import (
	"net/http"

	"Diploma/pkg/user"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	delegate user.Delegate
}

type Response struct {
	Message string
}

func (h *Handler) InitUserRoutes(router *gin.Engine) {
	userRouter := router.Group("/user")
	{
		userRouter.GET("/", h.GetUser)
	}
}

func SetupUserHandler(userDelegate user.Delegate) Handler {
	return Handler{
		delegate: userDelegate,
	}
}

func (h *Handler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Message: "something like user should be here"})
}
