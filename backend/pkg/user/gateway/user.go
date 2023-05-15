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

func (u *UserGatewayImpl) CreateUser(user *models.UserUsecase) (err error) {
	query := fmt.Sprintf(
		"INSERT INTO public.users (created_at, username, firstname, lastname, email, password_hash) VALUES (current_timestamp, '%v', '%v', '%v', '%v', '%v')",
		user.Username, user.Firstname, user.Lastname, user.Email, user.Password)
	_, err = u.DatabaseClient.Exec(query)
	if err != nil {
		//TODO: implement db duplicate error
		//if errors.As(err, ) {
		//	return
		//}
		return
	}
	return
}
