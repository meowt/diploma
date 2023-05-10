package models

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lib/pq"
)

type UserClaims struct {
	jwt.StandardClaims
	Username string
}

type UserDB struct {
	Id            int         `json:"id,omitempty"`
	Created_At    pq.NullTime `json:"created_at"`
	Updated_At    pq.NullTime `json:"updated_at"`
	Deleted_At    pq.NullTime `json:"deleted_at"`
	Username      string      `json:"username,omitempty"`
	Firstname     string      `json:"firstname,omitempty"`
	Lastname      string      `json:"lastname,omitempty"`
	Email         string      `json:"email,omitempty"`
	Password_Hash string      `json:"passwordHash,omitempty"`
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

func (user *UserHttp) ToUsecase() *UserUsecase {
	return &UserUsecase{
		Email:    user.Email,
		Password: user.Password,
	}
}

func (user *UserUsecase) ToDB() *UserDB {
	return &UserDB{
		Username:      user.Username,
		Firstname:     user.Firstname,
		Lastname:      user.Lastname,
		Email:         user.Email,
		Password_Hash: user.Password,
	}
}

func (user *UserDB) ToUsecase() *UserUsecase {
	return &UserUsecase{
		CreatedAt: user.Created_At.Time,
		UpdatedAt: user.Updated_At.Time,
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}
}
