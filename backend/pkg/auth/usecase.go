package auth

import "Diploma/pkg/models"

type UseCase interface {
	SignUp(user *models.UserUsecase) (accessToken, refreshToken string, err error)
	LogIn(user *models.UserUsecase) (accessToken, refreshToken string, err error)
	RefreshToken(username, oldRefreshToken string) (accessToken, refreshToken string, err error)
	ParseToken(token string) (claims *models.UserClaims, err error)
}
