package auth

import "Diploma/pkg/models"

type Gateway interface {
	LogIn(user *models.UserUsecase) (gotUser *models.UserUsecase, err error)
	SaveRefreshToken(user *models.UserUsecase, refreshToken string) (err error)
	CheckRefreshToken(user *models.UserUsecase, oldRefreshToken string) (err error)
}
