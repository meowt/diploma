package delegate

import (
	"errors"

	"Diploma/pkg/auth"
	"Diploma/pkg/models"
	"Diploma/pkg/service"
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

func (au *AuthDelegateImpl) SignUp(email, username, password string) (accessToken, refreshToken string, err error) {
	user := models.UserHttp{Username: username, Email: email, Password: password}
	hashManager := service.NewBCrypter()
	user.Password, err = hashManager.HashPassword(user.Password)
	if err != nil {
		return
	}
	userUsecase := user.ToUsecase()
	return au.UseCase.SignUp(userUsecase)
}

func (au *AuthDelegateImpl) LogIn(email, password string) (accessToken, refreshToken string, err error) {
	user := models.UserHttp{Email: email, Password: password}
	if err != nil {
		return
	}
	userUsecase := user.ToUsecase()
	return au.UseCase.LogIn(userUsecase)
}

func (au *AuthDelegateImpl) RefreshToken(username, oldRefreshToken string) (accessToken, refreshToken string, err error) {
	return au.UseCase.RefreshToken(username, oldRefreshToken)
}

func (au *AuthDelegateImpl) ParseToken(token string) (username string, err error) {
	claims, err := au.UseCase.ParseToken(token)
	if err != nil {
		return
	}
	if claims.Username == "" {
		err = errors.New("Empty username wtf ")
		return
	}
	username = claims.Username
	return
}
