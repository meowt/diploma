package user

import "Diploma/pkg/models"

type Delegate interface {
	GetUserById(userId int) (user *models.UserHttp, err error)
	GetUserByUsername(username string) (user *models.UserHttp, err error)
	UpdateUser(UpdateUser *models.UserUpdateInput) (user *models.UserHttp, err error)
}
