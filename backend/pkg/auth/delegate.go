package auth

import "Diploma/pkg/models"

type Delegate interface {
	SignUp(input *models.SignUpInput) (accessToken, refreshToken string, err error)
	LogIn(input *models.LogInInput) (accessToken, refreshToken string, err error)
	RefreshToken(username, oldRefreshToken string) (accessToken, refreshToken string, err error)
	ParseToken(token string) (username string, err error)
}
