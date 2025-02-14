package persistence

import (
	"encoding/json"
	"os"
	"task-cli/structs"
)

var fileName string = "tasks.json"
var file *os.File
var encoder *json.Encoder

func SetUpPersistence() {
	var err error
	file, err = os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	encoder = createEncoder()
}

func createEncoder() *json.Encoder {
	return json.NewEncoder(file)
}

func WriteInPersistence(content structs.TaskList) {
	if encoder == nil {
		createEncoder()
	}

	err := encoder.Encode(content)
	if err != nil {
		panic(err)
	}
}

func ReadPersistence() ([]byte, error) {
	return os.ReadFile(fileName)
}

func WipePersistenceContent() {
	os.Truncate(fileName, 0)
}

func ClosePersistence() {
	file.Close()
}
