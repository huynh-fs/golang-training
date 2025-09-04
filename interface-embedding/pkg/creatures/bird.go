package creatures

type Bird struct {
	animal
}

func (b Bird) Speak() string {
	return "Chíp chíp!"
}

func (b Bird) Move() string {
	return "Bay bằng đôi cánh"
}

func NewBird(name string, age int) Bird {
	return Bird{
		animal: animal{Name: name, Age: age},
	}
}
