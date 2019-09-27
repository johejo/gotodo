package main

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/johejo/gotodo/entity"
)

type ToDoRepository interface {
	Get(ctx context.Context, todoID string) (*entity.ToDo, error)
}
type toDoRepository struct {
	db *sqlx.DB
}

func NewToDoRepository(db *sql.DB) ToDoRepository {
	return &toDoRepository{db: sqlx.NewDb(db, "mysql")}
}

func (r *toDoRepository) Get(ctx context.Context, todoID string) (*entity.ToDo, error) {
	const q = "SELECT * FROM todo WHERE todo_id=?"
	var t entity.ToDo
	if err := r.db.GetContext(ctx, &t, q, todoID); err != nil {
		return nil, err
	}
	return &t, nil
}
