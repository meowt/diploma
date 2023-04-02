package handler

import (
	"fmt"
	"net/http"

	"Diploma/pkg/service/httpService"
	"Diploma/pkg/theme"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Handler struct {
	delegate theme.Delegate
}

func SetupThemeHandler(themeDelegate theme.Delegate) Handler {
	return Handler{
		delegate: themeDelegate,
	}
}

func (h *Handler) InitThemeRoutes(router *gin.Engine) {
	themeRouter := router.Group("/theme")
	{
		themeRouter.GET("/", h.GetInfo)
		themeRouter.GET("/:themeId", h.GetThemeById)
		themeRouter.GET("/getUploadPath", h.GetUploadPath)
	}
}

type Response struct {
	Message string
}

func (h *Handler) GetInfo(c *gin.Context) {
	client := httpService.SetupClient()
	req, err := httpService.SetupRequestByUrl(viper.GetString("http.cloud-api"))
	if err != nil {
		//error handling
		return
	}
	response, err := client.Do(&req)
	if err != nil {
		//error handling
		return
	}
	responseJson, err := httpService.ReadResponse(response)
	if err != nil {
		//error handling
		return
	}
	fmt.Println(responseJson)
	c.JSON(http.StatusOK, Response{Message: responseJson})
}

func (h *Handler) GetThemeById(c *gin.Context) {
	id := c.Param("themeId")
	c.JSON(http.StatusOK, Response{Message: id})
}

func (h *Handler) GetUploadPath(c *gin.Context) {
	client := httpService.SetupClient()
	req, err := httpService.SetupRequestByUrl(viper.GetString("http.cloud-api-upload"))
	if err != nil {
		//error handling
		return
	}
	fmt.Println(req.Header)
	response, err := client.Do(&req)
	if err != nil {
		//error handling
		return
	}
	responseJson, err := httpService.ReadResponse(response)
	if err != nil {
		//error handling
		return
	}
	fmt.Println(responseJson)
	c.JSON(http.StatusOK, Response{Message: responseJson})
}
