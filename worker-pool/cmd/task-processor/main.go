package main

import (
	"fmt"
	"time"

	"github.com/huynh-fs/worker-pool/internal/repository"
	"github.com/huynh-fs/worker-pool/internal/service"
	"github.com/huynh-fs/worker-pool/pkg/config"
	"github.com/huynh-fs/worker-pool/pkg/logger"
)

func main() {
	logger.Init()
	logger.Info.Println("Bắt đầu ứng dụng...")

	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		logger.Error.Fatalf("Lỗi khi tải cấu hình: %v", err)
	}
	logger.Info.Printf("Cấu hình đã được tải: %d workers, %d tasks.", cfg.WorkerPool.Workers, cfg.WorkerPool.Tasks)

	taskRepo := repository.NewMemTaskRepository()

	taskServiceCfg := service.Config{NumWorkers: cfg.WorkerPool.Workers}
	taskService := service.NewTaskService(taskRepo, logger.Info, taskServiceCfg)

	startTime := time.Now()
	taskService.ProcessTasks(cfg.WorkerPool.Tasks)
	duration := time.Since(startTime)

	logger.Info.Println("Tất cả các tác vụ đã được xử lý.")
	fmt.Printf("Chương trình chạy trong: %s\n", duration)
}