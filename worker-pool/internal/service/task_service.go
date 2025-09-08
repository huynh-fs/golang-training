package service

import (
	"log"
	"sync"
	"github.com/huynh-fs/worker-pool/internal/repository"
)

type TaskService struct {
	taskRepo repository.TaskRepository
	logger   *log.Logger
	cfg      Config
}

type Config struct {
	NumWorkers int
}

func NewTaskService(repo repository.TaskRepository, logger *log.Logger, cfg Config) *TaskService {
	return &TaskService{
		taskRepo: repo,
		logger:   logger,
		cfg:      cfg,
	}
}

func (s *TaskService) ProcessTasks(numTasks int) {
	tasksToProcess, err := s.taskRepo.GetTasks(numTasks)
	if err != nil {
		s.logger.Fatalf("Không thể lấy tác vụ: %v", err)
	}

	tasksChan := make(chan *repository.Task, len(tasksToProcess))
	var wg sync.WaitGroup

	s.logger.Printf("Khởi tạo %d workers...", s.cfg.NumWorkers)
	for i := 1; i <= s.cfg.NumWorkers; i++ {
		wg.Add(1)
		go s.worker(i, &wg, tasksChan)
	}
	for _, task := range tasksToProcess {
		tasksChan <- task
	}
	close(tasksChan) // đóng channel để báo hiệu hết việc

	wg.Wait()
}

// một goroutine sẽ xử lý công việc từ channel.
func (s *TaskService) worker(id int, wg *sync.WaitGroup, tasks <-chan *repository.Task) {
	defer wg.Done()
	s.logger.Printf("Worker %d đã bắt đầu", id)
	for task := range tasks {
		task.Process()
	}
	s.logger.Printf("Worker %d đã kết thúc", id)
}