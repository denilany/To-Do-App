package main

import "time"

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title: title,
		Completed: false,
		CompletedAt: nil,
		CreatedAt: time.Now(),
	}

	*todos = append(*todos, todo)
}
