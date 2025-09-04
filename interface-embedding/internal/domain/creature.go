package domain

type Creature interface {
	GetName() string
	Speak() string
	Move() string
}
