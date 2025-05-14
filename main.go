package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gian-guerra/task-manager-go-level-2/core"
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

			fmt.Printf("Task added successfully with ID %s\n", task.ID)

		case "list":
			tasks := service.ListTasks()
			fmt.Printf("%v\n", tasks)

		case "complete":
			if len(parts) < 2 {
				fmt.Println("Please provide the task ID")
				continue
			}
			id := parts[1]
			err := service.CompleteTask(id)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			fmt.Println("Task completed!")
		case "delete":
			if len(parts) < 2 {
				fmt.Println("Please provide the task ID")
				continue
			}
			id := parts[1]
			err := service.DeleteTask(id)
			if err != nil {
				fmt.Println("Error in delete process: ", err)
				return
			}
			fmt.Println("Task deleted!")
		case "exit":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}
