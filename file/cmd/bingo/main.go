package main

import (
	"github.com/huynh-fs/file/internal/service"
)

func main() {
	bingoGame := service.NewGame()
	bingoGame.Play()
}