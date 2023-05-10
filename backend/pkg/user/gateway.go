package user

import "Diploma/pkg/models"

type Gateway interface {
	GetUserByEmailOrUsername(useCaseUser *models.UserUsecase) (User *models.UserUsecase, err error)
}
