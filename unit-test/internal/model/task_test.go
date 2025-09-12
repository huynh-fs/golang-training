package model

import "testing"

func TestNewTask(t *testing.T) {
	t.Run("Tạo task thành công", func(t *testing.T) {
		title := "Học Unit Test"
		desc := "Học cách viết unit test trong Go"
		task, err := NewTask(1, title, desc)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi, nhưng nhận được: %v", err)
		}
		if task == nil {
			t.Fatal("Task không được là nil khi không có lỗi")
		}
		if task.Title != title {
			t.Errorf("got title %q, want %q", task.Title, title)
		}
		if task.Completed {
			t.Error("Task mới không nên được hoàn thành")
		}
	})

	t.Run("Thất bại khi tiêu đề trống", func(t *testing.T) {
		_, err := NewTask(1, "", "Mô tả")
		if err == nil {
			t.Fatal("Mong đợi lỗi khi tiêu đề trống, nhưng không có lỗi")
		}
	})
}