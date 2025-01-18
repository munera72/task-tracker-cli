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

func (tl *TaskList) FindTask(id int) *Task {
	var taskPtr *Task
	for i := range tl.tasks {
		if tl.tasks[i].Id() == id {
			taskPtr = &(tl.tasks[i])
			break
		}
	}
	return taskPtr
}

func (tl *TaskList) UpdateTaskDescription(id int, desc string) {
	task := tl.FindTask(id)
	(*task).SetDescription(desc)
}

func (tl *TaskList) UpdateTaskStatus(id int, st enums.TaskState) {
	task := tl.FindTask(id)
	(*task).SetStatus(st)
}

func (tl TaskList) All() {
	for _, task := range tl.tasks {
		fmt.Printf("%v\n", task)
	}
}

func (tl TaskList) AllDone() {
	for _, task := range tl.tasks {
		if task.status == enums.Done {
			fmt.Print(task)
		}
	}
}

func (tl TaskList) AllInProgress() {
	for _, task := range tl.tasks {
		if task.status == enums.InProgress {
			fmt.Print(task)
		}
	}
}

func (tl *TaskList) AllTodo() {
	for _, task := range tl.tasks {
		if task.status == enums.Todo {
			fmt.Print(task)
		}
	}
}
