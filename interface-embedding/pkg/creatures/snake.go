package creatures

type Snake struct {
	animal
}

func (s Snake) Speak() string {
	return "Xì xì!"
}

func (s Snake) Move() string {
	return "Lê trên mặt đất"
}

func NewSnake(name string, age int) Snake {
	return Snake{
		animal: animal{Name: name, Age: age},
	}
}