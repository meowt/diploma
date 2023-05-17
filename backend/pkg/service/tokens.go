package service

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"Diploma/pkg/models"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

// TokenManager provides logic for JWT & Refresh tokens generation and parsing.
type TokenManager interface {
	NewJWT(userId string, ttl time.Duration) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	SigningKey string
}

func NewManager(signingKey string) *Manager {
	return &Manager{SigningKey: signingKey}
}

func (m *Manager) NewJWT(user *models.UserUsecase) (res string, err error) {
	claims := models.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(viper.GetDuration("auth.access_token_ttl") * time.Second).Unix(),
		},
		UserIdentity: models.UserIdentity{
			UserId:   user.Id,
			Username: user.Username,
		},
	}
	ss := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	res, err = ss.SignedString([]byte(m.SigningKey))
	if err != nil {
		err = errors.New("error while signing jwt token")
		return
	}
	log.Printf("Generated access token for %v: %v", user.Username, res)
	return
}

func (m *Manager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.SigningKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("error get user claims from token")
	}

	return claims["sub"].(string), nil
}

func (m *Manager) NewRefreshToken() (str string, err error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err = r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
