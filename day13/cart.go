package main

import (
	"fmt"
	"sort"
	"strconv"
)

var RightTurns = map[rune]rune{'^': '>', '>': 'v', 'v': '<', '<': '^'}

type Cart struct {
	// one of <,>,^,v
	Direction       rune
	TrackUnderneath rune
	Crashed         bool
	LastTurn        Turn
	X               int
	Y               int
}

func NewCart(x, y int, dir rune) *Cart {
	t := '|'
	if dir == '>' || dir == '<' {
		t = '-'
	}

	return &Cart{
		Direction:       dir,
		TrackUnderneath: t,
		LastTurn:        RightTurn,
		X:               x,
		Y:               y,
	}
}

func (c Cart) String() string {
	return fmt.Sprintf("Dir: %s, Last: %v, [%v, %v]",
		string(c.Direction), c.LastTurn, c.X, c.Y)
}

func (c *Cart) RightTurn() { c.Direction = RightTurns[c.Direction] }

func (c *Cart) LeftTurn() {
	c.RightTurn()
	c.RightTurn()
	c.RightTurn()
}

func (c *Cart) Turn() {
	switch c.LastTurn {
	case Straight:
		c.RightTurn()
		c.LastTurn = RightTurn
	case LeftTurn:
		// straight
		c.LastTurn = Straight
	case RightTurn:
		c.LeftTurn()
		c.LastTurn = LeftTurn
	}
}

type Carts []*Cart

func (c Carts) Sort() {
	// update carts in sequence
	sort.Slice(c, func(i, j int) bool {
		if c[i].Y < c[j].Y {
			return true
		}
		if c[i].Y == c[j].Y {
			return c[i].X < c[j].X
		}
		return false
	})
}

func (carts Carts) Collisions() Carts {
	m := make(map[string]Cart, len(carts))
	coll := Carts{}
	for _, c := range carts {
		key := strconv.Itoa(c.X) + strconv.Itoa(c.Y)
		if v, ok := m[key]; ok {
			coll = append(coll, &v)
			coll = append(coll, c)
		}
		m[key] = *c
	}
	return coll
}
