package main

import (
	"os"
	"task-tracker-cli/controller"
	"task-tracker-cli/persistence"
)

func main() {
	persistence.SetUpPersistence()
	if len(os.Args) >= 2 {
		controller.ProcessArgs(os.Args)
	} else {
		panic("Please provide an action to perform")
	}
}
