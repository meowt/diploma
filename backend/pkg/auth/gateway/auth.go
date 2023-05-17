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

func (au *AuthGatewayImpl) LogIn(user *models.UserUsecase) (gotUser *models.UserUsecase, err error) {
	dbUser := models.UserDB{}
	query := fmt.Sprintf(
		"SELECT * FROM public.users WHERE email = '%v' OR username = '%v'",
		user.Email, user.Username)
	err = au.DatabaseClient.QueryRowx(query).StructScan(&dbUser)
	if err != nil {
		log.Println("[ERROR] SQL query didn't execute:", query)
		err = au.ErrCreator.New(err)
		return
	}
	log.Println("SQL query executed:", query)
	gotUser = dbUser.ToUsecase()
	return
}

func (au *AuthGatewayImpl) SaveRefreshToken(user *models.UserUsecase, refreshToken string) (err error) {
	query := fmt.Sprintf(
		"DELETE FROM public.tokens WHERE user_id = '%v';"+
			"INSERT INTO tokens (token, user_id) VALUES ('%v', '%v')",
		user.Id, refreshToken, user.Id)
	_, err = au.DatabaseClient.Exec(query)
	if err != nil {
		log.Println("[ERROR] SQL query didn't execute:", query)
		return au.ErrCreator.New(err)
	}
	log.Println("SQL query executed:", query)
	return
}

func (au *AuthGatewayImpl) CheckRefreshToken(user *models.UserUsecase, oldRefreshToken string) (err error) {
	query := fmt.Sprintf(
		"SELECT * FROM public.tokens WHERE user_id = '%v' AND token = '%v'",
		user.Id, oldRefreshToken)
	_, err = au.DatabaseClient.Exec(query)
	if err != nil {
		log.Println("[ERROR] SQL query didn't execute:", query)
		err = au.ErrCreator.New(err)
		return
	}
	log.Println("SQL query executed:", query)
	return
}
