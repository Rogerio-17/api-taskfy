package handler

import (
	"encoding/json"
	"net/http"
	"taskify/internal/helpers"
	"taskify/internal/usecase"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// CreateUserRequest representa a estrutura de dados para criar um usuário
type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserResponse representa a estrutura de dados para a resposta de criação de um usuário
type CreateUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// CreateUser é o handler para criar um novo usuário
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestBody CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Dados inválidos")
		return
	}

	createdUser, err := h.userUseCase.CreateUser(requestBody.Name, requestBody.Email, requestBody.Password)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseBody := CreateUserResponse{
		ID:    createdUser.Id,
		Email: createdUser.Email,
	}

	helpers.ResponseWithJSON(w, http.StatusOK, responseBody)
}

// LoginRequest representa a estrutura de dados para autenticar um usuário
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

// Login é o handler para autenticar um usuário
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var requestBody LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Dados invalidos")
		return
	}

	authenticatedUser, err := h.userUseCase.Login(requestBody.Email, requestBody.Password)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseBody := LoginResponse{
		ID:    authenticatedUser.Id,
		Email: authenticatedUser.Email,
		Token: authenticatedUser.Id,
	}

	helpers.ResponseWithJSON(w, http.StatusOK, responseBody)
}
