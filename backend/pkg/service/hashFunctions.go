package service

import "golang.org/x/crypto/bcrypt"

type HashManager interface {
	HashPassword(password string) (string, error)
	ComparePassword(password, hash string) bool
}

type BCrypter struct {
	rndm string
}

func NewBCrypter() *BCrypter {
	return new(BCrypter)
}

func (b *BCrypter) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (b *BCrypter) ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
