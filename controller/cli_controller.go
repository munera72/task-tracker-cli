package controller

import (
	"fmt"
)

func ProcessArgs(args []string) {
	switch args[1] {
	case "add":
		fmt.Print("adding")
	case "list":
		fmt.Print("listing")
	case "delete":
		fmt.Print("deleting")
	case "update":
		fmt.Print("updating")
	case "mark-in-progress":
		fmt.Print("marking in progress")
	case "mark-done":
		fmt.Print("marking done")
	}
}
