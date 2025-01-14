package structs

import (
	"task-tracker-cli/enums"
	"time"
)

type Task struct {
	id          int
	description string
	status      enums.TaskState
	createdAt   time.Time
	updatedAt   time.Time
}

type TaskList struct {
	tasks []Task
}
