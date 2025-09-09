package service

import (
	"fmt"
	"time"

	"github.com/huynh-fs/file/internal/bingo"
	"github.com/huynh-fs/file/internal/output"
	"github.com/huynh-fs/file/internal/random"
)

const outputFile = "bingo_result.csv"

type Game struct {
	ticket 	 *bingo.Ticket
	marked   [bingo.TicketSize][bingo.TicketSize]bool
	randGen *random.Generator
	calledNumbers []int
}

func NewGame() *Game {
	return &Game{
		ticket:   bingo.NewTicket(),
		randGen:  random.NewGenerator(),
	}
}

func (g *Game) Play() error {
	fmt.Println(("Bắt đầu trò chơi Bingo!"))
	fmt.Println("Vé Bingo của bạn là:")
	g.printticket(g.ticket)

	g.marked[bingo.TicketSize/2][bingo.TicketSize/2] = true // đánh dấu ô trống ở giữa

	var win bool 
	var winline string

	for {
		time.Sleep(1 * time.Second)
		num, err := g.randGen.Draw(1, 75)
		if err != nil {
			fmt.Println("Lỗi khi rút số:", err)
			break
		}
		g.calledNumbers = append(g.calledNumbers, num)
		fmt.Printf("Số được rút: %d\n", num)
		g.markNumber(num)
		win, winline = bingo.CheckWinner(g.marked)
		if win {
			fmt.Printf("Chúc mừng! Bạn đã thắng với %s!\n", winline)
			break
		}
	}

	finalTicket := g.generateFinalticket()
	result := output.ResultData{
		InitialTicket: *g.ticket,
		CalledNumbers: g.calledNumbers,
		Winline: winline,
		FinalTicket: finalTicket,
	}

	if err := output.WriteResultToCSV(outputFile, result); err != nil {
		panic(err) 
	} else {
		fmt.Printf("Kết quả đã được lưu vào %s\n", outputFile)
	}
	g.printMarked()
	return nil
}


func (g *Game) markNumber(num int) {
	for r := 0; r < bingo.TicketSize; r++ {
		for c := 0; c < bingo.TicketSize; c++ {
			if g.ticket[r][c] == num {
				g.marked[r][c] = true
				return
			}
		}
	}
}

func (g *Game) generateFinalticket() [bingo.TicketSize][bingo.TicketSize]int {
	var finalticket [bingo.TicketSize][bingo.TicketSize]int
	isCalled := make(map[int]bool)
	for _, num := range g.calledNumbers {
		isCalled[num] = true
	}

	for r := 0; r < bingo.TicketSize; r++ {
		for c := 0; c < bingo.TicketSize; c++ {
			num := g.ticket[r][c]
			if num == 0 || isCalled[num] {
				finalticket[r][c] = 0
			} else {
				finalticket[r][c] = num
			}
		}
	}
	return finalticket
}

func (g *Game) printticket(ticket *bingo.Ticket) {
	fmt.Println("-----------------")
	for r := 0; r < bingo.TicketSize; r++ {
		for c := 0; c < bingo.TicketSize; c++ {
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

func (g *Game) printMarked() {
	fmt.Println("Vé đã đánh dấu:")
	fmt.Println("----------------")
	for r := 0; r < bingo.TicketSize; r++ {
		for c := 0; c < bingo.TicketSize; c++ {
			if g.marked[r][c] {
				fmt.Printf(" * ")
			} else {
				fmt.Printf("%2d ", g.ticket[r][c])
			}	
		}
		fmt.Println()
	}
	fmt.Println("----------------")
}