package models

import "time"

type Task struct {
	ID      string
	Title   string
	Done    bool
	DueData time.Time
}

type TaskList []Task
