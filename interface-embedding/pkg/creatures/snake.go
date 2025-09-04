package creatures

type Snake struct {
	baseAnimal
}

func (s *Snake) GetName() string { return s.Name }
func (s *Snake) Speak() string { return "Xì xì!" }
func (s *Snake) Move() string { return "Lê trên mặt đất" }

func NewSnake(name string, age int) Snake {
	return Snake{baseAnimal: baseAnimal{Name: name, Age: age}}
}