package service

import (
	"github.com/huynh-fs/file/internal/model"
	"fmt"
	"time"
)

type GameService struct {
	ticket       *model.Ticket
	marked       [model.TicketSize][model.TicketSize]bool
	randSvc      *RandomService
	calledNumbers []int
}

func NewGameService() *GameService {
	return &GameService{
		ticket:  NewTicket(),
		randSvc: NewRandomService(),
	}
}

func (s *GameService) Play() *model.ResultData {
	fmt.Println("Bắt đầu trò chơi BINGO!")
	fmt.Println("Tấm vé của bạn:")
	s.printTicket(s.ticket)

	s.marked[2][2] = true // ô giữa là ô free

	var winLine string
	for {
		time.Sleep(2 * time.Second)

		num, err := s.randSvc.Draw(1, 75)
		if err != nil {
			fmt.Println("Lỗi:", err)
			break
		}
		s.calledNumbers = append(s.calledNumbers, num)
		fmt.Printf("Số vừa ra: %d\n", num)

		s.markNumber(num)

		win, line := CheckWin(s.marked)
		if win {
			winLine = line
			fmt.Printf("\n%s\n", winLine)
			break
		}
	}

	finalTicket := s.createFinalTicket()
	fmt.Println("\nTấm vé cuối cùng của bạn:")
	s.printMarked()

	return &model.ResultData{
		InitialTicket:   *s.ticket,
		CalledNumbers: s.calledNumbers,
		WinLine:       winLine,
		FinalTicket:     *finalTicket,
	}
}

func (s *GameService) markNumber(num int) {
	for r := 0; r < model.TicketSize; r++ {
		for c := 0; c < model.TicketSize; c++ {
			if s.ticket[r][c] == num {
				s.marked[r][c] = true
				return
			}
		}
	}
}

func (s *GameService) createFinalTicket() *model.Ticket {
	var finalTicket model.Ticket
	isCalled := make(map[int]bool)
	for _, num := range s.calledNumbers {
		isCalled[num] = true
	}

	for r := 0; r < model.TicketSize; r++ {
		for c := 0; c < model.TicketSize; c++ {
			num := s.ticket[r][c]
			if isCalled[num] || num == 0 {
				finalTicket[r][c] = 0
			} else {
				finalTicket[r][c] = num
			}
		}
	}
	return &finalTicket
}

func (s *GameService) printTicket(ticket *model.Ticket) {
	fmt.Println("-----------------")
	for r := 0; r < model.TicketSize; r++ {
		for c := 0; c < model.TicketSize; c++ {
			if ticket[r][c] == 0 {
				fmt.Printf(" * ")
			} else {
				fmt.Printf("%2d ", ticket[r][c])
			}
		}
		fmt.Println()
	}
	fmt.Println("-----------------")
}

func (s *GameService) printMarked() {
	fmt.Println("-----------------")
	for r := 0; r < model.TicketSize; r++ {
		for c := 0; c < model.TicketSize; c++ {
			if s.marked[r][c] || s.ticket[r][c] == 0 {
				fmt.Printf(" * ")
			} else {
				fmt.Printf("%2d ", s.ticket[r][c])
			}
		}
		fmt.Println()
	}
}
