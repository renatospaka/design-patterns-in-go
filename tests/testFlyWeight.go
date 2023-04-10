package tests

import (
	"fmt"

	"github.com/renatospaka/design-pattern-go/structural/flyweight"
)

func TestFlyWeight() {
	fmt.Println("Flyweight => text processor")
	text := "This is a brave new world"

	ft := flyweight.NewFormattedText(text)
	ft.Capitalize(10, 15)
	fmt.Println(ft.String())

	bft :=  flyweight.NewBetterFormattedText((text))
	bft.Range(16, 19).Capitalize = true
	fmt.Println(bft.String())
	
	fmt.Println("Flyweight => common names")
	john := flyweight.NewUser("John Doe")
	jane := flyweight.NewUser("Jane Doe")
	alsoJohn := flyweight.NewUser("John Smith")
	fmt.Println(john.FullName)
	fmt.Println(jane.FullName)
	fmt.Println(alsoJohn.FullName)
	fmt.Println("Memory taken by users:", 
		len([]byte(john.FullName)) +
		len([]byte(jane.FullName)) +
		len([]byte(alsoJohn.FullName)))
	
	john2 := flyweight.NewUser2("John Doe")
	jane2 := flyweight.NewUser2("Jane Doe")
	alsoJohn2 := flyweight.NewUser2("John Smith")
	fmt.Println(john2.FullName())
	fmt.Println(jane2.FullName())
	fmt.Println(alsoJohn2.FullName())
	fmt.Println("Memory taken by users2:", 
		john2.Memory() +
		jane2.Memory() +
		alsoJohn2.Memory())
}
