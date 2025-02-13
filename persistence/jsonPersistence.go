package persistence

import (
	"encoding/json"
	"errors"
	"os"
)

var fileName string = "tasks.json"
var file *os.File
var encoder *json.Encoder

func SetUpPersistence() {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		createPersistenceFile()
	}
	createEncoder()
}

func createPersistenceFile() {
	var err error
	file, err = os.Create(fileName)
	if err != nil {
		panic(err)
	}
}

func createEncoder() {
	encoder = json.NewEncoder(file)
}

func WriteInPersistence(content any) {
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
