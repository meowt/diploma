package usecase

import (
	"Diploma/pkg/auth"
	"Diploma/pkg/models"
)

type AuthUseCaseImpl struct {
	auth.Gateway
}

type AuthUseCaseModule struct {
	auth.UseCase
}

func SetupAuthUseCase(gateway auth.Gateway) AuthUseCaseModule {
	return AuthUseCaseModule{
		UseCase: &AuthUseCaseImpl{Gateway: gateway},
	}
}

func (au *AuthUseCaseImpl) SignUp(user *models.UserUsecase) (accessToken, refreshToken string, err error) {
	return "", "", au.Gateway.SignUp(user)
}
