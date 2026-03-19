package domain

import (
	"time"

	"github.com/google/uuid"
)

type TaskRepository interface {
	Create(newUser *Task) error
	GetById(id string) (*Task, error)
	FindMany(userId string) ([]*Task, error)
	Update(task *Task) error
}

type Task struct {
	Id          string
	UserId      string
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewTask é um tipo que representa uma tarefa
func NewTask(userId, title, description string) *Task {
	now := time.Now()

	return &Task{
		Id:          uuid.New().String(),
		UserId:      userId,
		Title:       title,
		Description: description,
		IsCompleted: false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// UpdateTask é um método que atualiza o título e a descrição de uma tarefa
func (t *Task) UpdateTask(newTitle, newDescriptio string) {
	t.Title = newTitle
	t.Description = newDescriptio
	t.UpdatedAt = time.Now()
}

// MarkAsCompleted é um método que marca uma tarefa como concluída
func (t *Task) MarkAsCompleted() {
	if t.IsCompleted {
		return
	}

	t.IsCompleted = true
	t.UpdatedAt = time.Now()
}

// MarkAsIncomplete é um método que marca uma tarefa como incompleta
func (t *Task) MarkAsIncomplete() {
	if !t.IsCompleted {
		return
	}

	t.IsCompleted = false
	t.UpdatedAt = time.Now()
}
