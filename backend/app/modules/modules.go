package modules

import (
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

// GatewayModule section
type GatewayModule struct {
	UserGateway  user.Gateway
	ThemeGateway theme.Gateway
}

func SetupGateway(DatabaseClient *sqlx.DB, StorageClient *s3.S3) GatewayModule {
	return GatewayModule{
		UserGateway:  userGateway.SetupUserGateway(DatabaseClient),
		ThemeGateway: themeGateway.SetupThemeGateway(DatabaseClient, StorageClient),
	}
}

// UseCaseModule section
type UseCaseModule struct {
	UserUseCase  user.UseCase
	ThemeUseCase theme.UseCase
}

func SetupUseCase(gatewayModule GatewayModule) UseCaseModule {
	return UseCaseModule{
		UserUseCase:  userUsecase.SetupUserUseCase(gatewayModule),
		ThemeUseCase: themeUsecase.SetupThemeUseCase(gatewayModule),
	}
}

// DelegateModule section
type DelegateModule struct {
	UserDelegate  user.Delegate
	ThemeDelegate theme.Delegate
}

func SetupDelegate(usecaseModule UseCaseModule) DelegateModule {
	return DelegateModule{
		UserDelegate:  userDelegate.SetupUserDelegate(usecaseModule),
		ThemeDelegate: themeDelegate.SetupThemeDelegate(usecaseModule),
	}
}

// HandlerModule section
type HandlerModule struct {
	UserHandler  userHttp.Handler
	ThemeHandler themeHttp.Handler
}

func SetupHandler(delegate DelegateModule) HandlerModule {
	return HandlerModule{
		UserHandler:  userHttp.SetupUserHandler(delegate.UserDelegate),
		ThemeHandler: themeHttp.SetupThemeHandler(delegate.ThemeDelegate),
	}
}
