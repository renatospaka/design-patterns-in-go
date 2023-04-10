package tests

import (
	"fmt"

	"github.com/renatospaka/design-pattern-go/creational/factory"
)

func TestFactoryPattern() {
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

	// -> Prototype
	managerProto := factory.NewEmployeeFacProto(factory.Manager)
	managerProto.Name = "Jane"
	fmt.Println("FACTORY -> Prototype", managerProto)
}