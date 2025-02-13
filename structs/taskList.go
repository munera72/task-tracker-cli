package structs

import (
	"fmt"
	"task-tracker-cli/enums"
	"time"
)

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func NewTaskList() *TaskList {
	return &TaskList{}
}

func (tl *TaskList) AddTask(task Task) {
	tl.Tasks = append(tl.Tasks, task)
}

func (tl *TaskList) FindTask(id int) (int, *Task) {
	var index int
	var taskPtr *Task
	for i := range tl.Tasks {
		if tl.Tasks[i].Id == id {
			taskPtr = &(tl.Tasks[i])
			index = i
			break
		}
	}
	return index, taskPtr
}

func (tl *TaskList) UpdateTaskDescription(id int, desc string) {
	_, task := tl.FindTask(id)
	task.SetDescription(desc)
	task.SetUpdatedAt(time.Now())
}

func (tl *TaskList) UpdateTaskStatus(id int, st enums.TaskState) {
	_, task := tl.FindTask(id)
	task.SetStatus(st)
	task.SetUpdatedAt(time.Now())
}

func (tl TaskList) All() {
	for _, task := range tl.Tasks {
		fmt.Printf("%v\n", task)
	}
}

func (tl TaskList) FilterByState(state enums.TaskState) {
	for _, task := range tl.Tasks {
		if task.Status == state {
			fmt.Printf("%v\n", task)
		}
	}
}

func (tl *TaskList) DeleteTask(id int) {
	i, _ := tl.FindTask(id)
	tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
}
