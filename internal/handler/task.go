package handler

import (
	"net/http"
	"taskify/internal/helpers"
	"taskify/internal/usecase"
	"taskify/middleware"
)

type TaskHandler struct {
	taskUseCase *usecase.TaskUseCase
}

func NewTaskHandler(taskUseCase *usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		taskUseCase: taskUseCase,
	}
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

	var tasks []*TaskResponse

	userTasks, err := h.taskUseCase.ListAll(userId)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	for _, t := range userTasks {
		tasks = append(tasks, &TaskResponse{
			ID:          t.Id,
			UserId:      t.UserId,
			Title:       t.Title,
			Description: t.Description,
			IsCompleted: t.IsCompleted,
			CreatedAt:   t.CreatedAt.String(),
			UpdatedAt:   t.UpdatedAt.String(),
		})
	}

	helpers.ResponseWithJSON(w, http.StatusOK, tasks)
}
