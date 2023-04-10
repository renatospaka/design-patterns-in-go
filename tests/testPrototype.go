package tests

import (
	"fmt"

	"github.com/renatospaka/design-pattern-go/creational/prototype"
)

func TestPrototypePattern() {
	john := prototype.NewMainOfficeEmployee("John", 123)
	jane := prototype.NewAuxOfficeEmployee("Jane", 987) 
	fmt.Println("PROTOTYPE =>", john)
	fmt.Println("PROTOTYPE =>", jane)

}