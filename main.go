package main

import (
	"ToDo/args"
	"ToDo/storage"
	"ToDo/todo"
)

func main() {
	todos := todo.Todos{}
	storage := storage.NewStorage[todo.Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := args.NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)
}
