package delegate

import (
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
