package main

import (
	"github.com/huynh-fs/file/internal/service"
	"github.com/huynh-fs/file/internal/handler"
)

func main() {
	bingoGame := service.NewGameService()
	result := bingoGame.Play()

	if err := handler.WriteResult(result); err != nil {
		panic(err)
	}
}

