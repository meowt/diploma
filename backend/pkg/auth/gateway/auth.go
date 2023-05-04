package gateway

import (
	"fmt"

	"Diploma/pkg/auth"
	"Diploma/pkg/errorPkg"
	"Diploma/pkg/models"

	"github.com/jmoiron/sqlx"
)

type AuthGatewayImpl struct {
	DatabaseClient *sqlx.DB
	ErrCreator     errorPkg.ErrCreator
}

type AuthGatewayModule struct {
	auth.Gateway
}

func SetupAuthGateway(databaseClient *sqlx.DB, errCreator errorPkg.ErrCreator) AuthGatewayModule {
	return AuthGatewayModule{
		Gateway: &AuthGatewayImpl{DatabaseClient: databaseClient, ErrCreator: errCreator},
	}
}

func (au *AuthGatewayImpl) SignUp(user *models.UserUsecase) (err error) {
	_, err = au.DatabaseClient.Exec(fmt.Sprintf(
		"INSERT INTO public.users (created_at, username, firstname, lastname, email, password_hash) VALUES (current_timestamp, '%v', '%v', '%v', '%v', '%v')",
		user.Username, user.Firstname, user.Lastname, user.Email, user.Password))
	if err != nil {
		return au.ErrCreator.New(err)
	}
	return nil
}

func (au *AuthGatewayImpl) LogIn(user *models.UserUsecase) (err error) {
	var userDB, emptyUserDB models.UserDB
	err = au.DatabaseClient.QueryRow(fmt.Sprintf(
		"SELECT * FROM public.users WHERE email = '%v' AND password_hash = '%v'", user.Email, user.Password)).Scan(&userDB)
	if err != nil {
		return au.ErrCreator.New(err)
	}
	if userDB == emptyUserDB {
		//TODO: create error
		return au.ErrCreator.New(fmt.Errorf(""))
	}
	return nil
}
