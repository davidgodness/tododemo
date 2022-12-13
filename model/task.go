package model

import "time"

const (
	Active uint8 = iota
	Done
)

type Task struct {
	Id        uint64    `json:"id"xml:"id"`
	Name      string    `json:"name"xml:"name"binding:"required"`
	Status    uint8     `json:"status"xml:"status"`
	CreatedAt time.Time `json:"created_at"xml:"created_at"`
	UpdatedAt time.Time `json:"updated_at"xml:"updated_at"`
}

var tasks []Task

func CreateTask(t Task) error {
	tasks = append(tasks, t)
	return nil
}

func GetTasks() ([]Task, error) {
	return tasks, nil
}
