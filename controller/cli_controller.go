package controller

import (
	"fmt"
	"strconv"
	"task-tracker-cli/enums"
)

func ProcessArgs(args []string) {
	switch args[1] {
	case "add":
		argument := ""
		if len(args) >= 3 {
			argument = args[2]
		}
		addTask(argument)

	case "list":
		argument := ""
		if len(args) >= 3 {
			argument = args[2]
		}
		listTasks(argument)

	case "delete":
		argument := 0
		if len(args) >= 3 {
			argument, _ = strconv.Atoi(args[2])
		}
		deleteTask(argument)

	case "update":
		if len(args) >= 4 {
			arg1, _ := strconv.Atoi(args[2])
			updateTask(arg1, args[3])
		} else {
			panic("Not enough arguments")
		}

	case "mark-in-progress":
		if len(args) >= 3 {
			argument, _ := strconv.Atoi(args[2])
			updateStatus(argument, enums.InProgress)
		}

	case "mark-done":
		if len(args) >= 3 {
			argument, _ := strconv.Atoi(args[2])
			updateStatus(argument, enums.Done)
		}
	}
}

func addTask(desc string) {
	fmt.Print("consuming service to add a task: " + desc)
}

func listTasks(status string) {
	fmt.Print("listing tasks: " + status)
}

func deleteTask(id int) {
	fmt.Print("deleting task: " + string(id))
}

func updateTask(id int, desc string) {
	fmt.Print("Updating task " + string(id) + " with " + desc)
}

func updateStatus(id int, status enums.TaskState) {
	fmt.Print("Changing status of task " + string(id) + " to " + status.String())

}
