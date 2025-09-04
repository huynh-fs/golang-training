package main

import (
	"fmt"
	"github.com/huynh-fs/interface-embedding/pkg/creatures"
	"github.com/huynh-fs/interface-embedding/internal/service"
)

func main() {
	zooService := service.NewZooService(nil)

	dog := creatures.NewDog("Mực", 3)
	snake := creatures.NewSnake("Lê", 2)
	bird := creatures.NewBird("Ổi", 1)

	zooService.AddCreature(&dog)
	zooService.AddCreature(&snake)
	zooService.AddCreature(&bird)

	fmt.Println("=======Báo cáo sở thú: ======")
	zooService.GenerateReport()
}