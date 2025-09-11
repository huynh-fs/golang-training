package service

import (
	"github.com/huynh-fs/file/internal/model"
	"fmt"
)

type DisplayService struct{}

func NewDisplayService() *DisplayService {
	return &DisplayService{}
}

func (s *DisplayService) PrintInitialPage(p *model.Ticket) {
	fmt.Println("Bắt đầu trò chơi BINGO!")
	fmt.Println("Tấm vé của bạn:")
	s.printTicketLayout(p)
}

func (s *DisplayService) PrintCalledNumber(num int) {
	fmt.Printf("Số vừa ra: %d\n", num)
}

func (s *DisplayService) PrintWinMessage(line string) {
	fmt.Printf("\n%s\n", line)
}

func (s *DisplayService) PrintFinalPage(p *model.Ticket) {
	fmt.Println("\nTấm vé cuối cùng:")
	s.printTicketLayout(p)
}

func (s *DisplayService) printTicketLayout(p *model.Ticket) {
	fmt.Println("-------------------------")
	for r := 0; r < model.TicketSize; r++ {
		for c := 0; c < model.TicketSize; c++ {
			fmt.Printf("%3d |", p[r][c])
		}
		fmt.Println()
	}
	fmt.Println("-------------------------")
}