package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sync"
)

type Result struct {
	URL        string
	Title      string
	Err        error
}

var titleRegex = regexp.MustCompile(`(?i)<title>(.*?)</title>`)

func FetchTitle(url string, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		results <- Result{URL: url, Err : err}
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		results <- Result{URL: url, Err : err}
		return
	}

	matches := titleRegex.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		results <- Result{URL: url, Err : fmt.Errorf("no title found")}
		return
	}

	results <- Result{URL: url, Title: matches[1]}
}