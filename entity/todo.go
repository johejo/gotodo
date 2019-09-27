package entity

import "time"

type ToDo struct {
	ToDoID    string    `json:"toDoId" db:"todo_id"`
	ToDo      string    `json:"toDo" db:"todo"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
