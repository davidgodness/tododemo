package model

import (
	"context"
	"database/sql"
	"github.com/davidgodeness/tododemo/database"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	Active uint8 = iota
	Done
)

type Task struct {
	Id        int64     `json:"id"`
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

	t.Id = id
	return t, nil
}

func GetTasks(ctx context.Context, query gin.H) ([]Task, error) {
	db, err := database.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	var r *sql.Rows

	if status, ok := query["status"]; ok {
		r, err = db.QueryContext(ctx,
			"select id, name, status, created_at, updated_at from task where status = ?", status)
		if err != nil {
			return nil, err
		}
	} else {
		r, err = db.QueryContext(ctx, "select id, name, status, created_at, updated_at from task")
		if err != nil {
			return nil, err
		}
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

func UpdateTask(ctx context.Context, id int64, data gin.H) (int64, error) {
	db, err := database.GetDb(ctx)
	if err != nil {
		return 0, err
	}

	if status, ok := data["status"]; ok {
		r, err := db.ExecContext(ctx, "update task set status = ? where id = ?", status, id)
		if err != nil {
			return 0, err
		}
		return r.RowsAffected()
	}

	return 0, nil
}

func DeleteTask(ctx context.Context, id int64) (int64, error) {
	db, err := database.GetDb(ctx)
	if err != nil {
		return 0, err
	}

	r, err := db.ExecContext(ctx, "delete from task where id = ?", id)
	if err != nil {
		return 0, err
	}

	return r.RowsAffected()
}

func GetTask(ctx context.Context, id int64) (Task, error) {
	var t Task
	db, err := database.GetDb(ctx)
	if err != nil {
		return t, err
	}

	r := db.QueryRowContext(ctx, "select id, name, status, created_at, updated_at from task where id = ?", id)

	err = r.Scan(&t.Id, &t.Name, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return t, err
	}

	return t, nil
}
