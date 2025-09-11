package model

type ResultData struct {
	InitialTicket   Ticket
	CalledNumbers []int
	WinLine       string
	FinalTicket     Ticket
}