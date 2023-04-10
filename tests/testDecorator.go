package tests

import (
	"fmt"

	"github.com/renatospaka/design-pattern-go/structural/decorator"
)

func TestDecorator() {
	fmt.Println("Decorator => multi aggregation")
	dragon := decorator.Dragon{}
	dragon.SetAge(10)
	dragon.Fly()
	dragon.Crawl()

	fmt.Println("Decorator => geometric structure")
	circle := decorator.Circle{Radius: 2.0}
	circle.Resize(2)
	fmt.Println(circle.Render())
	
	redCircle := decorator.ColoredShape{Shape: &circle, Color: "Red"}
	fmt.Println(redCircle.Render())

	rhsCircle := decorator.TransparentShape{&redCircle, 0.5}
	fmt.Println(rhsCircle.Render())
}