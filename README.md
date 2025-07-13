# Go Task Manager

A command-line task management application written in Go that provides full CRUD operations for tasks with persistent storage using JSON.

## Features

- Add new tasks with descriptions
- List all tasks with their status (Pending/Completed)
- Get details of a specific task by ID
- Mark tasks as completed
- Delete tasks
- Persistent storage using JSON file
- Efficient O(1) task lookups using map-based storage
- Auto-incrementing task IDs

## Installation

1. Clone the repository:
```bash
git clone https://github.com/ixmorrow/go-task-manager.git
cd go-task-manager
```

2. Build the application:
```bash
go build -o task-manager main.go
```

## Usage

```bash
./task-manager [command] [arguments]
```

### Commands

- `add [description]` - Add a new task with the given description
- `list` - Display all tasks with their current status
- `get_task [id]` - Get details of a specific task by ID
- `complete [id]` - Mark a task as completed by ID
- `delete [id]` - Delete a task by ID

### Examples

```bash
# Add a new task
./task-manager add "Complete Go project"

# List all tasks
./task-manager list

# Get details of task with ID 2
./task-manager get_task 2

# Mark task 2 as completed
./task-manager complete 2

# Delete task 3
./task-manager delete 3
```

## Data Storage

Tasks are stored in `tasks.json` file in the same directory as the executable. The file is automatically created when you add your first task. Tasks are internally managed using a map for efficient O(1) lookups but persisted as a JSON array.

## Project Structure

- `main.go` - Main application code containing task management logic
- `go.mod` - Go module definition
- `tasks.json` - JSON file for persistent task storage
- `task-manager` - Compiled executable

## Task Structure

Each task contains:
- `ID` - Unique identifier (auto-incremented)
- `Description` - Task description
- `Completed` - Boolean status indicating if the task is completed

## Implementation Details

- Uses a map (`map[int]Task`) for efficient task storage and retrieval
- Automatic ID management with `nextID` tracking
- Error handling for invalid task IDs
- JSON marshaling/unmarshaling for persistence

## Requirements

- Go 1.24.5 or higher