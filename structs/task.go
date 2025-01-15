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

func NewTask(id int, description string, status enums.TaskState, createdAT, updatedAt time.Time) *Task {
	return &Task{
		id,
		description,
		status,
		createdAT,
		updatedAt,
	}
}
