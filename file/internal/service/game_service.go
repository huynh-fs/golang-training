package service

import (
	"github.com/huynh-fs/file/internal/model"
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

func (s *GameService) Play(events chan <- model.GameEvent) *model.ResultData {
	defer close(events)

	s.marked[2][2] = true // ô giữa là ô free

	events <- model.GameEvent{Type: model.InitialTicketRender, Ticket: s.ticket}

	var winLine string
	for {
		time.Sleep(2 * time.Second)

		num, err := s.randSvc.Draw(1, 75)
		if err != nil {
			break
		}
		s.calledNumbers = append(s.calledNumbers, num)
		s.markNumber(num)

		events <- model.GameEvent{Type: model.NumberDrawn, Number: num}

		win, line := CheckWin(s.marked)
		if win {
			winLine = line
			events <- model.GameEvent{Type: model.GameWon, Message: line}
			break
		}
	}

	finalTicket := s.createFinalTicket()

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
