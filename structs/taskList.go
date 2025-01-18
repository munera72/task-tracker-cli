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
	return &TaskList{}
}

func (tl *TaskList) AddTask(task Task) {
	tl.tasks = append(tl.tasks, task)
}

func (tl TaskList) LastAddedId() int {
	return tl.lastAddedId
}

func (tl *TaskList) FindTask(id int) (int, *Task) {
	var index int
	var taskPtr *Task
	for i := range tl.tasks {
		if tl.tasks[i].Id() == id {
			taskPtr = &(tl.tasks[i])
			index = i
			break
		}
	}
	return index, taskPtr
}

func (tl *TaskList) UpdateTaskDescription(id int, desc string) {
	_, task := tl.FindTask(id)
	(*task).SetDescription(desc)
}

func (tl *TaskList) UpdateTaskStatus(id int, st enums.TaskState) {
	_, task := tl.FindTask(id)
	(*task).SetStatus(st)
}

func (tl TaskList) All() {
	for _, task := range tl.tasks {
		fmt.Printf("%v\n", task)
	}
}

func (tl TaskList) FilterByState(state enums.TaskState) {
	for _, task := range tl.tasks {
		if task.status == state {
			fmt.Printf("%v\n", task)
		}
	}
}

func (tl *TaskList) DeleteTask(id int) {
	i, _ := tl.FindTask(id)
	tl.tasks = append(tl.tasks[:i], tl.tasks[i+1:]...)
}
