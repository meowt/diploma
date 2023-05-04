package auth

import "Diploma/pkg/models"

type UseCase interface {
	SignUp(user *models.UserUsecase) (accessToken, refreshToken string, err error)
}
