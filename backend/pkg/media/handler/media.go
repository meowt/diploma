package handler

import (
	"Diploma/pkg/media"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	delegate media.Delegate
}

func SetupMediaHandler(mediaDelegate media.Delegate) Handler {
	return Handler{
		delegate: mediaDelegate,
	}
}

func (h *Handler) InitMediaRoutes(router *gin.Engine) {
	mediaRouter := router.Group("/media")
	_ = mediaRouter
	{
		mediaRouter.POST("/like", h.SetLike)
		mediaRouter.DELETE("/like", h.DeleteLike)
		mediaRouter.POST("/follow", h.FollowUser)
		mediaRouter.DELETE("/follow", h.UnfollowUser)
	}
}

func (h *Handler) SetLike(c *gin.Context) {

}

func (h *Handler) DeleteLike(c *gin.Context) {

}

func (h *Handler) FollowUser(c *gin.Context) {

}

func (h *Handler) UnfollowUser(c *gin.Context) {

}
