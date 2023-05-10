package gateway

import (
	"fmt"

	"Diploma/pkg/errorPkg"
	"Diploma/pkg/models"
	"Diploma/pkg/user"

	"github.com/jmoiron/sqlx"
)

type UserGatewayImpl struct {
	DatabaseClient *sqlx.DB
	ErrCreator     errorPkg.ErrCreator
}

type UserGatewayModule struct {
	user.Gateway
}

func SetupUserGateway(postgresClient *sqlx.DB) UserGatewayModule {
	return UserGatewayModule{
		Gateway: &UserGatewayImpl{DatabaseClient: postgresClient},
	}
}

func (u *UserGatewayImpl) GetUserByEmailOrUsername(useCaseUser *models.UserUsecase) (User *models.UserUsecase, err error) {
	var dbUser models.UserDB
	err = u.DatabaseClient.QueryRowx(fmt.Sprintf(
		"SELECT * FROM public.users WHERE email = '%v' OR username = '%v'",
		useCaseUser.Email, useCaseUser.Username)).StructScan(&dbUser)
	if err != nil {
		err = u.ErrCreator.New(err)
		return
	}
	return dbUser.ToUsecase(), err
}
