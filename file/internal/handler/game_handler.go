package handler

import (
	"github.com/huynh-fs/file/internal/model"
	"github.com/huynh-fs/file/internal/service"
	"github.com/huynh-fs/file/pkg/output"
	"sync"
)

type GameHandler struct {
	gameSvc    *service.GameService
	displaySvc *service.DisplayService
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		gameSvc:    service.NewGameService(),
		displaySvc: service.NewDisplayService(),
	}
}

func (h *GameHandler) PlayGame() error {
	events := make(chan model.GameEvent)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for event := range events {
			switch event.Type {
			case model.InitialTicketRender:
				h.displaySvc.PrintInitialPage(event.Ticket)
			case model.NumberDrawn:
				h.displaySvc.PrintCalledNumber(event.Number)
			case model.GameWon:
				h.displaySvc.PrintWinMessage(event.Message)
			}
		}
	}()

	resultData := h.gameSvc.Play(events)

	wg.Wait()

	h.displaySvc.PrintFinalPage(&resultData.FinalTicket)

	return output.WriteToCSV(resultData)
}