package auth

import (
	"Diploma/pkg/models"

	"github.com/gin-gonic/gin"
)

type Delegate interface {
	SignUp(input *models.SignUpInput) (accessToken, refreshToken string, err error)
	LogIn(input *models.LogInInput) (accessToken, refreshToken string, err error)
	RefreshToken(username, oldRefreshToken string) (accessToken, refreshToken string, err error)
	ParseIdentity(c *gin.Context) (userIdentity *models.UserIdentity, err error)
}
