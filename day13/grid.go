package main

import (
	"fmt"
	"strings"
)

const possibles = "-|/\\+"

type Grid struct {
	Tracks   []string
	AllCarts Carts
}

func (g Grid) String() string {
	return strings.Join(g.Tracks, "\n")
}

func (g *Grid) NextTrack(x, y int) byte {
	//fmt.Println("Trying to read:", c.X, nextY)
	//fmt.Println("X:", g.Tracks)
	//fmt.Println("Reading:", y, x)

	if x == 150 && y == 145 {
		dump(g)
	}
	return g.Tracks[y][x]
}

func (g *Grid) UpdateMove(c *Cart, xdif, ydif int) (*Cart, *Cart) {
	if c.Crashed {
		return nil, nil
	}

	newX := c.X + xdif
	newY := c.Y + ydif

	// see if it's a collision
	var coll *Cart
	for _, cc := range g.AllCarts {
		if cc.X == newX && cc.Y == newY {
			coll = cc
			break
		}
	}

	// set the track back to what is under the cart
	row := g.Tracks[c.Y]
	g.Tracks[c.Y] = row[:c.X] + string(c.TrackUnderneath) + row[c.X+1:]

	// set the cart to the new location
	c.X = newX
	c.Y = newY

	if coll != nil {
		coll.Crashed = true
		c.Crashed = true
		// remove the collider by putting the track back
		row = g.Tracks[coll.Y]
		g.Tracks[coll.Y] = row[:coll.X] + string(coll.TrackUnderneath) + row[coll.X+1:]
	} else {
		// record what was under the cart
		c.TrackUnderneath = rune(g.Tracks[c.Y][c.X])
		// redraw the cart in the new location
		row = g.Tracks[c.Y]
		g.Tracks[c.Y] = row[:c.X] + string(c.Direction) + row[c.X+1:]

		if i := strings.IndexRune(possibles, c.TrackUnderneath); i < 0 {
			dump(g)
			panic(fmt.Sprintf("We blorfed at [%v, %v] => %s", c.X, c.Y, string(c.TrackUnderneath)))
		}
	}

	return coll, c
}

func (g *Grid) Click() (a, b *Cart) {
	g.AllCarts.Sort()
	for _, c := range g.AllCarts {
		//fmt.Println("Click cart:", *c)
		nextX := 0
		nextY := 0
		switch c.Direction {
		case '>', '<':
			nextX = 1
			if c.Direction == '<' {
				nextX *= -1
			}
			nextTrack := g.NextTrack(c.X+nextX, c.Y)
			switch nextTrack {
			case '\\':
				c.RightTurn()
			case '/':
				c.LeftTurn()
			case '+':
				c.Turn()
			}

		case 'v', '^':
			nextY = 1
			if c.Direction == '^' {
				nextY *= -1
			}
			nextTrack := g.NextTrack(c.X, c.Y+nextY)
			switch nextTrack {
			case '\\':
				c.LeftTurn()
			case '/':
				c.RightTurn()
			case '+':
				c.Turn()
			}
		}

		// update cart and grid state
		a, b = g.UpdateMove(c, nextX, nextY)
	}

	carts := Carts{}
	for _, c := range g.AllCarts {
		if !c.Crashed {
			carts = append(carts, c)
		} else {
			fmt.Println("Removing:", *c)
		}
	}
	(*g).AllCarts = carts
	return nil, nil
}

func (g *Grid) Run2(maxTicks int) (int, *Cart) {
	for i := 0; i < maxTicks; i++ {
		g.Click()
		if len(g.AllCarts) == 1 {
			return i, g.AllCarts[0]
		}
	}
	return maxTicks, nil
}

// run simulation and return location and tick of first collision
func (g *Grid) Run(maxTicks int) (int, *Cart, *Cart) {
	for i := 0; i < maxTicks; i++ {
		a, b := g.Click()
		if a != nil {
			return i, a, b
		}
	}

	return maxTicks, nil, nil
}

func dump(g *Grid) {
	fmt.Println(g)

	/*
		for _, c := range g.AllCarts {
			fmt.Println(string(c.Direction), c.X, c.Y)
		}
	*/
}
