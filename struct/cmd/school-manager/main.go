package main

import (
	"github.com/huynh-fs/struct/internal/app/school-manager/handler"
	"github.com/huynh-fs/struct/internal/app/school-manager/service"
)

func main() {
	classService := service.NewClassService()
	cliHandler := handler.NewCLIHandler(classService)
	cliHandler.Run()
}

