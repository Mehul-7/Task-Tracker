package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"github.com/google/uuid"
)

var fileName = "storage.json"
func init(){
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, creationErr := os.Create(fileName)
		if creationErr != nil {
			log.Fatal(creationErr)
		}
		defer file.Close()
		fmt.Println("File created successfully")
	}
}

func getFileContent() (map[uuid.UUID]task, error) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil{
		fmt.Println("Unable to read file content!")
		return nil, err
	}

	existingTasks := make(map[uuid.UUID]task)
	if len(fileContent) > 0 {
		err = json.Unmarshal(fileContent, &existingTasks)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	return existingTasks, nil
}

func write(tasks map[uuid.UUID]task) (bool, error){
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON: ", err)
		return false, err
	}

	writeErr := os.WriteFile("storage.json", jsonData, 0644)
	if writeErr != nil {
		fmt.Println("ERror writing file: ", writeErr)
		return false, err
	}

	return true, nil
}