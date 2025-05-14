package models

import "time"

type Task struct {
	ID        string
	Title     string
	Done      bool
	CreatedAt time.Time
}

type TaskList []Task
