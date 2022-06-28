package tests

import (
	"fmt"

	"github.com/renatospaka/design-pattern-go/builder"
)

func TestBuilderPattern() {
	// -> Function
	b := builder.PersonBuilderF{}
	personF := b.Called("João").
		WorksAsA("Bartender").
		Build() 
	fmt.Println("BUILDER -> Functional", *personF)

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

	// -> Parameter
	builder.SendEmail(func(b *builder.EmailBuilder) {
		b.
			From("bar@thisclub.com").
			To("bartender@thedomain.com").
			Subject("Work tonight?").
			Body("I have a position for tonight. Are you interested?")
	})
	fmt.Println("BUILDER -> Parameter Finished")
}
