package modules

import (
	"Diploma/pkg/auth"
	authDelegate "Diploma/pkg/auth/delegate"
	authGateway "Diploma/pkg/auth/gateway"
	authHttp "Diploma/pkg/auth/handler"
	authUsecase "Diploma/pkg/auth/usecase"

	"Diploma/pkg/media"
	mediaDelegate "Diploma/pkg/media/delegate"
	mediaGateway "Diploma/pkg/media/gateway"
	mediaHttp "Diploma/pkg/media/handler"
	mediaUsecase "Diploma/pkg/media/usecase"

	"Diploma/pkg/user"
	userDelegate "Diploma/pkg/user/delegate"
	userGateway "Diploma/pkg/user/gateway"
	userHttp "Diploma/pkg/user/handler"
	userUsecase "Diploma/pkg/user/usecase"

	"Diploma/pkg/theme"
	themeDelegate "Diploma/pkg/theme/delegate"
	themeGateway "Diploma/pkg/theme/gateway"
	themeHttp "Diploma/pkg/theme/handler"
	themeUsecase "Diploma/pkg/theme/usecase"

	"Diploma/pkg/errorPkg"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmoiron/sqlx"
)

// GatewayModule section
type GatewayModule struct {
	AuthGateway  auth.Gateway
	UserGateway  user.Gateway
	ThemeGateway theme.Gateway
	MediaGateway media.Gateway
}

func SetupGateway(DatabaseClient *sqlx.DB, StorageClient *s3.S3, errorCreator *errorPkg.ErrorCreator) GatewayModule {
	return GatewayModule{
		AuthGateway:  authGateway.SetupAuthGateway(DatabaseClient, errorCreator),
		UserGateway:  userGateway.SetupUserGateway(DatabaseClient, errorCreator),
		ThemeGateway: themeGateway.SetupThemeGateway(DatabaseClient, StorageClient),
		MediaGateway: mediaGateway.SetupMediaGateway(DatabaseClient),
	}
}

// UseCaseModule section
type UseCaseModule struct {
	AuthUseCase  auth.UseCase
	UserUseCase  user.UseCase
	ThemeUseCase theme.UseCase
	MediaUseCase media.UseCase
}

func SetupUseCase(gatewayModule GatewayModule, creator *errorPkg.ErrorCreator) UseCaseModule {
	return UseCaseModule{
		AuthUseCase:  authUsecase.SetupAuthUseCase(gatewayModule.AuthGateway, gatewayModule.UserGateway, creator),
		UserUseCase:  userUsecase.SetupUserUseCase(gatewayModule.UserGateway),
		ThemeUseCase: themeUsecase.SetupThemeUseCase(gatewayModule.ThemeGateway),
		MediaUseCase: mediaUsecase.SetupMediaUseCase(gatewayModule.MediaGateway),
	}
}

// DelegateModule section
type DelegateModule struct {
	AuthDelegate  auth.Delegate
	UserDelegate  user.Delegate
	ThemeDelegate theme.Delegate
	MediaDelegate media.Delegate
}

func SetupDelegate(usecaseModule UseCaseModule) DelegateModule {
	return DelegateModule{
		AuthDelegate:  authDelegate.SetupAuthDelegate(usecaseModule.AuthUseCase),
		UserDelegate:  userDelegate.SetupUserDelegate(usecaseModule.UserUseCase),
		ThemeDelegate: themeDelegate.SetupThemeDelegate(usecaseModule.ThemeUseCase),
		MediaDelegate: mediaDelegate.SetupMediaDelegate(usecaseModule.MediaUseCase),
	}
}

// HandlerModule section
type HandlerModule struct {
	AuthHandler  authHttp.Handler
	UserHandler  userHttp.Handler
	ThemeHandler themeHttp.Handler
	MediaHandler mediaHttp.Handler
}

func SetupHandler(delegate DelegateModule, errorManager *errorPkg.ErrorManager) HandlerModule {
	return HandlerModule{
		AuthHandler:  authHttp.SetupAuthHandler(delegate.AuthDelegate, errorManager.ErrorProcessor, errorManager.ErrorCreator),
		UserHandler:  userHttp.SetupUserHandler(delegate.UserDelegate, delegate.AuthDelegate, errorManager.ErrorProcessor, errorManager.ErrorCreator),
		ThemeHandler: themeHttp.SetupThemeHandler(delegate.ThemeDelegate, delegate.AuthDelegate),
		MediaHandler: mediaHttp.SetupMediaHandler(delegate.MediaDelegate, delegate.AuthDelegate, errorManager.ErrorProcessor, errorManager.ErrorCreator),
	}
}
