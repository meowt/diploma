package service

import (
	"testing"
	"time"

	"Diploma/pkg/models"
)

func TestManager_NewRefreshToken(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		mngr := NewManager("super_secret_access_signing_key")
		res, err := mngr.NewRefreshToken()
		if err != nil {
			t.Errorf("Smthng went wrong: %v", err)
			return
		}
		t.Log(res)
	})
}

func TestManager_NewJWT(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		mngr := NewManager("super_secret_access_signing_key")
		user := &models.UserUsecase{
			DefaultModel: models.DefaultModel{
				Id:        3,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: time.Now(),
			},
			Username:  "meowt",
			Firstname: "ilya",
			Lastname:  "matveev",
			Email:     "meowt@gmail.com",
			Password:  "123",
			Followers: 0,
		}
		res, err := mngr.NewJWT(user)
		if err != nil {
			t.Errorf("Smthng went wrong: %v", err)
			return
		}
		t.Log(res)
	})
}
