package models

import "time"

type UserDB struct {
	Id           int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	Username     string
	Firstname    string
	Lastname     string
	Email        string
	PasswordHash string
}

type UserUsecase struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Username  string
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

type UserHttp struct {
	Username  string
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

func (user *UserHttp) ToUsecase() UserUsecase {
	return UserUsecase{
		Email:    user.Email,
		Password: user.Password,
	}
}

func (user *UserUsecase) ToDB() UserDB {
	return UserDB{
		Username:     user.Username,
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
		Email:        user.Email,
		PasswordHash: user.Password,
	}
}

func (user *UserDB) ToUsecase() UserUsecase {
	return UserUsecase{
		CreatedAt: user.CreatedAt,
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}
}
