package utils

import (
	"fmt"

	"github.com/gian-guerra/task-manager-go-level-2/models"
)

func PrintTasks(tasks []models.Task) {
	for _, t := range tasks {
		status := "❌"
		if t.Done {
			status = "✅"
		}
		fmt.Printf("[%s] %s - %s (%s)\n", t.ID, status, t.Title, t.CreatedAt.Format("2006-01-02"))
	}
}

func PrintMissingID() {
	fmt.Println("Please provide the task ID")
}

func PrintSuccessExecution(id string) {
	fmt.Printf("Operation successed for task: %s\n", id)
}

func PrintError(err error) {
	fmt.Println("Error in execution: ", err)
}
