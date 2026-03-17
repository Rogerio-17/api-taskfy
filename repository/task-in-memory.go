package repository

import (
	"fmt"
	"taskify/internal/domain"
	apperros "taskify/pkg/errors"
)

// TaskRepositoryInMemory repositório de usuários em memória
type TaskRepositoryInMemory struct {
	tasks map[string]*domain.Task
}

// NewTaskRepositoryInMemory cria uma instancia do repositório de usuários em memória
func NewTaskRepositoryInMemory() *TaskRepositoryInMemory {
	return &TaskRepositoryInMemory{
		tasks: map[string]*domain.Task{},
	}
}

func (r *TaskRepositoryInMemory) Create(newTask *domain.Task) error {
	r.tasks[newTask.Id] = newTask
	fmt.Println(r.tasks)
	return nil
}

func (r *TaskRepositoryInMemory) GetById(id string) (*domain.Task, error) {
	if _, ok := r.tasks[id]; !ok {
		return nil, apperros.ErrTaskNotFound
	}

	return r.tasks[id], nil
}

func (r *TaskRepositoryInMemory) FindMany(userId string) ([]*domain.Task, error) {
	var userTasks []*domain.Task

	for _, t := range r.tasks {
		if t.UserId != userId {
			continue
		}
		userTasks = append(userTasks, t)
	}

	return userTasks, nil
}
