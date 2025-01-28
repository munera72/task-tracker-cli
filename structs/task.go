package structs

import (
	"fmt"
	"task-tracker-cli/enums"
	"time"
)

var currentIdCount int = 0

type Task struct {
	Id          int             `json:"id"`
	Description string          `json:"description"`
	Status      enums.TaskState `json:"status"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}

func NewTask(description string, status enums.TaskState, createdAt time.Time) *Task {
	currentIdCount++
	return &Task{
		Id:          currentIdCount,
		Description: description,
		Status:      status,
		CreatedAt:   createdAt,
	}
}

func (t *Task) SetId(id int) {
	t.Id = id
}

func (t *Task) SetDescription(d string) {
	t.Description = d
}

func (t *Task) SetStatus(newStatus enums.TaskState) {
	t.Status = newStatus
}

func (t *Task) SetUpdatedAt(updatedAt time.Time) {
	t.UpdatedAt = updatedAt
}

func (t Task) String() string {
	return fmt.Sprintf("id: %v\n"+
		"desctiption: %s\n"+
		"status: %s\n"+
		"created at: %s\n\n\n",
		t.Id, t.Description, t.Status.String(), t.CreatedAt)
}
