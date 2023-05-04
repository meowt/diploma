package handler

import (
	"net/http"

	"Diploma/pkg/auth"
	"Diploma/pkg/errorPkg"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	delegate     auth.Delegate
	errProcessor errorPkg.ErrProcessor
}

func SetupAuthHandler(authDelegate auth.Delegate, errProcessor errorPkg.ErrProcessor) Handler {
	return Handler{
		delegate:     authDelegate,
		errProcessor: errProcessor,
	}
}

func (h *Handler) InitAuthRoutes(router *gin.Engine) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/reg", h.SignUp)
		authRouter.POST("/log", h.LogIn)
	}
}

type signUpInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var err error
	signInInput := &signUpInput{}
	if err = c.BindJSON(signInInput); err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	response := authResponse{}
	response.AccessToken, response.RefreshToken, err = h.delegate.SignUp(signInInput.Email, signInInput.Password)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	setRefreshToken(refreshToken, c)

	refreshToken2, _ := c.Request.Cookie("token2")

	c.JSON(http.StatusOK, response)
}

func (h *Handler) LogIn(c *gin.Context) {
	signInInput := &signUpInput{}
	if err := c.BindJSON(signInInput); err != nil {
		h.errProcessor.ProcessError(err)
		return
	}

	accessToken, refreshToken, err := h.delegate.SignUp(signInInput.Email, signInInput.Password)
	if err != nil {
		h.errProcessor.ProcessError(err)
		return
	}

	setRefreshToken(refreshToken, c)

	refreshToken2, _ := c.Request.Cookie("token2")

	c.JSON(http.StatusOK, authResponse{
		AccessToken: accessToken,
	})
}
