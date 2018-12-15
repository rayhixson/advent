package main

import (
	"fmt"
	"strings"
)

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
	return g.Tracks[y][x]
}

func (g *Grid) UpdateMove(c *Cart, xdif, ydif int) (*Cart, *Cart) {
	newX := c.X + xdif
	newY := c.Y + ydif

	// see if it's a collision
	for _, cc := range g.AllCarts {
		if cc.X == newX && cc.Y == newY {
			return cc, c
		}
	}

	// set the track back to what is under the cart
	row := g.Tracks[c.Y]
	g.Tracks[c.Y] = row[:c.X] + string(c.TrackUnderneath) + row[c.X+1:]

	// set the cart to the new location
	c.X = newX
	c.Y = newY

	// record what was under the cart
	c.TrackUnderneath = rune(g.Tracks[c.Y][c.X])

	// redraw the cart in the new location
	row = g.Tracks[c.Y]
	g.Tracks[c.Y] = row[:c.X] + string(c.Direction) + row[c.X+1:]

	return nil, nil
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
		if a != nil {
			return a, b
		}
	}
	return nil, nil
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

	for _, c := range g.AllCarts {
		fmt.Println(string(c.Direction), c.X, c.Y)
	}
}
