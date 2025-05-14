package ports

import "github.com/gian-guerra/task-manager-go-level-2/models"

type TaskRepositoryPort interface {
	Save(task models.Task) models.Task
	GetAll() []models.Task
	Complete(id string) error
	Delete(id string) error
}
