package handler

import (
	"fmt"
	"github.com/huynh-fs/interface-embedding/internal/app/zoo/service"
	"github.com/huynh-fs/interface-embedding/pkg/creatures"
)

type CLIHandler struct {
	zooService *service.ZooService
}

func NewCLIHandler(s *service.ZooService) *CLIHandler {
	return &CLIHandler{zooService: s}
}

func (h *CLIHandler) Run() {
	fmt.Println("CHÀO MỪNG ĐẾN VỚI SỞ THÚ GO!")

	dog := creatures.NewDog("Chó Mực", 5)
	snake := creatures.NewSnake("Rắn Hổ Mang", 3)
	bird := creatures.NewBird("Chim Chích Chòe", 2)

	h.zooService.AddCreature(&dog)
	h.zooService.AddCreature(&snake)
	h.zooService.AddCreature(&bird)

	h.zooService.GenerateReport()
}