package tests

import (
	"fmt"

	"github.com/renatospaka/design-pattern-go/behavioral/chainofresponsability"
)

func TestChainOfResponsability() {
	fmt.Println("Chain Of Responsability -> method Chain")

	goblin := chainofresponsability.NewCreature("Goblin", 1, 1)
	fmt.Println(goblin.String())
	
	root := chainofresponsability.NewCreatureModifier(goblin)

	// I really didn't get what happened here...
	root.Add(chainofresponsability.NewNoBonusesModifier(goblin))

	root.Add(chainofresponsability.NewDoubleAttackModifier(goblin))
	root.Add(chainofresponsability.NewIncreaseDefenseModifier(goblin))
	root.Add(chainofresponsability.NewDoubleAttackModifier(goblin))
	root.Handle()
	fmt.Println(goblin.String())

	fmt.Println()
	fmt.Println("Chain Of Responsability -> method Chain")
	game := &chainofresponsability.Game{}
	goblin2 := chainofresponsability.NewCreature2(game, "Strong Goblin", 2, 2)
	fmt.Println(goblin2.String())

	{
		m := chainofresponsability.NewDoubleAttackModifier2(game, goblin2)
		fmt.Println(goblin2.String())
		m.Close()
	}
	fmt.Println(goblin2.String())
}
