package usecase

import (
	"Diploma/pkg/auth"
	"Diploma/pkg/errorPkg"
	"Diploma/pkg/models"
	"Diploma/pkg/service"
	"Diploma/pkg/user"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type AuthUseCaseImpl struct {
	authGateway auth.Gateway
	userGateway user.Gateway
	*service.Manager
	*errorPkg.ErrorCreator
}

type AuthUseCaseModule struct {
	auth.UseCase
}

func SetupAuthUseCase(authGateway auth.Gateway, userGateway user.Gateway, creator *errorPkg.ErrorCreator) AuthUseCaseModule {
	tokenManager := service.NewManager(viper.GetString("auth.signing_key"))
	return AuthUseCaseModule{
		UseCase: &AuthUseCaseImpl{
			authGateway:  authGateway,
			userGateway:  userGateway,
			Manager:      tokenManager,
			ErrorCreator: creator,
		},
	}
}

func (au *AuthUseCaseImpl) SignUp(user *models.UserUsecase) (accessToken, refreshToken string, err error) {
	accessToken, err = au.Manager.NewJWT(user.Username)
	if err != nil {
		return
	}
	refreshToken, err = au.Manager.NewRefreshToken()
	if err != nil {
		return
	}
	err = au.authGateway.SignUp(user)
	if err != nil {
		return
	}
	err = au.authGateway.SaveRefreshToken(user.Username, refreshToken)
	return
}

func (au *AuthUseCaseImpl) LogIn(user *models.UserUsecase) (accessToken, refreshToken string, err error) {
	passwordHash, err := au.authGateway.LogIn(user)
	if err != nil {
		//
		return
	}
	hashManager := service.NewBCrypter()
	if hashManager.ComparePassword(user.Password, passwordHash) {
		user, err = au.userGateway.GetUserByEmailOrUsername(user)
		if err != nil {
			return
		}
		accessToken, err = au.Manager.NewJWT(user.Username)
		if err != nil {
			//Access token generating error
			return
		}
		refreshToken, err = au.Manager.NewRefreshToken()
		if err != nil {
			//Refresh token generating error
			return
		}
		err = au.authGateway.SaveRefreshToken(user.Username, refreshToken)
		return
	}
	return accessToken, refreshToken, au.ErrorCreator.NewErrWrongPassword()
}

func (au *AuthUseCaseImpl) RefreshToken(username, oldRefreshToken string) (accessToken, refreshToken string, err error) {
	err = au.authGateway.CheckRefreshToken(username, oldRefreshToken)
	if err != nil {
		return
	}
	accessToken, err = au.Manager.NewJWT(username)
	if err != nil {
		//Access token generating error
		return
	}
	refreshToken, err = au.Manager.NewRefreshToken()
	if err != nil {
		//Refresh token generating error
		return
	}
	err = au.authGateway.SaveRefreshToken(username, refreshToken)
	return
}

func (au *AuthUseCaseImpl) ParseToken(token string) (claims *models.UserClaims, err error) {
	key := viper.GetString("auth.signing_key")
	data, err := jwt.ParseWithClaims(token, &models.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
	if err != nil {
		return
	}

	claims, ok := data.Claims.(*models.UserClaims)
	if !ok {
		//TODO: implement custom err
		return
	}
	return
}
