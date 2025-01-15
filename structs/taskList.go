package structs

import (
	"fmt"
	"task-tracker-cli/enums"
)

type TaskList struct {
	lastAddedId int
	tasks       []Task
}

func NewTaskList() *TaskList {
	return &TaskList{
		0,
		make([]Task, 0, 10),
	}
}

func (tl *TaskList) AddTask(task Task) {
	tl.lastAddedId++
	task.id = tl.lastAddedId
	tl.tasks = append(tl.tasks, task)
}

func (tl *TaskList) GetLastAddedId() int {
	return tl.lastAddedId
}

func (tl *TaskList) All() {
	for _, task := range tl.tasks {
		fmt.Printf("id: %v\n"+
			"desctiption: %s\n"+
			"status: %s\n"+
			"created at: %s\n\n\n",
			task.GetId(), task.GetDescription(), task.GetStatus().String(), task.GetCreatedAt())
	}
}

func (tl *TaskList) AllDone() {
	for _, task := range tl.tasks {
		if task.status == enums.Done {
			fmt.Printf("id: %v\n"+
				"desctiption: %s\n"+
				"status: %s\n"+
				"created at: %s\n\n\n",
				task.GetId(), task.GetDescription(), task.GetStatus().String(), task.GetCreatedAt())
		}
	}
}
