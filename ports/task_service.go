package ports

import "github.com/gian-guerra/task-manager-go-level-2/models"

type TaskServicePort interface {
	AddTask(title string) models.Task
	CompleteTask(id string) error
	DeleteTask(id string) error
	ListTasks() []models.Task
}
