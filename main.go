package main

import (
	"os"
	"task-cli/controller"
	"task-cli/persistence"
)

func main() {
	persistence.SetUpPersistence()
	defer persistence.ClosePersistence()
	if len(os.Args) >= 2 {
		controller.ProcessArgs(os.Args)
	} else {
		panic("Please provide an action to perform")
	}
}
