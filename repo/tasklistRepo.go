package repo

import (
	"encoding/json"
	"task-tracker-cli/persistence"
	"task-tracker-cli/structs"
)

func Save(tl structs.TaskList) {
	persistence.WriteInPersistence(tl)
}

func Delete() {
	persistence.WipePersistenceContent()
}

func Read() (tl *structs.TaskList) {
	content, _ := persistence.ReadPersistence()
	json.Unmarshal(content, tl)
	return
}
