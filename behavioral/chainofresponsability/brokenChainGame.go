package chainofresponsability

import (
	"fmt"
	"sync"
)

// CoR, Mediator, Observer, CQS

type Argument int
const (
	Attack Argument = iota
	Defense
)

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

type Observer interface {
	Handle(q *Query)
}

type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Fire(q *Query)
}

type Game struct {
	observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
	g.observers.Store(o, struct{}{})
}

func (g *Game) Unsubscribe(o Observer) {
	g.observers.Delete(o)
}

func (g *Game) Fire(q *Query) {
	g.observers.Range(
		func(key, value interface{}) bool {
			if key == nil {
				return false
			}

			// must cast to Observer due to it is an interface{}
			key.(Observer).Handle(q)
			return true
		},
	)
}

type Creature2 struct {
	game            *Game
	Name            string
	attack, defense int
}

func NewCreature2(game *Game, name string, attack int, defense int) *Creature2 {
	return &Creature2{
		game:    game,
		Name:    name,
		attack:  attack,
		defense: defense,
	}
}

func (c *Creature2) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature2) Defense() int {
	q := Query{c.Name, Defense, c.defense}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature2) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack(), c.Defense())
}

type CreatureModifier2 struct {
	game     *Game
	creature *Creature2
}

func (c *CreatureModifier2) Handle(q *Query) {
	// 
}

type DoubleAttackModifier2 struct {
	CreatureModifier2
}

func NewDoubleAttackModifier2(g *Game, creature *Creature2) *DoubleAttackModifier2 {
	d := &DoubleAttackModifier2{CreatureModifier2{g, creature}}
	g.Subscribe(d)
	return d
}

func (d *DoubleAttackModifier2) Handle(q *Query) {
	if q.CreatureName == d.creature.Name &&
	q.WhatToQuery == Attack {
		q.Value *= 2
	} 
}

func (d *DoubleAttackModifier2) Close() error {
	d.game.Unsubscribe(d)
	return nil
}
