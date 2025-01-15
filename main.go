package main

import (
	"fmt"
	"task-tracker-cli/enums"
	"task-tracker-cli/structs"
	"time"
)

func main() {
	myTaskList := structs.NewTaskList()
	myTask := structs.NewTask("Hello world :p", enums.Todo, time.Now())

	myTaskList.AddTask(*myTask)

	for task := range myTaskList.All() {
		fmt.Printf("id: %v\n"+
			"desctiption: %s\n"+
			"status: %s\n",
			task.GetId(), task.GetDescription(), task.GetStatus().String())
	}
}
