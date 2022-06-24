package main

import (
	"fmt"

	"github.com/renatospaka/design-pattern-go/builder"
	"github.com/renatospaka/design-pattern-go/factory"
)

func main() {
	// **********
	// ** BUILDER
	// **********
	testBuilderPattern() 
	
	// **********
	// ** FACTORY
	// **********
	testFactoryPattern()
}

func testBuilderPattern() {
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

func testFactoryPattern() {
	// -> Function
	personFuncFac := factory.NewPersonFuncFactory("John", 33)
	fmt.Println("FACTORY -> Functional", personFuncFac)

	// -> Interface
	personInterfFac := factory.NewPersonInterfFactory("John", 33)
	personInterfFac.SayHello()
	personInterfFac = factory.NewPersonInterfFactory("Sam", 83)
	personInterfFac.SayHello()
	fmt.Println("FACTORY -> Interface", personInterfFac)

	// -> Generator
	developerFactory := factory.NewEmployeeFactory("developer", 60000)
	developer := developerFactory("John")
	fmt.Println("FACTORY -> Generator (Function)", developer)
	
	managerFactory := factory.NewEmployeeFactory("manager", 80000)
	manager := managerFactory("Jane")
	fmt.Println("FACTORY -> Generator (Function)", manager)

	bossFactory := factory.NewEmployeeFactory2("CEO", 110000)
	boss := bossFactory.Create("Sam")
	fmt.Println("FACTORY -> Generator (Struct)", boss)
}

