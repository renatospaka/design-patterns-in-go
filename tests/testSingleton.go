package tests

import (
	"fmt"

	"github.com/renatospaka/design-pattern-go/singleton"
)

func TestSingleton() {
	db := singleton.GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Printf("SINGLETON => Population of %s = %d\n", "Seoul", pop)
}