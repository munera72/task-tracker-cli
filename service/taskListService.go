package service

import (
	"task-tracker-cli/enums"
	"task-tracker-cli/repo"
	"task-tracker-cli/structs"
	"time"
)

var taskList structs.TaskList = repo.ReadTaskList()

func ListTasks(status string) {
	if status == "" {
		taskList.All()
	} else {
		taskList.FilterByState(enums.GetStatus(status))
	}
}

func AddTask(desc string) {
	taskList.AddTask(*structs.NewTask(desc, enums.Todo, time.Now()))
	repo.SaveTaskList(taskList)
}

func DeleteTask(id int) {
	taskList.DeleteTask(id)
	repo.SaveTaskList(taskList)
}

func UpdateTaskDescription(id int, desc string) {
	taskList.UpdateTaskDescription(id, desc)
	repo.SaveTaskList(taskList)
}

func UpdateTaskStatus(id int, status enums.TaskState) {
	taskList.UpdateTaskStatus(id, status)
	repo.SaveTaskList(taskList)
}
