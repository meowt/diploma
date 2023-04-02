package gateway

import (
	"Diploma/pkg/theme"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmoiron/sqlx"
)

type ThemeGatewayImpl struct {
	DatabaseClient *sqlx.DB
	StorageClient  *s3.S3
}

type ThemeGatewayModule struct {
	theme.Gateway
}

func SetupThemeGateway(databaseClient *sqlx.DB, storageClient *s3.S3) ThemeGatewayModule {
	return ThemeGatewayModule{
		Gateway: &ThemeGatewayImpl{DatabaseClient: databaseClient, StorageClient: storageClient},
	}
}
