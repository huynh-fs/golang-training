package creatures
type Bird struct {
	baseAnimal
}

func (d *Bird) GetName() string { return d.Name }
func (d *Bird) Speak() string   { return "Chíp chíp!" }
func (d *Bird) Move() string    { return "Bay bằng đôi cánh" }

func NewBird(name string, age int) Bird {
	return Bird{baseAnimal: baseAnimal{Name: name, Age: age}}
}