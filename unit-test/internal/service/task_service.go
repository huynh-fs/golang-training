package service

import (
	"fmt"
	"github.com/huynh-fs/unit-test/internal/model"
	"sync"
)

type TaskService struct {
	tasks  map[int]*model.Task
	nextID int
	mu     sync.Mutex 
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  make(map[int]*model.Task),
		nextID: 1,
	}
}

func (s *TaskService) CreateTask(title, description string) (*model.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, err := model.NewTask(s.nextID, title, description)
	if err != nil {
		return nil, err
	}

	s.tasks[s.nextID] = task
	s.nextID++
	return task, nil
}

func (s *TaskService) GetTask(id int) (*model.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return nil, fmt.Errorf("không tìm thấy task với ID: %d", id)
	}
	return task, nil
}

func (s *TaskService) CompleteTask(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return fmt.Errorf("không tìm thấy task với ID: %d", id)
	}

	if task.Completed {
		return fmt.Errorf("task với ID %d đã được hoàn thành trước đó", id)
	}

	task.Completed = true
	return nil
}