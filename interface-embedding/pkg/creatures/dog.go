package creatures

type animal struct {
	Name string
	Age int
}

func (a animal) GetName() string {
	return a.Name
}

type Dog struct {
	animal
}

func (d Dog) Speak() string {
	return "Gâu gâu!"
}

func (d Dog) Move() string {
	return "Chạy bằng bốn chân"
}

func NewDog(name string, age int) Dog {
	return Dog{
		animal: animal{ Name: name, Age: age },
	}
}
