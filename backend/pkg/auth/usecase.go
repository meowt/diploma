package auth

import "Diploma/pkg/models"

type UseCase interface {
	SignUp(user *models.UserUsecase) (accessToken, refreshToken string, err error)
	LogIn(user *models.UserUsecase) (accessToken, refreshToken string, err error)
	RefreshToken(user *models.UserUsecase, oldRefreshToken string) (accessToken, refreshToken string, err error)
	ParseToken(token string) (identity *models.UserIdentity, err error)
}
