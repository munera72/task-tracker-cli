package structs

import (
	"fmt"
	"task-tracker-cli/enums"
	"time"
)

var currentIdCount int = 0

type Task struct {
	id          int
	description string
	status      enums.TaskState
	createdAt   time.Time
	updatedAt   time.Time
}

func NewTask(description string, status enums.TaskState, createdAt time.Time) *Task {
	currentIdCount++
	return &Task{
		id:          currentIdCount,
		description: description,
		status:      status,
		createdAt:   createdAt,
	}
}

func (t *Task) SetId(id int) {
	t.id = id
}

func (t Task) Id() int {
	return t.id
}

func (t *Task) SetDescription(d string) {
	t.description = d
}

func (t Task) Description() string {
	return t.description
}

func (t Task) Status() enums.TaskState {
	return t.status
}

func (t *Task) SetStatus(newStatus enums.TaskState) {
	t.status = newStatus
}

func (t Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t Task) UpdatedAt() time.Time {
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
		t.Id(), t.Description(), t.Status().String(), t.CreatedAt())
}
