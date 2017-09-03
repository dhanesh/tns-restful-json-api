package main

import "time"

/*Todo data structure for holding todo*/
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

/*Todos collection of Todo data*/
type Todos []Todo
