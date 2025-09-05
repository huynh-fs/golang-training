package main

import (
	"fmt"
	"os"
	"github.com/huynh-fs/concurrency/internal/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Lỗi: %v\n", err)
		os.Exit(1)
	}
}