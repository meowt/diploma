package gateway

import (
	"fmt"
	"log"

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
	query := fmt.Sprintf(
		"INSERT INTO public.users (created_at, username, firstname, lastname, email, password_hash) VALUES (current_timestamp, '%v', '%v', '%v', '%v', '%v')",
		user.Username, user.Firstname, user.Lastname, user.Email, user.Password)
	_, err = au.DatabaseClient.Exec(query)
	if err != nil {
		return au.ErrCreator.New(err)
	}
	return nil
}

func (au *AuthGatewayImpl) LogIn(user *models.UserUsecase) (passwordHash string, err error) {
	query := fmt.Sprintf(
		"SELECT password_hash FROM public.users WHERE email = '%v' OR username = '%v'",
		user.Email, user.Username)
	err = au.DatabaseClient.QueryRow(query).Scan(&passwordHash)
	if err != nil {
		return passwordHash, au.ErrCreator.New(err)
	}
	return passwordHash, err
}

func (au *AuthGatewayImpl) SaveRefreshToken(username, refreshToken string) (err error) {
	query := fmt.Sprintf(
		"DELETE FROM public.tokens WHERE user_id = (SELECT id FROM public.users WHERE username = '%v');"+
			"INSERT INTO tokens (token, user_id) VALUES ('%v', (SELECT id FROM public.users WHERE username = '%v'))",
		username, refreshToken, username)
	_, err = au.DatabaseClient.Exec(query)
	log.Println(query)
	if err != nil {
		return au.ErrCreator.New(err)
	}
	return nil
}

func (au *AuthGatewayImpl) CheckRefreshToken(username, oldRefreshToken string) (err error) {
	query := fmt.Sprintf(
		"SELECT * FROM public.tokens WHERE user_id = (SELECT id FROM public.users WHERE username = '%v') AND token = '%v'",
		username, oldRefreshToken)
	res, err := au.DatabaseClient.Exec(query)
	log.Println(res)
	return
}
