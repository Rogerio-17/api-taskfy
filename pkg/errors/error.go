package errors

import "errors"

var (
	// General errors
	ErrInternalServerError = errors.New("Internal server error")
	ErrInvalidRequest      = errors.New("Invalid request")
	ErrInvalidData         = errors.New("Invalid data")

	// User errors
	ErrEmptyName          = errors.New("Name is required")
	ErrEmptyEmail         = errors.New("Email is required")
	ErrEmptyPassword      = errors.New("Password is required")
	ErrInvalidCredentials = errors.New("Invalid email or password")
	ErrUserNotFound       = errors.New("User not found")
	ErrEmailAlreadyExists = errors.New("Email already exists")

	// Task errors
	ErrEmptyUserId  = errors.New("User is requered")
	ErrEmptyTitle   = errors.New("Title is required")
	ErrTaskNotFound = errors.New("Task not found")
	ErrUnauthorized = errors.New("Unauthorized")

	// Auth errors
	ErrInvalidAuthorization = errors.New("Token não informado")
	ErrInvalidToken         = errors.New("token inválido")
)
