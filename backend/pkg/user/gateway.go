package user

import "Diploma/pkg/models"

type Gateway interface {
	GetUserByEmailOrUsername(useCaseUser *models.UserUsecase) (User *models.UserUsecase, err error)
	CreateUser(user *models.UserUsecase) (err error)
	GetUserById(userId uint) (User *models.UserUsecase, err error)
	UpdateUser(UpdateUser *models.UserUpdateInput) (user *models.UserUsecase, err error)
}
