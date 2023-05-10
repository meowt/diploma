package modules

import (
	"Diploma/pkg/auth"
	authDelegate "Diploma/pkg/auth/delegate"
	authGateway "Diploma/pkg/auth/gateway"
	authHttp "Diploma/pkg/auth/handler"
	authUsecase "Diploma/pkg/auth/usecase"

	"Diploma/pkg/errorPkg"

	"Diploma/pkg/theme"
	themeHttp "Diploma/pkg/theme/handler"

	"Diploma/pkg/user"
	userDelegate "Diploma/pkg/user/delegate"
	userGateway "Diploma/pkg/user/gateway"
	userHttp "Diploma/pkg/user/handler"
	userUsecase "Diploma/pkg/user/usecase"

	themeDelegate "Diploma/pkg/theme/delegate"
	themeGateway "Diploma/pkg/theme/gateway"
	themeUsecase "Diploma/pkg/theme/usecase"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmoiron/sqlx"
)

// ErrorModule section
type ErrorModule struct {
	ErrCreator   errorPkg.ErrCreator
	ErrProcessor errorPkg.ErrProcessor
}

// GatewayModule section
type GatewayModule struct {
	AuthGateway  auth.Gateway
	UserGateway  user.Gateway
	ThemeGateway theme.Gateway
}

func SetupGateway(DatabaseClient *sqlx.DB, StorageClient *s3.S3, errCreator errorPkg.ErrCreator) GatewayModule {
	return GatewayModule{
		AuthGateway:  authGateway.SetupAuthGateway(DatabaseClient, errCreator),
		UserGateway:  userGateway.SetupUserGateway(DatabaseClient),
		ThemeGateway: themeGateway.SetupThemeGateway(DatabaseClient, StorageClient),
	}
}

// UseCaseModule section
type UseCaseModule struct {
	AuthUseCase  auth.UseCase
	UserUseCase  user.UseCase
	ThemeUseCase theme.UseCase
}

func SetupUseCase(gatewayModule GatewayModule, creator *errorPkg.ErrorCreator) UseCaseModule {
	return UseCaseModule{
		AuthUseCase:  authUsecase.SetupAuthUseCase(gatewayModule.AuthGateway, gatewayModule.UserGateway, creator),
		UserUseCase:  userUsecase.SetupUserUseCase(gatewayModule.UserGateway),
		ThemeUseCase: themeUsecase.SetupThemeUseCase(gatewayModule.ThemeGateway),
	}
}

// DelegateModule section
type DelegateModule struct {
	AuthDelegate  auth.Delegate
	UserDelegate  user.Delegate
	ThemeDelegate theme.Delegate
}

func SetupDelegate(usecaseModule UseCaseModule) DelegateModule {
	return DelegateModule{
		AuthDelegate:  authDelegate.SetupAuthDelegate(usecaseModule.AuthUseCase),
		UserDelegate:  userDelegate.SetupUserDelegate(usecaseModule.UserUseCase),
		ThemeDelegate: themeDelegate.SetupThemeDelegate(usecaseModule.ThemeUseCase),
	}
}

// HandlerModule section
type HandlerModule struct {
	AuthHandler  authHttp.Handler
	UserHandler  userHttp.Handler
	ThemeHandler themeHttp.Handler
}

func SetupHandler(delegate DelegateModule, errorProcessor *errorPkg.ErrorProcessor) HandlerModule {
	return HandlerModule{
		AuthHandler:  authHttp.SetupAuthHandler(delegate.AuthDelegate, errorProcessor),
		UserHandler:  userHttp.SetupUserHandler(delegate.UserDelegate),
		ThemeHandler: themeHttp.SetupThemeHandler(delegate.ThemeDelegate),
	}
}
