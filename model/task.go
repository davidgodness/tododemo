package model

import (
	"context"
	"github.com/davidgodeness/tododemo/database"
	"time"
)

const (
	Active uint8 = iota
	Done
)

type Task struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	Status    uint8     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateTask(ctx context.Context, name string) (Task, error) {
	t := Task{
		Name:      name,
		Status:    Active,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db, err := database.GetDb(ctx)
	if err != nil {
		return t, err
	}

	r, err := db.ExecContext(ctx, "insert into task (name, status, created_at, updated_at) values (?, ?, ?, ?)",
		t.Name, t.Status, t.CreatedAt, t.UpdatedAt)
	if err != nil {
		return t, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return t, err
	}

	t.Id = uint64(id)
	return t, nil
}

func GetTasks(ctx context.Context) ([]Task, error) {
	db, err := database.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	r, err := db.QueryContext(ctx, "select id, name, status, created_at, updated_at from task")
	if err != nil {
		return nil, err
	}

	tasks := make([]Task, 0)
	for r.Next() {
		var t Task
		if err := r.Scan(&t.Id, &t.Name, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}
