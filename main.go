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
var tasks = make(map[int]Task) // Map for O(1) lookups by ID
var nextID = 1                 // Track next available ID (load/max from file if needed)
var filePath = "tasks.json"

func loadTasks() {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return // Start empty if no file
	}
	var loadedTasks []Task // Temporarily load into slice from JSON
	if err := json.Unmarshal(data, &loadedTasks); err != nil {
		return
	}
	for _, t := range loadedTasks {
		tasks[t.ID] = t
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}
	json.Unmarshal(data, &tasks)
}

func saveTasks() {
	var taskList []Task // Convert map to slice for JSON persistence
	for _, t := range tasks {
		taskList = append(taskList, t)
	}
	data, _ := json.Marshal(taskList)
	ioutil.WriteFile(filePath, data, 0644)
}

func addTask(description string) {
	id := len(tasks) + 1
	tasks[nextID] = Task{ID: id, Description: description, Completed: false}
	nextID++
	saveTasks()
}

func getTask(taskId string) {
	id, err := strconv.Atoi(taskId)
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}
	t := tasks[id]
	fmt.Printf("ID: %d, Description: %s, Completed: %s\n", t.ID, t.Description, _getStatus(t))
}

// Internal use only
func _getStatus(t Task) string {
	status := "Pending"
	if t.Completed {
		status = "Completed"
	}
	return status
}

func completeTask(taskId string) {
	id, err := strconv.Atoi(taskId)
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
		return
	}
	if t, ok := tasks[id]; ok {
		t.Completed = true
		tasks[id] = t // Maps are reference types, but struct update requires re-assignment
		saveTasks()
		fmt.Println("Task", id, "marked as completed.")
	} else {
		fmt.Println("Task not found: ", id)
	}
}

func deleteTask(taskId string) {
	id, err := strconv.Atoi(taskId)
	if err != nil {
		log.Fatalf("Error converting string to in: %v", err)
		return
	}
	if _, ok := tasks[id]; ok {
		delete(tasks, id)
		saveTasks()
		fmt.Println("Task", id, "deleted.")
	} else {
		fmt.Println("Task not found.")
	}
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

func main() {
	loadTasks()
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-manager [add|list|complete|delete] [args]")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: add <description>")
			return
		}
		addTask(os.Args[2])
	case "list":
		listTasks()
	case "get_task":
		if len(os.Args) < 3 {
			fmt.Println("Usage: get_task <id>")
			return
		}
		getTask(os.Args[2])
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: complete <id>")
			return
		}
		completeTask(os.Args[2])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: delete <id>")
			return
		}
		deleteTask(os.Args[2])
	}
}
