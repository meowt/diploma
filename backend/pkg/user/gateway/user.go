package gateway

import (
	"database/sql"
	"fmt"
	"log"

	"Diploma/pkg/errorPkg"
	"Diploma/pkg/models"
	"Diploma/pkg/user"

	"github.com/jmoiron/sqlx"
)

type UserGatewayImpl struct {
	DatabaseClient *sqlx.DB
	ErrCreator     *errorPkg.ErrorCreator
}

type UserGatewayModule struct {
	user.Gateway
}

func SetupUserGateway(postgresClient *sqlx.DB, creator *errorPkg.ErrorCreator) UserGatewayModule {
	return UserGatewayModule{
		Gateway: &UserGatewayImpl{DatabaseClient: postgresClient, ErrCreator: creator},
	}
}

func (u *UserGatewayImpl) GetUserByEmailOrUsername(useCaseUser *models.UserUsecase) (User *models.UserUsecase, err error) {
	var dbUser models.UserDB
	query := fmt.Sprintf(
		"SELECT * FROM public.users WHERE email = '%v' OR username = '%v' LIMIT 1",
		useCaseUser.Email, useCaseUser.Username)
	err = u.DatabaseClient.QueryRowx(query).StructScan(&dbUser)
	switch err {
	case nil:
		log.Println("SQL query executed:", query)
	case sql.ErrNoRows:
		err = u.ErrCreator.NewErrSQLNoRows()
		return
	default:
		log.Printf("[ERROR] SQL query didn't execute: %v, error: %v", query, err)
		return
	}

	return dbUser.ToUsecase(), err
}

func (u *UserGatewayImpl) CreateUser(user *models.UserUsecase) (err error) {
	query := fmt.Sprintf(
		"INSERT INTO public.users (created_at, username, firstname, lastname, email, password_hash) VALUES (current_timestamp, '%v', '%v', '%v', '%v', '%v') RETURNING id",
		user.Username, user.Firstname, user.Lastname, user.Email, user.Password)
	err = u.DatabaseClient.QueryRowx(query).Scan(&user.Id)
	if err != nil {
		switch err {
		//TODO: implement db duplicate error
		default:
			log.Printf("[ERROR] SQL query didn't execute: %v, error: %v\n", query, err)
			return
		}
	}
	log.Println("SQL query executed:", query)
	return
}

func (u *UserGatewayImpl) GetUserById(userId uint) (User *models.UserUsecase, err error) {
	var dbUser models.UserDB
	query := fmt.Sprintf("SELECT * FROM public.users WHERE id = '%v' LIMIT 1", userId)
	err = u.DatabaseClient.QueryRowx(query).StructScan(&dbUser)
	switch err {
	case nil:
		log.Println("SQL query executed:", query)
	case sql.ErrNoRows:
		err = u.ErrCreator.NewErrSQLNoRows()
		return
	default:
		log.Printf("[ERROR] SQL query didn't execute: %v, error: %v\n", query, err)
		return
	}

	return dbUser.ToUsecase(), err
}

func (u *UserGatewayImpl) UpdateUser(UpdateUser *models.UserUpdateInput) (user *models.UserUsecase, err error) {
	var dbUser models.UserDB
	query := fmt.Sprintf("UPDATE users SET username = '%v', firstname = '%v', lastname = '%v', updated_at = current_timestamp\nWHERE id = '%v'\nRETURNING *;",
		UpdateUser.NewUsername, UpdateUser.Firstname, UpdateUser.Lastname, UpdateUser.UpdatingUserId)
	err = u.DatabaseClient.QueryRowx(query).StructScan(&dbUser)
	switch err {
	case nil:
		log.Println("SQL query executed:", query)
	case sql.ErrNoRows:
		err = u.ErrCreator.NewErrSQLNoRows()
		return
	default:
		log.Printf("[ERROR] SQL query didn't execute: %v, error: %v\n", query, err)
		return
	}

	return dbUser.ToUsecase(), err
}
