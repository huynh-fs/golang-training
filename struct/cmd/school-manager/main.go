package main

import (
	"github.com/huynh-fs/struct/internal/app/school-manager/handler"
	"github.com/huynh-fs/struct/internal/app/school-manager/service"
)

func main() {
	schoolService := service.NewSchoolService()
	cliHandler := handler.NewCLIHandler(schoolService)
	cliHandler.Run()
}

