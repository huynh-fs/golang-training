package repository

import (
	"fmt"
	"time"
)

type Task struct {
	ID      int
	Payload string
}

// mô phỏng việc thực thi một tác vụ.
func (t *Task) Process() {
	fmt.Printf("Bắt đầu xử lý tác vụ %d: %s\n", t.ID, t.Payload)
	time.Sleep(1 * time.Second)
	fmt.Printf(">> Hoàn thành xử lý tác vụ %d\n", t.ID)
}

type TaskRepository interface {
	GetTasks(count int) ([]*Task, error)
}

type memTaskRepository struct{}

func NewMemTaskRepository() TaskRepository {
	return &memTaskRepository{}
}

// giả lập việc lấy các tác vụ từ database.
func (r *memTaskRepository) GetTasks(count int) ([]*Task, error) {
	tasks := make([]*Task, count)
	for i := 0; i < count; i++ {
		tasks[i] = &Task{
			ID:      i + 1,
			Payload: fmt.Sprintf("Data for task %d", i+1),
		}
	}
	return tasks, nil
}