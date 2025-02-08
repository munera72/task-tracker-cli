package main

import (
	"encoding/json"
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

	//Create the json file
	file, err := os.Create("tasks.json")
	if err != nil {
		panic(err)
	}

	//Create the encoder for the json to write in the file we just created
	encoder := json.NewEncoder(file)
	err = encoder.Encode(myTaskList)
	if err != nil {
		panic(err)
	}

	//Read the json file and put the content on a new taskList
	content, err := os.ReadFile("tasks.json")
	if err != nil {
		panic(err)
	}

	newTasksList := structs.NewTaskList()
	err = json.Unmarshal(content, newTasksList)
	if err != nil {
		panic(err)
	}

	newTasksList.All()

}
