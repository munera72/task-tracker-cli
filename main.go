package main

import (
	"task-tracker-cli/enums"
	"task-tracker-cli/structs"
	"time"
)

func main() {
	myTaskList := structs.NewTaskList()

	myTaskList.AddTask(*structs.NewTask("...", enums.Todo, time.Now()))
	myTaskList.AddTask(*structs.NewTask("Hello wor ", enums.InProgress, time.Now()))
	myTaskList.AddTask(*structs.NewTask("Hello world :p", enums.Done, time.Now()))

	myTaskList.UpdateTaskDescription(1, "Hel")
	myTaskList.UpdateTaskStatus(1, enums.InProgress)

	myTaskList.All()

}
