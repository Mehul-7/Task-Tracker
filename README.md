# Task Tracker CLI

# Project
Project Roadmap: https://roadmap.sh/projects/task-tracker

## Overview
Task Tracker is a simple, lightweight command-line interface (CLI) application for managing your tasks efficiently. Store, track, and organize your tasks directly from the terminal.

## Features
- Add new tasks
- Update existing tasks
- Delete tasks
- Mark tasks as in progress
- Mark tasks as done
- List tasks with various filters

## Installation

### Prerequisites
- Go 1.16+ installed
- Git (optional)

### Download
Download the latest release from the [Releases](link-to-github-releases) page for your operating system.

### Build from Source
```bash
# Clone the repository
git clone https://github.com/yourusername/task-tracker.git

# Navigate to the project directory
cd task-tracker

# Build the application
go build -o task-cli
```

## Usage

### Adding a New Task
```bash
# Add a task with a description
task-cli add "Buy groceries"
# Expected Output: Task added successfully (ID: 1)
```

### Updating a Task
```bash
# Update a task by its ID
task-cli update 1 "Buy groceries and cook dinner"
```

### Deleting a Task
```bash
# Delete a task by its ID
task-cli delete 1
```

### Marking Task Status
```bash
# Mark a task as in progress
task-cli mark-in-progress 1

# Mark a task as done
task-cli mark-done 1
```

### Listing Tasks

#### List All Tasks
```bash
task-cli list
```

#### List Tasks by Status
```bash
# List completed tasks
task-cli list done

# List tasks to do
task-cli list todo

# List tasks in progress
task-cli list in-progress
```

## Task Statuses
- `todo`: Default status for new tasks
- `in-progress`: Task is currently being worked on
- `done`: Task is completed

## Storage
- Tasks are stored in a local JSON file
- Persistent storage across sessions
- Easy to backup and manage

## Examples

```bash
# Add a new task
$ task-cli add "Prepare presentation"
Task added successfully (ID: 1)

# Mark task as in progress
$ task-cli mark-in-progress 1

# List tasks in progress
$ task-cli list in-progress
```

## Error Handling
- Clear error messages for invalid commands
- Helpful usage instructions

## Contributing
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License
[Choose an appropriate license, e.g., MIT]

## Contact
[Your Name or Contact Information]
[Project Repository Link]
