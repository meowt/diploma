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

func (au *AuthDelegateImpl) SignUp(input *models.SignUpInput) (accessToken, refreshToken string, err error) {
	user := models.UserHttp{Username: input.Username, Email: input.Email, Password: input.Password}
	hashManager := service.NewHashManager()
	user.Password, err = hashManager.HashPassword(user.Password)
	if err != nil {
		return
	}
	userUsecase := user.ToUsecase()
	return au.UseCase.SignUp(userUsecase)
}

func (au *AuthDelegateImpl) LogIn(input *models.LogInInput) (accessToken, refreshToken string, err error) {
	user := &models.UserHttp{Email: input.Email, Password: input.Password}
	userUsecase := user.ToUsecase()
	return au.UseCase.LogIn(userUsecase)
}

func (au *AuthDelegateImpl) RefreshToken(username, oldRefreshToken string) (accessToken, refreshToken string, err error) {
	return au.UseCase.RefreshToken(&models.UserUsecase{Username: username}, oldRefreshToken)
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
