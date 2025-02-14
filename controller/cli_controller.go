package controller

import (
	"strconv"
	"task-cli/enums"
	"task-cli/service"
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
	default:
		panic("Not a valid operation")
	}
}

func addTask(desc string) {
	service.AddTask(desc)
}

func listTasks(status string) {
	service.ListTasks(status)
}

func deleteTask(id int) {
	service.DeleteTask(id)
}

func updateTask(id int, desc string) {
	service.UpdateTaskDescription(id, desc)
}

func updateStatus(id int, status enums.TaskState) {
	service.UpdateTaskStatus(id, status)
}
