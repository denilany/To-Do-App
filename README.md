# ToDo CLI Application

This is a simple command-line application written in Go for managing a to-do list. The application allows users to add, delete, edit, toggle the completion status of tasks, and list all tasks. The to-do data is persisted (stored) in a JSON file.

## Features

- **Add a Task**: Add a new task with a title.
- **Delete a Task**: Remove a task by its index.
- **Edit a Task**: Edit the title of a task by its index.
- **Toggle Task Completion**: Mark a task as completed or uncompleted.
- **List All Tasks**: Display all tasks with their status.

## Project Structure

- `main.go`: Entry point of the application that loads, processes, and saves tasks.
- `todo/todo.go`: Handles to-do item logic (adding, deleting, editing, toggling, printing tasks).
- `storage/storage.go`: Handles saving and loading to-do data to/from a JSON file.
- `args/args.go`: Handles command-line argument parsing and execution logic.

## How It Works

1. The program uses command-line flags to specify operations (`--add`, `--edit`, `--del`, `--toggle`, `--list`).
2. To-do items are stored in a `Todos` slice, with each item represented by a `Todo` struct containing:
   - `Title`: The task title.
   - `Completed`: A boolean indicating whether the task is complete.
   - `CreatedAt`: Timestamp when the task was created.
   - `CompletedAt`: Timestamp when the task was marked as completed.
3. The list of tasks is stored in a `todos.json` file using JSON format for future reference.

## Commands

Here are the available commands and how to use them:

### Add a Task

```bash
./main --add "Task Title"
```

### Delete a Task

```bash
./main --del 2
```

Where `2` is the index of the task you want to delete.

### Edit a Task

```bash
./main --edit "2:New Task Title"
```

Where `2` is the index of the task you want to edit, followed by the new title.

### Toggle Task Completion

```bash
./main --toggle 2
```

Where `2` is the index of the task to toggle.

### List All Tasks

```bash
./main --list
```

This command will display all tasks, showing their index, title, completion status, and timestamps.

## Data Storage

Tasks are stored in a file named `todos.json` in the same directory as the executable. This file will be loaded on application startup and updated on application exit.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/denilany/To-Do-App.git
   cd To-Do-App
   ```

2. Build the project:

   ```bash
   go build -o main
   ```

3. Run the program with the desired flags.

## Dependencies

- **Go**: Make sure you have Go installed on your system.
- **Table Package**: For rendering the to-do list in a table format. Install it using:

  ```bash
  go get github.com/aquasecurity/table
  ```

## Future Improvements

- Add more comprehensive error handling.
- Add the ability to prioritize tasks.
- Implement task categorization (e.g., work, personal).
- Enhance the user interface for a more detailed display of task information.

## License

This project is licensed under the [MIT License](./LICENSE).
