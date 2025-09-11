package service

import (
	"github.com/huynh-fs/file/internal/model"
	"fmt"
	"math/rand"
)

func NewTicket() *model.Ticket {
	var ticket model.Ticket

	for col := 0; col < model.TicketSize; col++ {
		min := col*15 + 1
		max := (col + 1) * 15
		usedInCol := make(map[int]bool)

		for row := 0; row < model.TicketSize; row++ {
			if col == 2 && row == 2 {
				continue
			}
			for {
				num := rand.Intn(max-min+1) + min
				if !usedInCol[num] {
					ticket[row][col] = num
					usedInCol[num] = true
					break
				}
			}
		}
	}
	ticket[2][2] = 0 // ô giữa là ô free
	return &ticket
}

func CheckWin(marked [model.TicketSize][model.TicketSize]bool) (bool, string) {
	// kiểm tra hàng ngang
	for row := 0; row < model.TicketSize; row++ {
		if marked[row][0] && marked[row][1] && marked[row][2] && marked[row][3] && marked[row][4] {
			return true, fmt.Sprintf("BINGO theo hàng ngang %d", row+1)
		}
	}
	// kiểm tra cột dọc
	for col := 0; col < model.TicketSize; col++ {
		if marked[0][col] && marked[1][col] && marked[2][col] && marked[3][col] && marked[4][col] {
			return true, fmt.Sprintf("BINGO theo cột dọc %d", col+1)
		}
	}
	// kiểm tra đường chéo chính
	if marked[0][0] && marked[1][1] && marked[2][2] && marked[3][3] && marked[4][4] {
		return true, "BINGO theo đường chéo chính (\\)"
	}
	// kiểm tra đường chéo phụ
	if marked[0][4] && marked[1][3] && marked[2][2] && marked[3][1] && marked[4][0] {
		return true, "BINGO theo đường chéo phụ (/)"
	}
	return false, ""
}