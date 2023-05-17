package models

import (
	"github.com/golang-jwt/jwt"
)

type UserUpdateInput struct {
	UpdatingUserId uint   `json:"updating_user_id,omitempty"`
	NewUsername    string `json:"new_username"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
}

type UserClaims struct {
	jwt.StandardClaims
	UserIdentity
}

type UserIdentity struct {
	UserId   uint
	Username string
}

type UserHttp struct {
	DefaultModel
	Username  string `json:"username"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	Followers int    `json:"followers"`
}

type UserUsecase struct {
	DefaultModel
	Username  string
	Firstname string
	Lastname  string
	Email     string
	Password  string
	Followers int
}

type UserDB struct {
	DBDefaultModel
	Username      string
	Firstname     string
	Lastname      string
	Email         string
	Password_Hash string
	Followers     int
}

func (user *UserUsecase) ToHttp() *UserHttp {
	return &UserHttp{
		DefaultModel: user.DefaultModel,
		Username:     user.Username,
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
		Email:        user.Email,
		//Password:     user.Password,
		Followers: user.Followers,
	}
}

func (user *UserHttp) ToUsecase() *UserUsecase {
	return &UserUsecase{
		DefaultModel: user.DefaultModel,
		Username:     user.Username,
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
		Email:        user.Email,
		Password:     user.Password,
		Followers:    user.Followers,
	}
}

func (user *UserDB) ToUsecase() *UserUsecase {
	return &UserUsecase{
		DefaultModel: user.ToDefaultModel(),
		Username:     user.Username,
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
		Email:        user.Email,
		Password:     user.Password_Hash,
		Followers:    user.Followers,
	}
}

func (user *UserUsecase) ToDB() *UserDB {
	return &UserDB{
		DBDefaultModel: user.ToDBDefaultModel(),
		Username:       user.Username,
		Firstname:      user.Firstname,
		Lastname:       user.Lastname,
		Email:          user.Email,
		Password_Hash:  user.Password,
		Followers:      user.Followers,
	}
}
