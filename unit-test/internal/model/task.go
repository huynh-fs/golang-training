package model

import "fmt"

type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

func NewTask(id int, title, description string) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("tiêu đề không được để trống")
	}
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   false,
	}, nil
}