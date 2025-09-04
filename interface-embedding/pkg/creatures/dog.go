package creatures

type baseAnimal struct {
	Name string
	Age  int
}

type Dog struct {
	baseAnimal
}

func (d *Dog) GetName() string { return d.Name }
func (d *Dog) Speak() string   { return "Gâu gâu!" }
func (d *Dog) Move() string    { return "Chạy bằng bốn chân." }

func NewDog(name string, age int) Dog {
	return Dog{baseAnimal: baseAnimal{Name: name, Age: age}}
}