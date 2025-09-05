package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sync"
)

var titleRegex = regexp.MustCompile(`(?i)<title>(.*?)</title>`)

func FetchAndPrintTitle(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Lỗi khi lấy %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Lỗi khi đọc %s: %v\n", url, err)
		return
	}

	matches := titleRegex.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		fmt.Printf("Lỗi khi lấy %s: không tìm thấy thẻ title\n", url)
		return
	}

	fmt.Printf("%s -> \"%s\"\n", url, matches[1])
}