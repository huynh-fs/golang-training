package main

import (
	"github.com/huynh-fs/interface-embedding/internal/app/zoo/handler"
	"github.com/huynh-fs/interface-embedding/internal/app/zoo/service"
)

func main() {
	zooService := service.NewZooService()
	cliHandler := handler.NewCLIHandler(zooService)
	cliHandler.Run()
}