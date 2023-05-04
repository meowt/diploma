package auth

import "Diploma/pkg/models"

type Gateway interface {
	SignUp(user *models.UserUsecase) (err error)
}
