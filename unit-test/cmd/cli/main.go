package main

import (
	"github.com/huynh-fs/unit-test/internal/handler"
)

func main() {
	cli := handler.NewCLIHandler()

	cli.Run()
}