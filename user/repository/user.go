package repository

import "github.com/rafaelgfirmino/authion/user/domain"

type UserRepository interface {
	FindByID(id int64) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	ConfirmationToken(confirmationToken string) error
	RegisterNewUser(user *domain.User) (*domain.User, error)
	Authenticate(user *domain.User) error
}
