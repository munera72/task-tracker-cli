package main

import (
	"os"
	"task-tracker-cli/controller"
)

func main() {
	if len(os.Args) >= 2 {
		controller.ProcessArgs(os.Args)
	} else {
		panic("Please provide an action to perform")
	}
}
