package model

type GameEventType = int

const (
	InitialTicketRender GameEventType = iota // thời điểm khởi tạo ticket
	NumberDrawn // thời điểm rút số
	GameWon // thời điểm BINGO
)

type GameEvent struct {
	Type GameEventType
	Ticket *Ticket
	Number int // dùng cho sự kiện có số mới được rút ra
	Message string // dùng cho sự kiện BINGO để lưu winline
}