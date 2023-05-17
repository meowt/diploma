package handler

import (
	"net/http"
	"strconv"

	"Diploma/pkg/auth"
	"Diploma/pkg/errorPkg"
	"Diploma/pkg/models"
	"Diploma/pkg/user"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	user.Delegate
	authDelegate auth.Delegate
	errProcessor *errorPkg.ErrorProcessor
	errCreator   *errorPkg.ErrorCreator
}

func SetupUserHandler(
	userDelegate user.Delegate,
	authDelegate auth.Delegate,
	processor *errorPkg.ErrorProcessor,
	creator *errorPkg.ErrorCreator) Handler {
	return Handler{
		Delegate:     userDelegate,
		authDelegate: authDelegate,
		errProcessor: processor,
		errCreator:   creator,
	}
}

func (h *Handler) InitUserRoutes(router *gin.Engine) {
	userRouter := router.Group("/user")
	{
		userRouter.GET("/getById/:userId", h.GetUserById)
		userRouter.GET("/getByUsername/:username", h.GetUserByUsername)
		userRouter.PUT("/update", h.UpdateUser)
	}
}

type Response struct {
	Message string
}

func (h *Handler) GetUserById(c *gin.Context) {
	_, err := h.authDelegate.ParseIdentity(c)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	userIdStr := c.Param("userId")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	User, err := h.Delegate.GetUserById(userIdInt)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	c.JSON(http.StatusOK, User)
}

func (h *Handler) GetUserByUsername(c *gin.Context) {
	_, err := h.authDelegate.ParseIdentity(c)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	username := c.Param("username")
	User, err := h.Delegate.GetUserByUsername(username)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	c.JSON(http.StatusOK, User)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	identity, err := h.authDelegate.ParseIdentity(c)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	UserUpdate := &models.UserUpdateInput{}
	err = c.BindJSON(UserUpdate)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	if UserUpdate.UpdatingUserId == 0 {
		UserUpdate.UpdatingUserId = identity.UserId
	}

	User, err := h.Delegate.UpdateUser(UserUpdate)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	c.JSON(http.StatusOK, User)
}
