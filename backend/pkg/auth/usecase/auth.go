package usecase

import (
	"log"

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
	accessToken, err = au.Manager.NewJWT(user)
	if err != nil {
		return
	}
	refreshToken, err = au.Manager.NewRefreshToken()
	if err != nil {
		return
	}
	err = au.userGateway.CreateUser(user)
	if err != nil {
		return
	}
	err = au.authGateway.SaveRefreshToken(user, refreshToken)
	return
}

func (au *AuthUseCaseImpl) LogIn(user *models.UserUsecase) (accessToken, refreshToken string, err error) {
	notHashedPassword := user.Password
	user, err = au.authGateway.LogIn(user)
	if err != nil {
		//
		return
	}
	hashManager := service.NewHashManager()
	if !hashManager.ComparePassword(notHashedPassword, user.Password) {
		log.Println(notHashedPassword, user.Password)
		return accessToken, refreshToken, au.ErrorCreator.NewErrWrongPassword()
	}
	accessToken, err = au.Manager.NewJWT(user)
	if err != nil {
		//Access token generating error
		return
	}
	refreshToken, err = au.Manager.NewRefreshToken()
	if err != nil {
		//Refresh token generating error
		return
	}
	err = au.authGateway.SaveRefreshToken(user, refreshToken)
	return
}

func (au *AuthUseCaseImpl) RefreshToken(user *models.UserUsecase, oldRefreshToken string) (accessToken, refreshToken string, err error) {
	user, err = au.userGateway.GetUserByEmailOrUsername(user)
	if err != nil {
		return
	}
	err = au.authGateway.CheckRefreshToken(user, oldRefreshToken)
	if err != nil {
		return
	}
	accessToken, err = au.Manager.NewJWT(user)
	if err != nil {
		//Access token generating error
		return
	}
	refreshToken, err = au.Manager.NewRefreshToken()
	if err != nil {
		//Refresh token generating error
		return
	}
	err = au.authGateway.SaveRefreshToken(user, refreshToken)
	return
}

func (au *AuthUseCaseImpl) ParseToken(token string) (claims *models.UserClaims, err error) {
	key := viper.GetString("auth.signing_key")
	data, err := jwt.ParseWithClaims(token, &models.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
	if err != nil {
		err = au.ErrorCreator.NewErrParsingToken()
		return
	}

	claims, ok := data.Claims.(*models.UserClaims)
	if !ok {
		err = au.ErrorCreator.NewErrParsingToken()
		return
	}
	return
}
