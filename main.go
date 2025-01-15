package main

import (
	"task-tracker-cli/enums"
	"task-tracker-cli/structs"
	"time"
)

func main() {
	myTaskList := structs.NewTaskList()
	myTask1 := structs.NewTask("...", enums.Todo, time.Now())
	myTask2 := structs.NewTask("Hello wor ", enums.InProgress, time.Now())
	myTask3 := structs.NewTask("Hello world :p", enums.Done, time.Now())

	myTaskList.AddTask(*myTask1)
	myTaskList.AddTask(*myTask2)
	myTaskList.AddTask(*myTask3)

	myTaskList.AllDone()
	myTaskList.All()

}
