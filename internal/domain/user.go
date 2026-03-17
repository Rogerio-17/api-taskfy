package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(newUser *User) error
	FindByEmail(email string) (*User, error)
	GetById(id string) (*User, error)
}

type User struct {
	Id        string
	Email     string
	Name      string
	Password  string
	CreatedAt time.Time
}

func NewUser(name, email, password string) *User {
	return &User{
		Id:        uuid.New().String(),
		Email:     email,
		Name:      name,
		Password:  password,
		CreatedAt: time.Now(),
	}
}
