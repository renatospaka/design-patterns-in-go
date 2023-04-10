package tests

import (
	"fmt"

	"github.com/renatospaka/design-pattern-go/creational/singleton"
)

func TestSingleton() {
	db := singleton.GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Printf("SINGLETON => Population of %s = %d\n", "Seoul", pop)

	//UT
	cities := []string{"Seoul", "Mexico City"}
	tp := singleton.GetTotalPopulation(cities)
	fmt.Printf("SINGLETON => Population of %s = %d => OK=(%t)\n", cities, tp, tp == (17500000 + 17400000))

	//UT + DIP
	cities = []string{"alpha", "gamma"}
	tp = singleton.GetTotalPopulationEx(&singleton.DummyDatabase{}, cities)
	fmt.Printf("SINGLETON => Population of %s = %d => OK=(%t)\n", cities, tp, tp == 4)
}
