package util

import (
	"fmt"
	"log"
	// "reflect"
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

func UpdateTask(id string, newVal string) (bool, error){
	tasks, err := getFileContent()
	if err != nil {
		log.Fatal("Failed to retrieve tasks from storage!")
	}

	uid, parseErr := uuid.Parse(id)
	if parseErr != nil {
		log.Fatal("Invalid uid!")
	}

	targetTask, exists := tasks[uid]
	if !exists {
		log.Fatal("Task not found")
	}

	targetTask.Description = newVal

	tasks[uid] = targetTask

	write(tasks)
	//logic for dynamically selecting and manipulating fields 
	// taskObj := reflect.ValueOf(&targetTask).Elem()

	// targetField := taskObj.FieldByName(field)

	// if !targetField.IsValid() {
	// 	return false, fmt.Errorf("field %s does not exist", field)
	// }

	// if !targetField.CanSet() {
	// 	return false, fmt.Errorf("cannot modify field %s", field)
	// }

	// if targetField.Type() == reflect.TypeOf("") {
	// 	targetField.SetString(newVal.(string))
	// 	targetTask.UpdatedAt = time.Now()
	// } else if targetField.Type() == reflect.TypeOf(1) {
	// 	newVal, isInt := newVal.(int)
	// 	if !isInt {
	// 		log.Fatal("Invalid argument value for the selected field")
	// 	}
	// 	targetField.SetInt(int64(newVal))
	// 	targetTask.UpdatedAt = time.Now()
	// } else {
	// 	log.Fatal("cannot modify field : ", field)
	// }

	// tasks[uid] = targetTask

	return true, nil
}

func ListTasks(status string){
	tasks, err := getFileContent()
	
	if err != nil {
		log.Fatal("Failed to retrieve tasks from storage!")
	} 

	statusVal := -1
	switch status {
		case "todo":
			statusVal = 0
		case "in-progress":
			statusVal = 1
		case "done":
			statusVal = 2 
	}

	for uid, task := range tasks {
		if task.Status == statusVal || statusVal == -1 {
			fmt.Printf("UUID: %v, Task: %+v\n", uid, task.Description)
		}
	}
}

func DeleteTask(id string){
	taskList, err := getFileContent()
	if err != nil {
		log.Fatal("Failed to retrieve tasks from storage")
	}

	uid, parseErr := uuid.Parse(id)
	if parseErr != nil {
		log.Fatal("Invalid uid!")
	}

	_, exists := taskList[uid]
	if !exists {
		log.Fatal("Task not found")
	}

	delete(taskList, uid)

	if _, exists := taskList[uid]; !exists {
		fmt.Println("Task was successfully deleted")
	}

	write(taskList)

}

func UpdateStatus(id string, status string) {
	tasks, err := getFileContent()
	
	if err != nil {
		log.Fatal("Failed to retrieve tasks from storage!")
	} 

	uid, parseErr := uuid.Parse(id)
	if parseErr != nil {
		log.Fatal("Invalid uid!")
	}

	targetTask, exists := tasks[uid]
	if !exists {
		log.Fatal("Task not found")
	}

	switch status {
	case "in-progress":
		targetTask.Status = 1
	case "done":
		targetTask.Status = 2
	}

	tasks[uid] = targetTask

	write(tasks)

}
