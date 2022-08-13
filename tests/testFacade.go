package tests

import (
	"fmt"

	"github.com/renatospaka/design-pattern-go/facade"
)

func TestFacade() {
	fmt.Println("Facade => buffer")

	c := facade.NewConsole()
	u := c.GetCharacterAt(1)
	fmt.Println(u)
}
