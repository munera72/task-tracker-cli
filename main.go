package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	myJson, err := json.MarshalIndent(myTaskList, "", " ")

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Print(string(myJson))

}
