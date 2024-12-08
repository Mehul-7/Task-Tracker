package util

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type task struct {
	Id          uuid.UUID
	Description string
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New(description string) task {
	taskObj := task{uuid.New(), description, 0, time.Now(), time.Now()}
	//Write to storage
	return taskObj
}

func AddTask(task task){
	existingTasks, getErr := getFileContent()
	if getErr != nil{
		fmt.Println("Unable to get file content")
		return
	}

	existingTasks[task.Id] = task

	write(existingTasks)
}

func UpdateTask(id string, newVal any, field string) (bool, error){
	tasks, err := getFileContent()
	if err != nil {
		log.Fatal("Failed to retrieve tasks from storage!")
	}

	uid, parseErr := uuid.Parse(id)
	if parseErr != nil {
		log.Fatal("Invalid uid!")
	}

	taskObj := reflect.ValueOf(tasks[uid]).Elem()

	targetField := taskObj.FieldByName(field)

	if !targetField.IsValid() {
		return false, fmt.Errorf("filed %s does not exist", field)
	}

	if !targetField.CanSet() {
		return false, fmt.Errorf("cannot modify field %s", field)
	}

	if targetField.Type() == reflect.TypeOf("") {
		targetField.SetString(newVal.(string))
	} else if targetField.Type() == reflect.TypeOf(1) {
		newVal, isInt := newVal.(int)
		if !isInt {
			log.Fatal("Invalid argument value for the selected field")
		}
		targetField.SetInt(int64(newVal))
	} else {
		log.Fatal("cannot modify field : ", field)
	}

	fmt.Println(taskObj.)

	return true, nil
}
