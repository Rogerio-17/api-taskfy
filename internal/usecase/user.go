package usecase

import (
	"taskify/internal/domain"
	"taskify/pkg/errors"
)

type UserUseCase struct {
	userRepository domain.UserRepository
}

// NewUserUseCase cria uma instancia dos casos de uso de usuário
func NewUserUseCase(paramUserRepository domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: paramUserRepository,
	}
}

func (uc *UserUseCase) CreateUser(name, email, password string) (*domain.User, error) {
	if name == "" {
		return nil, errors.ErrEmptyName
	}

	if email == "" {
		return nil, errors.ErrEmptyEmail
	}

	if password == "" {
		return nil, errors.ErrEmptyPassword
	}

	user := domain.NewUser(name, email, password)

	err := uc.userRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) Login(email, password string) (*domain.User, error) {
	if email == "" {
		return nil, errors.ErrEmptyEmail
	}

	if password == "" {
		return nil, errors.ErrEmptyPassword
	}

	user, err := uc.userRepository.FindByEmail(email)

	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	if user.Password != password {
		return nil, errors.ErrInvalidCredentials
	}

	return user, nil
}
