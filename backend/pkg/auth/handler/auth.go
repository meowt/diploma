package handler

import (
	"log"
	"net/http"
	"strings"

	"Diploma/pkg/auth"
	"Diploma/pkg/errorPkg"
	"Diploma/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Methods' naming principle:
// If method starts in lowercase - method doesn't process errors (only creates them)
// If method starts in Uppercase - method and creates, and processes errors

type Handler struct {
	delegate     auth.Delegate
	errProcessor *errorPkg.ErrorProcessor
	errCreator   *errorPkg.ErrorCreator
}

func SetupAuthHandler(authDelegate auth.Delegate, processor *errorPkg.ErrorProcessor, creator *errorPkg.ErrorCreator) Handler {
	return Handler{
		delegate:     authDelegate,
		errProcessor: processor,
		errCreator:   creator,
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

func (h *Handler) SignUp(c *gin.Context) {
	var (
		input    = models.SignUpInput{}
		response = authResponse{}
	)

	err := c.BindJSON(&input)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	response.AccessToken, response.RefreshToken, err = h.delegate.SignUp(&input)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}
	h.setRefreshToken(response.RefreshToken, c)

	log.Println(input.Email, input.Username, "signed up")
	c.JSON(http.StatusOK, response)
}

func (h *Handler) LogIn(c *gin.Context) {
	var (
		response authResponse
		input    = models.LogInInput{}
		err      error
	)

	if err = c.BindJSON(&input); err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	response.AccessToken, response.RefreshToken, err = h.delegate.LogIn(&input)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	h.setRefreshToken(response.RefreshToken, c)

	log.Println(input.Email, "logged in")
	c.JSON(http.StatusOK, response)
}

func (h *Handler) Refresh(c *gin.Context) {
	var (
		response authResponse
		err      error
	)

	identity, err := h.parseIdentity(c)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	response.RefreshToken, err = h.getRefreshTokenFromCookie(c)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	accessToken, refreshToken, err := h.delegate.RefreshToken(identity.Username, response.RefreshToken)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	log.Println(identity.Username, "refreshed his/her tokens")
	c.JSON(http.StatusOK, authResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// LogOut sets empty refresh token to user
func (h *Handler) LogOut(c *gin.Context) {
	h.setRefreshToken("", c)
	c.Status(http.StatusOK)
}

// Check method checks are equal user's refresh token and refresh token stored in DB
func (h *Handler) Check(c *gin.Context) {
	username, err := h.parseIdentity(c)
	if err != nil {
		h.errProcessor.ProcessError(c, err)
		return
	}

	log.Println(username, "checked his/her authentication")
	c.JSON(http.StatusOK, username)
}

func (h *Handler) parseIdentity(c *gin.Context) (userClaims *models.UserIdentity, err error) {
	header := c.GetHeader("Authorization")
	if header == "" {
		err = h.errCreator.NewErrEmptyAuthHeader()
		return
	}

	_, bearerToken, ok := strings.Cut(header, " ")
	if !ok {
		err = h.errCreator.NewErrStandardLibrary()
		return
	}

	userClaims.Username, err = h.delegate.ParseToken(bearerToken)
	if err != nil {
		//TODO: implement custom err
		//log.Println("Parsing token error:", err.Error())
		return
	}
	return
}

// getRefreshTokenFromCookie parses client's cookie with name "refresh_token" and returns result and error (if there is no such cookie)
func (h *Handler) getRefreshTokenFromCookie(c *gin.Context) (refreshToken string, err error) {
	refreshToken, err = c.Cookie("refresh_token")
	if err != nil {
		err = h.errCreator.NewErrEmptyCookie()
		return
	}
	return
}

func (h *Handler) setRefreshToken(refreshToken string, c *gin.Context) {
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
