package decorator

import "fmt"

type Aged interface {
	Age() int
	SetAge(age int)
}

type Bird struct {
	age int
}

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying.")
	}
}

func (b *Bird) Age() int {
	return b.age
}

func (b *Bird) SetAge(age int) {
	b.age = age
}

type Lizard struct {
	age int
}

func (l *Lizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling.")
	}
}

func (l *Lizard) Age() int {
	return l.age
}

func (l *Lizard) SetAge(age int) {
	l.age = age
}

type Dragon struct {
	bird Bird
	lizard Lizard
}

func NreDragon() *Dragon {
	return &Dragon{
		Bird{},
		Lizard{},
	}
}

func (d *Dragon) Age() int {
	return d.bird.age
}

func (d *Dragon) SetAge(age int) {
	d.bird.age = age
	d.lizard.age = age
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}