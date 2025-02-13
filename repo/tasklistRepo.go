package repo

import (
	"encoding/json"
	"task-tracker-cli/persistence"
	"task-tracker-cli/structs"
)

func SaveTaskList(tl structs.TaskList) {
	persistence.WipePersistenceContent()
	persistence.WriteInPersistence(tl)
}

func DeleteTaskList() {
	persistence.WipePersistenceContent()
}

func ReadTaskList() structs.TaskList {
	var tl structs.TaskList
	content, _ := persistence.ReadPersistence()
	json.Unmarshal(content, &tl)
	return tl
}
