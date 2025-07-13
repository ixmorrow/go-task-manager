package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

// Array of type Task named tasks
var tasks []Task
var filePath = "tasks.json"

func loadTasks() {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return // Start empty if no file
	}
	json.Unmarshal(data, &tasks)
}

func saveTasks() {
	data, _ := json.Marshal(tasks)
	ioutil.WriteFile(filePath, data, 0644)
}

func addTask(description string) {
	id := len(tasks) + 1
	tasks = append(tasks, Task{ID: id, Description: description, Completed: false})
	saveTasks()
}

func getTask(task_id string) {
	num, err := strconv.Atoi((task_id))
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}
	for _, t := range tasks {
		if t.ID == num {
			fmt.Printf("ID: %d, Description: %s, Completed: %s\n", t.ID, t.Description, _getStatus(t))
		}
	}
}

// Internal use only
func _getStatus(t Task) string {
	status := "Pending"
	if t.Completed {
		status = "Completed"
	}
	return status
}

func listTasks() {
	for _, t := range tasks {
		status := "Pending"
		if t.Completed {
			status = "Completed"
		}
		fmt.Printf("Id: %d, Description: %s, Completed: %s\n", t.ID, t.Description, status)
	}
}

// Add more functions like completeTask, deleteTask...

func main() {
	loadTasks()
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-manager [add|list|complete|delete] [args]")
		return
	}

	switch os.Args[1] {
	case "add":
		addTask(os.Args[2])
	case "list":
		listTasks()
	case "get_task":
		getTask(os.Args[2])
	}
}
