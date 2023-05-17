package usecase

import (
	"Diploma/pkg/models"
	"Diploma/pkg/user"
)

type UserUseCaseImpl struct {
	user.Gateway
}

type UserUseCaseModule struct {
	user.UseCase
}

func SetupUserUseCase(gateway user.Gateway) UserUseCaseModule {
	return UserUseCaseModule{
		UseCase: &UserUseCaseImpl{Gateway: gateway},
	}
}

func (u *UserUseCaseImpl) GetUserById(userId uint) (user *models.UserUsecase, err error) {
	return u.Gateway.GetUserById(userId)
}

func (u *UserUseCaseImpl) GetUserByUsername(username string) (user *models.UserUsecase, err error) {
	userUC := &models.UserUsecase{Username: username}
	return u.Gateway.GetUserByEmailOrUsername(userUC)
}

func (u *UserUseCaseImpl) UpdateUser(UpdateUser *models.UserUpdateInput) (user *models.UserUsecase, err error) {
	return u.Gateway.UpdateUser(UpdateUser)
}
