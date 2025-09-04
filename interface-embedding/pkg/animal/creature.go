package animal

type Creature interface {
	GetName() string
	Speak() string
	Move() string
}
