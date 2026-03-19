package handler

import (
	"encoding/json"
	"net/http"
	"taskify/internal/domain"
	"taskify/internal/helpers"
	"taskify/internal/usecase"
	"taskify/middleware"
	"taskify/pkg/errors"
	"time"
)

type TaskHandler struct {
	taskUseCase *usecase.TaskUseCase
}

func NewTaskHandler(taskUseCase *usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		taskUseCase: taskUseCase,
	}
}

// CreateTaskRequest representa a estrutura de dados para criar uma tarefa
type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	userId, err := middleware.GetUserIDFromContext(r.Context())

	if err != nil {
		helpers.ResponseWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	var requestBody CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, errors.ErrInvalidData.Error())
		return
	}

	taskCreated, err := h.taskUseCase.Create(userId, requestBody.Title, requestBody.Description)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseBody := fromTaskToResponse(taskCreated)

	helpers.ResponseWithJSON(w, http.StatusOK, responseBody)
}

type TaskResponse struct {
	ID          string `json:"id"`
	UserId      string `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	userId, err := middleware.GetUserIDFromContext(r.Context())

	if err != nil {
		helpers.ResponseWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	tasks := []*TaskResponse{}

	userTasks, err := h.taskUseCase.ListAll(userId)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	for _, t := range userTasks {
		tasks = append(tasks, fromTaskToResponse(t))
	}

	helpers.ResponseWithJSON(w, http.StatusOK, tasks)
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	userId, err := middleware.GetUserIDFromContext(r.Context())

	if err != nil {
		helpers.ResponseWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	taskId := r.PathValue("id")

	if taskId == "" {
		helpers.ResponseWithError(w, http.StatusBadRequest, errors.ErrInvalidRequest.Error())
		return
	}

	var requestBody UpdateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, errors.ErrInvalidData.Error())
		return
	}

	taskUpdated, err := h.taskUseCase.Update(userId, taskId, requestBody.Title, requestBody.Description)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseBody := fromTaskToResponse(taskUpdated)

	helpers.ResponseWithJSON(w, http.StatusOK, responseBody)
}

func (h *TaskHandler) MarkTaskAsCompleted(w http.ResponseWriter, r *http.Request) {
	userId, err := middleware.GetUserIDFromContext(r.Context())

	if err != nil {
		helpers.ResponseWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	taskId := r.PathValue("id")

	taskUpdated, err := h.taskUseCase.MarkAsCompleted(userId, taskId)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseBody := fromTaskToResponse(taskUpdated)

	helpers.ResponseWithJSON(w, http.StatusOK, responseBody)
}

func (h *TaskHandler) MarkTaskAsIncomplete(w http.ResponseWriter, r *http.Request) {
	userId, err := middleware.GetUserIDFromContext(r.Context())

	if err != nil {
		helpers.ResponseWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	taskId := r.PathValue("id")

	taskUpdated, err := h.taskUseCase.MarkAsIncomplete(userId, taskId)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseBody := fromTaskToResponse(taskUpdated)

	helpers.ResponseWithJSON(w, http.StatusOK, responseBody)
}

func fromTaskToResponse(task *domain.Task) *TaskResponse {
	return &TaskResponse{
		ID:          task.Id,
		UserId:      task.UserId,
		Title:       task.Title,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CreatedAt:   task.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   task.UpdatedAt.Format(time.RFC3339),
	}
}
