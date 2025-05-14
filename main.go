package main

import (
	"bufio"
	"fmt"

	"os"
	"strings"

	"github.com/gian-guerra/task-manager-go-level-2/core"
	"github.com/gian-guerra/task-manager-go-level-2/utils"
)

func main() {
	repository := core.NewRepository()
	service := core.NewTaskService(repository)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("commands: add, list, complete <id>, delete <id>, exit")
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]
		switch cmd {
		case "add":
			fmt.Print("Enter task title: ")

			scanner.Scan()
			tittle := scanner.Text()
			task := service.AddTask(tittle)

			utils.PrintSuccessExecution(task.ID)
		case "list":
			tasks := service.ListTasks()
			utils.PrintTasks(tasks)

		case "complete":
			if len(parts) < 2 {
				utils.PrintMissingID()
				continue
			}
			id := parts[1]
			err := service.CompleteTask(id)
			if err != nil {
				utils.PrintError(err)
				return
			}

			utils.PrintSuccessExecution(id)
		case "delete":
			if len(parts) < 2 {
				utils.PrintMissingID()
				continue
			}
			id := parts[1]
			err := service.DeleteTask(id)
			if err != nil {
				utils.PrintError(err)
				return
			}
			utils.PrintSuccessExecution(id)
		case "exit":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}
