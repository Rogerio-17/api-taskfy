package main

import (
	"fmt"
	"log"
	"net/http"
	"taskify/internal/handler"
	"taskify/internal/usecase"
	"taskify/middleware"
	"taskify/repository"
)

func main() {
	// Repositórios
	userInMemoryRepository := repository.NewUserRepositoryInMemory()
	taskInMemoryRepository := repository.NewTaskRepositoryInMemory()

	// Casos de uso
	userUseCase := usecase.NewUserUseCase(userInMemoryRepository)
	taskUseCase := usecase.NewTaskUseCase(taskInMemoryRepository)

	// Handlers
	userHandler := handler.NewUserHandler(userUseCase)
	taskHandler := handler.NewTaskHandler(taskUseCase)

	router := http.NewServeMux()

	// Middlewares
	authMiddleare := middleware.NewAuthMiddleware(userInMemoryRepository)

	// Rotas publicas
	router.HandleFunc("POST /users", userHandler.CreateUser)
	router.HandleFunc("POST /login", userHandler.Login)

	// Rotas autenticadas
	router.HandleFunc("GET /tasks", authMiddleare.VerifyAuthentication(taskHandler.ListTasks))

	port := ":8080"
	fmt.Printf("Servidor rodando na porta %s\n", port)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
