package service

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/huynh-fs/channels/internal/app/log-processor/aggregator"
	"github.com/huynh-fs/channels/internal/app/log-processor/parser"
)

type ProcessorService struct {
	numParserWorkers int
}

func New(numParserWorkers int) *ProcessorService {
	return &ProcessorService{numParserWorkers: numParserWorkers}
}

func (s *ProcessorService) Run(logFilePath string) (*aggregator.Stats, error) {
	file, err := os.Open(logFilePath)
	if err != nil {
		return nil, fmt.Errorf("không thể mở file: %w", err)
	}
	defer file.Close()

	linesChan := make(chan string, 1000)
	entriesChan := make(chan parser.LogEntry, 1000)
	var wgParsers sync.WaitGroup

	finalStatsChan := aggregator.Run(entriesChan)

	wgParsers.Add(s.numParserWorkers)
	for i := 0; i < s.numParserWorkers; i++ {
		go parser.Worker(&wgParsers, linesChan, entriesChan)
	}	

	go func() {
		wgParsers.Wait()
		close(entriesChan)
		fmt.Println("Service: Tất cả parser đã hoàn thành.")
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesChan <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("lỗi khi đọc file: %w", err)
	}
	close(linesChan) // đóng linesChan để báo cho các parser biết đã hết việc
	fmt.Println("Service: Đã đọc xong file.")

	finalStats := <-finalStatsChan
	return finalStats, nil
}