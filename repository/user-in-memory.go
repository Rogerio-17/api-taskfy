package repository

import (
	"fmt"
	"taskify/internal/domain"
	apperros "taskify/pkg/errors"
)

// UserRepositoryInMemory repositório de usuários em memória
type UserRepositoryInMemory struct {
	users map[string]*domain.User
}

// NewUserRepositoryInMemory cria uma instancia do repositório de usuários em memória
func NewUserRepositoryInMemory() *UserRepositoryInMemory {
	return &UserRepositoryInMemory{
		users: map[string]*domain.User{},
	}
}

func (r *UserRepositoryInMemory) Create(newUser *domain.User) error {
	for _, u := range r.users {
		if u.Email == newUser.Email {
			return apperros.ErrEmailAlreadyExists
		}
	}

	r.users[newUser.Id] = newUser
	fmt.Println(r.users)
	return nil
}

func (r *UserRepositoryInMemory) FindByEmail(email string) (*domain.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, apperros.ErrUserNotFound
}

func (r *UserRepositoryInMemory) GetById(id string) (*domain.User, error) {
	if _, ok := r.users[id]; !ok {
		return nil, apperros.ErrUserNotFound
	}

	return r.users[id], nil
}
