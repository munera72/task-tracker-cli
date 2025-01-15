package structs

import (
	"iter"
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

type TaskList struct {
	lastAddedId int
	tasks       []Task
}

func NewTaskList() *TaskList {
	return &TaskList{
		1,
		make([]Task, 0, 10),
	}
}

func (tl *TaskList) AddTask(task Task) {
	tl.tasks = append(tl.tasks, task)
}

func (tl *TaskList) GetLastAddedId() int {
	return tl.lastAddedId
}

func (tl *TaskList) All() iter.Seq[Task] {
	return func(yield func(Task) bool) {
		for i := 0; i < len(tl.tasks); i++ {
			if !yield(tl.tasks[i]) {
				return
			}
		}
	}
}
