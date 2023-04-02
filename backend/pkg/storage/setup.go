package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

func Setup() (svc *s3.S3, err error) {
	access := viper.GetString("storage.accessKey")
	secret := viper.GetString("storage.secretKey")

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ru-msk"),
		Endpoint:    aws.String(viper.GetString("storage.endpointUrl")),
		Credentials: credentials.NewStaticCredentials(access, secret, ""),
	}))

	svc = s3.New(sess)
	return
}
