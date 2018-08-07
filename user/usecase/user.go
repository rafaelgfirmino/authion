package usecase

import (
	"github.com/rafaelgfirmino/authion/exceptions"
	"github.com/rafaelgfirmino/authion/user/domain"
	"github.com/rafaelgfirmino/authion/user/repository"
)

type UserUseCase interface {
	FindByID(id int64) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	ConfirmationToken(token string) error
	RegisterNewUser(user *domain.User) (*domain.User, error)
	Authenticate(user *domain.User) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(repository repository.UserRepository) UserUseCase {
	return &userUsecase{
		userRepository: repository,
	}
}

func (usecase *userUsecase) FindByEmail(email string) (*domain.User, error) {
	return usecase.userRepository.FindByEmail(email)
}

func (usecase *userUsecase) FindByID(id int64) (*domain.User, error) {

	existUser, err := usecase.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return existUser, nil
}

func (usecase *userUsecase) ConfirmationToken(confirmationToken string) error {
	if len(confirmationToken) < 10 {
		return exceptions.ErrorTokenNotFound
	}
	return usecase.userRepository.ConfirmationToken(confirmationToken)
}

func (usecase *userUsecase) RegisterNewUser(user *domain.User) (*domain.User, error) {
	return usecase.userRepository.RegisterNewUser(user)
}

func (usecase *userUsecase) Authenticate(user *domain.User) error {
	return usecase.userRepository.Authenticate(user)
}
