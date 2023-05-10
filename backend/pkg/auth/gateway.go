package auth

import "Diploma/pkg/models"

type Gateway interface {
	SignUp(user *models.UserUsecase) (err error)
	LogIn(user *models.UserUsecase) (passwordHash string, err error)
	SaveRefreshToken(username, refreshToken string) (err error)
	CheckRefreshToken(username, oldRefreshToken string) (err error)
}
