package main

import (
	"github.com/huynh-fs/file/internal/handler"
	"log"
)

func main() {
	gameHandler := handler.NewGameHandler()

	if err := gameHandler.PlayGame(); err != nil {
		log.Fatalf("Trò chơi kết thúc với lỗi: %v", err)
	}
}