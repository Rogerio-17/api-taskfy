package usecase

import (
	"taskify/internal/domain"
	"taskify/pkg/errors"
)

type TaskUseCase struct {
	taskRepository domain.TaskRepository
}

// NewTaskUseCase cria uma instancia dos casos de uso de usuário
func NewTaskUseCase(taskRepository domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepository: taskRepository,
	}
}

func (uc *TaskUseCase) CreateTask(userId, title, description string) (*domain.Task, error) {
	if userId == "" {
		return nil, errors.ErrEmptyUserId
	}

	if title == "" {
		return nil, errors.ErrEmptyTitle
	}

	task := domain.NewTask(userId, title, description)

	err := uc.taskRepository.Create(task)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *TaskUseCase) ListAll(userId string) ([]*domain.Task, error) {
	return uc.taskRepository.FindMany(userId)
}
