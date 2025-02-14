package persistence

import (
	"encoding/json"
	"fmt"
	"os"
	"task-tracker-cli/structs"
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

// func createPersistenceFile() *os.File {
// 	file, err := os.Create(fileName)
// 	fmt.Print(file)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return file
// }

func createEncoder() *json.Encoder {
	return json.NewEncoder(file)
}

func WriteInPersistence(content structs.TaskList) {
	fmt.Print(encoder)
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
