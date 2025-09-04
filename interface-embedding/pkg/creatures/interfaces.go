package creatures

type Speaker interface {
	Speak() string
}

type Mover interface {
	Move() string
}