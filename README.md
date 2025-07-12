# Go Task Manager

A simple command-line task management application written in Go that allows you to add and list tasks with persistent storage.

## Features

- Add new tasks with descriptions
- List all tasks with their status (Pending/Completed)
- Persistent storage using JSON file
- Simple command-line interface

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

### Examples

```bash
# Add a new task
./task-manager add "Complete Go project"

# List all tasks
./task-manager list
```

## Data Storage

Tasks are stored in `tasks.json` file in the same directory as the executable. The file is automatically created when you add your first task.

## Project Structure

- `main.go` - Main application code containing task management logic
- `go.mod` - Go module definition
- `tasks.json` - JSON file for persistent task storage
- `task-manager` - Compiled executable

## Task Structure

Each task contains:
- `ID` - Unique identifier (auto-incremented)
- `Description` - Task description
- `Completed` - Boolean status (currently always false in this version)

## Note

This is a basic implementation with add and list functionality. Additional features like completing and deleting tasks are mentioned in the code comments but not yet implemented.

## Requirements

- Go 1.24.5 or higher