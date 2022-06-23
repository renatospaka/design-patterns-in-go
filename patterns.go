package main

import (
	"fmt"

	"github.com/renatospaka/design-pattern-go/builder"
)

func main() {
	// **********
	// ** BUILDER
	// **********
	testBuilderPattern()
}

func testBuilderPattern() {
	// -> Function
	b := builder.PersonBuilderF{}
	personF := b.Called("João").
		WorksAsA("Bartender").
		Build() 
	fmt.Println("BUILDER -> Function", *personF)

	// -> Facet
	pb := builder.NewPersonBuilderFa()
	pb.
		Lives().
			At("456, Any Street").
			In("London").
			WithPostcode("00A222X").
		Works().
			At("This Club").
			AsA("Bartender").
			Earning(123000)
	personFa := pb.Build() 
	fmt.Println("BUILDER -> Facet", *personFa)
}