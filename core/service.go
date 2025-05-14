package core

import (
	"time"

	"github.com/gian-guerra/task-manager-go-level-2/models"
	"github.com/gian-guerra/task-manager-go-level-2/ports"
)

type TaskService struct {
	repository *TaskRepository
}

var _ ports.TaskServicePort = &TaskService{}

func NewTaskService(repository *TaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (s *TaskService) AddTask(title string) models.Task {
	task := models.Task{
		Title:     title,
		CreatedAt: time.Now(),
		Done:      false,
	}
	return s.repository.Save(task)
}

func (s *TaskService) CompleteTask(id string) error {
	return s.repository.Complete(id)
}

func (s *TaskService) DeleteTask(id string) error {
	return s.repository.Delete(id)
}

func (s *TaskService) ListTasks() []models.Task {
	return s.repository.GetAll()
}
