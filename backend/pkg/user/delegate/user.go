package delegate

import (
	"Diploma/pkg/models"
	"Diploma/pkg/user"
)

type UserDelegateImpl struct {
	user.UseCase
}

type UserDelegateModule struct {
	user.Delegate
}

func SetupUserDelegate(usecase user.UseCase) UserDelegateModule {
	return UserDelegateModule{
		Delegate: &UserDelegateImpl{UseCase: usecase},
	}
}

func (u *UserDelegateImpl) GetUserById(userId int) (user *models.UserHttp, err error) {
	userUC, err := u.UseCase.GetUserById(uint(userId))
	if err != nil {
		return
	}
	return userUC.ToHttp(), err
}

func (u *UserDelegateImpl) GetUserByUsername(username string) (user *models.UserHttp, err error) {
	userUC, err := u.UseCase.GetUserByUsername(username)
	if err != nil {
		return
	}
	return userUC.ToHttp(), err
}

func (u *UserDelegateImpl) UpdateUser(UpdateUser *models.UserUpdateInput) (user *models.UserHttp, err error) {
	useCaseUser, err := u.UseCase.UpdateUser(UpdateUser)
	if err != nil {
		return
	}
	return useCaseUser.ToHttp(), err
}
