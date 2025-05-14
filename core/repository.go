package core

import (
	"errors"

	"github.com/gian-guerra/task-manager-go-level-2/models"
	"github.com/gian-guerra/task-manager-go-level-2/ports"
	"github.com/google/uuid"
)

type TaskRepository struct {
	tasks map[string]models.Task
}

var _ ports.TaskRepositoryPort = &TaskRepository{}

func NewRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]models.Task),
	}
}

func (s *TaskRepository) Save(task models.Task) models.Task {
	task.ID = uuid.NewString()
	s.tasks[task.ID] = task
	return task
}

func (s *TaskRepository) GetAll() []models.Task {
	list := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		list = append(list, task)
	}
	return list
}

func (s *TaskRepository) Complete(id string) error {
	task, ok := s.tasks[id]
	if !ok {
		return errors.New("task not found")
	}

	task.Done = true
	s.tasks[id] = task
	return nil
}

func (s *TaskRepository) Delete(id string) error {
	_, ok := s.tasks[id]
	if !ok {
		return errors.New("task not found")
	}
	delete(s.tasks, id)
	return nil
}
