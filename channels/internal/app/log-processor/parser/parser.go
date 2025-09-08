package parser

import (
	"strings"
	"sync"
)

type LogEntry struct {
	Level   string // INFO, ERROR, WARN, DEBUG
	Message string
}

func Worker(wg *sync.WaitGroup, lines <-chan string, entries chan<- LogEntry) {
	defer wg.Done()
	for line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			continue // bỏ qua dòng không hợp lệ
		}
		level := strings.TrimSpace(parts[0])
		if level != "INFO" && level != "WARN" && level != "ERROR" && level != "DEBUG" {
			continue // bỏ qua các level không xác định
		}

		entries <- LogEntry{
			Level:   level,
			Message: strings.TrimSpace(parts[1]),
		}
	}
}