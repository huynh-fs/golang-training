package aggregator

import (
	"fmt"
	"time"

	"github.com/huynh-fs/channels/internal/app/log-processor/parser"
)

type Stats struct {
	CountByLevel map[string]int
	TotalParsed  int
}

func Run(entries <-chan parser.LogEntry) <-chan *Stats {
	finalStatsChan := make(chan *Stats, 1) // buffered 1 để tránh block

	go func() {
		stats := &Stats{
			CountByLevel: make(map[string]int),
		}
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		fmt.Println("Aggregator: Bắt đầu thu thập kết quả...")

		for {
			select {
			case entry, ok := <-entries:
				if !ok { // channel đã đóng
					fmt.Println("Aggregator: Đã xử lý xong. Gửi báo cáo cuối cùng.")
					finalStatsChan <- stats
					close(finalStatsChan)
					return
				}
				stats.CountByLevel[entry.Level]++
				stats.TotalParsed++
			case <-ticker.C:
				fmt.Printf("-> Tiến độ: Đã xử lý %d dòng...\n", stats.TotalParsed)
			}
		}
	}()

	return finalStatsChan
}