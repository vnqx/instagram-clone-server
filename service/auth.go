package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type SignUpInput struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}


func (api *api) ValidateUser(email, password string) (*User, error) {
	user, err := api.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	isPasswordValid := checkPasswordHash(password, user.PasswordHash)
	if !isPasswordValid {
		return nil, errors.New(http.StatusText(http.StatusUnauthorized))
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func checkPasswordHash(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}