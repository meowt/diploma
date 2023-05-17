package handler

import (
	"net/http"
	"strconv"

	"Diploma/pkg/auth"
	"Diploma/pkg/theme"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	delegate     theme.Delegate
	authDelegate auth.Delegate
}

func SetupThemeHandler(themeDelegate theme.Delegate, authDelegate auth.Delegate) Handler {
	return Handler{
		delegate:     themeDelegate,
		authDelegate: authDelegate,
	}
}

func (h *Handler) InitThemeRoutes(router *gin.Engine) {
	themeRouter := router.Group("/theme")
	{
		themeRouter.POST("/upload", h.UploadTheme)
		themeRouter.GET("/getThemeById/:themeId", h.GetThemeById)
		themeRouter.GET("/getByUsername/:username", h.GetByUsername)
		themeRouter.GET("/getByFollows", h.GetByFollows)
		themeRouter.GET("/getLastThemesByUsername/:username", h.GetLastThemesByUsername)
		themeRouter.GET("/getLastThemes", h.GetLastThemes)
	}
}

type Response struct {
	Message string
}

func (h *Handler) UploadTheme(c *gin.Context) {
	//TODO: implement endpoint
	data, err := c.FormFile("File")
	if err != nil {
		h.
		return
	}

	c.JSON(http.StatusOK, Response{Message: "something like theme should be here"})
}

func (h *Handler) GetThemeById(c *gin.Context) {
	//TODO: implement endpoint
	idStr := c.Param("themeId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, Response{Message: strconv.Itoa(id)})
}

func (h *Handler) GetByUsername(c *gin.Context) {
	//TODO: implement endpoint
	c.JSON(http.StatusOK, Response{Message: "something like theme should be here"})
}

func (h *Handler) GetByFollows(c *gin.Context) {
	//TODO: implement endpoint
	c.JSON(http.StatusOK, Response{Message: "something like theme should be here"})
}

func (h *Handler) GetLastThemesByUsername(c *gin.Context) {
	//TODO: implement endpoint
	c.JSON(http.StatusOK, Response{Message: "something like theme should be here"})
}

func (h *Handler) GetLastThemes(c *gin.Context) {
	//TODO: implement endpoint
	c.JSON(http.StatusOK, Response{Message: "something like theme should be here"})
}
