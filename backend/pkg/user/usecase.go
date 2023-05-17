package user

import "Diploma/pkg/models"

type UseCase interface {
	GetUserById(userId uint) (user *models.UserUsecase, err error)
	GetUserByUsername(username string) (user *models.UserUsecase, err error)
	UpdateUser(UpdateUser *models.UserUpdateInput) (user *models.UserUsecase, err error)
}
