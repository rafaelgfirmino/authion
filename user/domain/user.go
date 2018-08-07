package domain

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID                int64     `json:"id,omitempty"`
	Email             string    `json:"email" validate:"required,email"`
	Password          string    `json:"password,omitempty" validate:"required"`
	Phone             string    `json:"phone,omitempty"`
	ConfirmationToken string    `json:"-"`
	Enabled           bool      `json:"-"`
	PasswordRequestAt time.Time `json:"-"`
	PasswordToken     string    `json:"password_token,omitempty"`
}

var ErroEncriptyPassword = errors.New("can not encrypt the password")

func (user User) GetPassword() ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, ErroEncriptyPassword
	}
	return hashedPassword, nil
}

func (user User) CheckPasswordHash(hash string) bool {

	fmt.Println(hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password))

	return err == nil
}
