package handler

import (
	"Diploma/pkg/auth"
	"Diploma/pkg/errorPkg"
	"Diploma/pkg/media"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	delegate     media.Delegate
	authDelegate auth.Delegate
	errProcessor *errorPkg.ErrorProcessor
	errCreator   *errorPkg.ErrorCreator
}

func SetupMediaHandler(mediaDelegate media.Delegate,
	authDelegate auth.Delegate,
	processor *errorPkg.ErrorProcessor,
	creator *errorPkg.ErrorCreator) Handler {
	return Handler{
		delegate:     mediaDelegate,
		authDelegate: authDelegate,
		errProcessor: processor,
		errCreator:   creator,
	}
}

func (h *Handler) InitMediaRoutes(router *gin.Engine) {
	mediaRouter := router.Group("/media")
	{
		mediaRouter.POST("/like/:theme_id", h.SetLike)
		mediaRouter.DELETE("/like", h.DeleteLike)
		mediaRouter.POST("/follow", h.FollowUser)
		mediaRouter.DELETE("/follow", h.UnfollowUser)
		mediaRouter.PUT("/background", h.UpdateBackground)
		mediaRouter.PUT("/avatar", h.UpdateAvatar)
		mediaRouter.PUT("/description", h.UpdateDescription)
	}
}

func (h *Handler) SetLike(c *gin.Context) {
	identity, err := h.authDelegate.ParseIdentity(c)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	themeId := c.GetUint(":theme_id")
	err = h.delegate.SetLike(identity.UserId, themeId)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}
}

func (h *Handler) DeleteLike(c *gin.Context) {

}

func (h *Handler) FollowUser(c *gin.Context) {

}

func (h *Handler) UnfollowUser(c *gin.Context) {

}

func (h *Handler) UpdateBackground(c *gin.Context) {

}

func (h *Handler) UpdateAvatar(c *gin.Context) {

}

func (h *Handler) UpdateDescription(c *gin.Context) {

}
