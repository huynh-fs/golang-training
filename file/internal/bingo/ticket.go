package bingo

import (
	"fmt"
	"math/rand"
)

const TicketSize = 5

type Ticket [TicketSize][TicketSize]int

func NewTicket() *Ticket {
	var ticket Ticket
	for col := 0; col < TicketSize; col++ {
		min := col*15 + 1
		max := (col + 1) * 15
		usedInCol := make(map[int]bool)
		for row := 0; row < TicketSize; row++ {
			for {
				num := rand.Intn(max-min+1) + min
				if !usedInCol[num] {
					usedInCol[num] = true
					ticket[row][col] = num
					break
				}
			}
		}
	}
	ticket[TicketSize/2][TicketSize/2] = 0 // ô giữa là ô free
	return &ticket
}

func CheckWinner(marked [TicketSize][TicketSize]bool) (bool, string) {
	for row := 0; row < TicketSize; row++ {
		if marked[row][0] && marked[row][1] && marked[row][2] && marked[row][3] && marked[row][4] {
			return true, fmt.Sprintf("hàng ngang %d", row+1)
		}
	}

	for col := 0; col < TicketSize; col++ {
		if marked[0][col] && marked[1][col] && marked[2][col] && marked[3][col] && marked[4][col] {
			return true, fmt.Sprintf("hàng dọc %d", col+1)
		}
	}

	if marked[0][0] && marked[1][1] && marked[2][2] && marked[3][3] && marked[4][4] {
		return true, "đường chéo chính (/)"
	}

	if marked[0][4] && marked[1][3] && marked[2][2] && marked[3][1] && marked[4][0] {
		return true, "đường chéo phụ (\\)"
	}
	return false, ""
}