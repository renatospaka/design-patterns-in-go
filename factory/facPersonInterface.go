package factory

import "fmt"

type PersonInterfFactory interface {
	SayHello()
}

type person struct {
	name string
	age int
}

type tiredPerson struct {
	name string
	age int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old\n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Sorry, I'm too tired\n")
}
func NewPersonInterfFactory(name string, age int) PersonInterfFactory {
	// Notice that there is no need to use a pointer 
	// in PersonInterfFactory due to it is an interface

	if age > 80 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}
