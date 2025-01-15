package structs

import (
	"fmt"
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

func NewTask(description string, status enums.TaskState, createdAt time.Time) *Task {
	return &Task{
		description: description,
		status:      status,
		createdAt:   createdAt,
	}
}

func (t *Task) SetId(id int) {
	t.id = id
}

func (t *Task) GetId() int {
	return t.id
}

func (t *Task) SetDescription(d string) {
	t.description = d
}

func (t *Task) GetDescription() string {
	return t.description
}

func (t *Task) GetStatus() enums.TaskState {
	return t.status
}

func (t *Task) SetStatus(newStatus enums.TaskState) {
	t.status = newStatus
}

func (t *Task) GetCreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) GetUpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Task) SetUpdatedAt(updatedAt time.Time) {
	t.updatedAt = updatedAt
}

func (t Task) String() string {
	return fmt.Sprintf("id: %v\n"+
		"desctiption: %s\n"+
		"status: %s\n"+
		"created at: %s\n\n\n",
		t.GetId(), t.GetDescription(), t.GetStatus().String(), t.GetCreatedAt())
}
