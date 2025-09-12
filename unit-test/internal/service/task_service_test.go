package service

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	s := NewTaskService()

	// dùng Table-Driven
	testCases := []struct {
		name        string
		title       string
		description string
		expectError bool
	}{
		{"Tạo thành công", "Task 1", "Mô tả 1", false},
		{"Tiêu đề trống", "", "Mô tả 2", true},
		{"Chỉ có tiêu đề", "Task 3", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			task, err := s.CreateTask(tc.title, tc.description)

			if tc.expectError {
				if err == nil {
					t.Errorf("Mong đợi lỗi nhưng nhận được nil")
				}
			} else {
				if err != nil {
					t.Errorf("Không mong đợi lỗi nhưng nhận được: %v", err)
				}
				if task.Title != tc.title {
					t.Errorf("got title %q, want %q", task.Title, tc.title)
				}
			}
		})
	}
}

func TestGetTask(t *testing.T) {
	s := NewTaskService()
	createdTask, _ := s.CreateTask("Task hiện có", "Mô tả")

	t.Run("Lấy task tồn tại", func(t *testing.T) {
		task, err := s.GetTask(createdTask.ID)
		if err != nil {
			t.Fatalf("Không mong đợi lỗi nhưng nhận được: %v", err)
		}
		if task.ID != createdTask.ID {
			t.Errorf("got ID %d, want %d", task.ID, createdTask.ID)
		}
	})

	t.Run("Lấy task không tồn tại", func(t *testing.T) {
		_, err := s.GetTask(999) // id không tồn tại
		if err == nil {
			t.Fatal("Mong đợi lỗi nhưng nhận được nil")
		}
	})
}

func TestCompleteTask(t *testing.T) {
	s := NewTaskService()
	task, _ := s.CreateTask("Task chưa hoàn thành", "Mô tả")

	t.Run("Hoàn thành task tồn tại", func(t *testing.T) {
		err := s.CompleteTask(task.ID)
		if err != nil {
			t.Fatalf("Không mong đợi lỗi nhưng nhận được: %v", err)
		}

		completedTask, _ := s.GetTask(task.ID)
		if !completedTask.Completed {
			t.Error("Task đáng lẽ phải được hoàn thành")
		}
	})

	t.Run("Hoàn thành task không tồn tại", func(t *testing.T) {
		err := s.CompleteTask(999)
		if err == nil {
			t.Fatal("Mong đợi lỗi nhưng nhận được nil")
		}
	})

	t.Run("Hoàn thành task đã hoàn thành", func(t *testing.T) {
		s.CompleteTask(task.ID)
		err := s.CompleteTask(task.ID)
		if err == nil {
			t.Fatal("Mong đợi lỗi khi hoàn thành task đã hoàn thành, nhưng nhận được nil")
		}
	})
}