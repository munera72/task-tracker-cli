package structs

import "iter"

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

func (tl *TaskList) All() iter.Seq[Task] {
	return func(yield func(Task) bool) {
		for i := 0; i < len(tl.tasks); i++ {
			if !yield(tl.tasks[i]) {
				return
			}
		}
	}
}
