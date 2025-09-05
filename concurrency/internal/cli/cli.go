package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"github.com/huynh-fs/concurrency/internal/fetcher"
)

func displayMenu() {
	fmt.Println("\n--- URL Fetcher Menu ---")
	fmt.Println("1. Lấy tiêu đề từ các URL")
	fmt.Println("2. Thoát")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func processURLs(urls []string) {
	if len(urls) == 0 {
		fmt.Println("Không có URL nào được cung cấp để xử lý.")
		return
	}

	var wg sync.WaitGroup
	resultsChan := make(chan fetcher.Result, len(urls))

	for _, url := range urls {
		wg.Add(1)
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		go fetcher.FetchTitle(url, &wg, resultsChan)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	fmt.Println("\n--- Kết quả ---")
	for result := range resultsChan {
		if result.Err != nil {
			fmt.Printf("Lỗi khi lấy %s: %v\n", result.URL, result.Err)
		} else {
			fmt.Printf("%s -> \"%s\"\n", result.URL, result.Title)
		}
	}
	fmt.Println("--- Hoàn thành ---")
}

func Run() error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		displayMenu()

		var choice string
		if scanner.Scan() {
			choice = scanner.Text()
		} else {
			break
		}
		
		switch strings.TrimSpace(choice) {
		case "1":
			fmt.Println("\nNhập từng URL và nhấn Enter. Nhấn Enter trên dòng trống để bắt đầu:")
			
			var urls []string

			for {
				fmt.Print("URL> ")
				if !scanner.Scan() {
					break
				}
				
				url := strings.TrimSpace(scanner.Text())

				if url == "" {
					break
				}

				urls = append(urls, url)
			}
			
			processURLs(urls)

		case "2":
			fmt.Println("Tạm biệt!")
			return nil
		default:
			fmt.Println("Lựa chọn không hợp lệ, vui lòng thử lại.")
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("lỗi đọc input: %w", err)
	}

	return nil
}