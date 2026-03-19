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

func (uc *TaskUseCase) Create(userId, title, description string) (*domain.Task, error) {
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

func (uc *TaskUseCase) Update(userId, taskId, title, description string) (*domain.Task, error) {
	if title == "" {
		return nil, errors.ErrEmptyTitle
	}

	task, err := uc.taskRepository.GetById(taskId)

	if err != nil {
		return nil, err
	}

	if task.UserId != userId {
		return nil, errors.ErrUnauthorized
	}

	task.UpdateTask(title, description)

	err = uc.taskRepository.Update(task)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *TaskUseCase) MarkAsCompleted(userId, taskId string) (*domain.Task, error) {
	task, err := uc.taskRepository.GetById(taskId)

	if err != nil {
		return nil, err
	}

	if task.UserId != userId {
		return nil, errors.ErrUnauthorized
	}

	task.MarkAsCompleted()

	err = uc.taskRepository.Update(task)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (uc *TaskUseCase) MarkAsIncomplete(userId, taskId string) (*domain.Task, error) {
	task, err := uc.taskRepository.GetById(taskId)

	if err != nil {
		return nil, err
	}

	if task.UserId != userId {
		return nil, errors.ErrUnauthorized
	}

	task.MarkAsIncomplete()

	err = uc.taskRepository.Update(task)

	if err != nil {
		return nil, err
	}

	return task, nil
}
