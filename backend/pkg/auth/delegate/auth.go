package delegate

import (
	"Diploma/pkg/auth"
	"Diploma/pkg/models"
)

type AuthDelegateImpl struct {
	auth.UseCase
}

type AuthDelegateModule struct {
	auth.Delegate
}

func SetupAuthDelegate(usecase auth.UseCase) AuthDelegateModule {
	return AuthDelegateModule{
		Delegate: &AuthDelegateImpl{UseCase: usecase},
	}
}

func (au *AuthDelegateImpl) SignUp(email, password string) (accessToken, refreshToken string, err error) {
	user := models.UserHttp{Email: email, Password: password}
	userUsecase := user.ToUsecase()
	return au.UseCase.SignUp(&userUsecase)
}
