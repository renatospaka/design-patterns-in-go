package main

import (
	"fmt"

	"github.com/renatospaka/design-pattern-go/builder"
)

func main() {
	// Testing BUILDER
	b := builder.PersonBuilder{}
	person := b.Called("João").
		WorksAsA("Bartender").
		Build() 
	fmt.Println(*person)
}
