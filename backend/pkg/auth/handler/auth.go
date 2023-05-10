package handler

import (
	"log"
	"net/http"
	"strings"

	"Diploma/pkg/auth"
	"Diploma/pkg/errorPkg"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Handler struct {
	delegate     auth.Delegate
	errProcessor *errorPkg.ErrorProcessor
}

func SetupAuthHandler(authDelegate auth.Delegate, processor *errorPkg.ErrorProcessor) Handler {
	return Handler{
		delegate:     authDelegate,
		errProcessor: processor,
	}
}

func (h *Handler) InitAuthRoutes(router *gin.Engine) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/sign-up", h.SignUp)
		authRouter.POST("/log-in", h.LogIn)
		authRouter.POST("/log-out", h.LogOut)
		authRouter.GET("/refresh", h.Refresh)
		authRouter.GET("/check", h.Check)
	}
}

type authResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type signUpInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var (
		input    = signUpInput{}
		response = authResponse{}
	)

	err := c.BindJSON(&input)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	response.AccessToken, response.RefreshToken, err = h.delegate.SignUp(input.Email, input.Username, input.Password)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}
	h.SetRefreshToken(response.RefreshToken, c)

	log.Println(input.Email, input.Username, "signed up")
	c.JSON(http.StatusOK, response)
}

type logInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) LogIn(c *gin.Context) {
	var (
		response authResponse
		input    = logInInput{}
		err      error
	)

	if err = c.BindJSON(&input); err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	response.AccessToken, response.RefreshToken, err = h.delegate.LogIn(input.Email, input.Password)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	h.SetRefreshToken(response.RefreshToken, c)

	log.Println(input.Email, "logged in")
	c.JSON(http.StatusOK, response)
}

func (h *Handler) Refresh(c *gin.Context) {
	var (
		response authResponse
		err      error
	)

	username, err := h.ParseIdentity(c)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	response.RefreshToken, err = h.GetRefreshTokenFromCookie(c)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	accessToken, refreshToken, err := h.delegate.RefreshToken(username, response.RefreshToken)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	log.Println(username, "refreshed his/her tokens")
	c.JSON(http.StatusOK, authResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// LogOut sets empty refresh token to user
func (h *Handler) LogOut(c *gin.Context) {
	h.SetRefreshToken("", c)
	c.Status(http.StatusOK)
}

// Check method checks are equal user's refresh token and refresh token stored in DB
func (h *Handler) Check(c *gin.Context) {
	username, err := h.ParseIdentity(c)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	log.Println(username, "checked his/her authentication")
	c.JSON(http.StatusOK, username)
}

func (h *Handler) ParseIdentity(c *gin.Context) (username string, err error) {
	header := c.GetHeader("Authorization")
	if header == "" {
		//TODO: implement custom err
		return
	}

	_, bearerToken, ok := strings.Cut(header, " ")
	if !ok {
		//TODO: implement custom err
		return
	}

	username, err = h.delegate.ParseToken(bearerToken)
	if err != nil {
		//TODO: implement custom err
		//log.Println("Parsing token error:", err.Error())
		return
	}
	return
}

// GetRefreshTokenFromCookie parses client's cookie with name "refresh_token" and returns result and error (if there is no such cookie)
func (h *Handler) GetRefreshTokenFromCookie(c *gin.Context) (refreshToken string, err error) {
	refreshToken, err = c.Cookie("refresh_token")
	if err != nil {
		// possible http.ErrNoCookie
		h.errProcessor.ProcessError(c, err)
		return
	}
	return
}

func (h *Handler) SetRefreshToken(refreshToken string, c *gin.Context) {
	c.SetCookie(
		"refresh_token",
		refreshToken,
		viper.GetInt("auth.refresh_token_ttl"),
		"/auth/",
		viper.GetString("auth.domain"),
		false,
		true,
	)
}
