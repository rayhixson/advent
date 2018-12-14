package main

import (
	"fmt"
	"strings"
)

// map left,up,down,right to what would be underneath that pattern
func Reveal(tomatch string) rune {
	val, ok := coveredSpotPatterns[tomatch]

	/*
		for k, v := range CoveredSpotPatterns {
			b, err := regexp.MatchString(tomatch, k)
			if err != nil {
				panic(fmt.Sprintf("Bad regexp: %s, %v", tomatch, k))
			}
			if b {
				return v
			}
		}
	*/

	if !ok {
		panic(fmt.Sprintf("Illegal covered spot state: [%s]\n", tomatch))
	}
	return val
}

var coveredSpotPatterns = map[string]rune{
	"-  -":  '-',
	" || ":  '|',
	" \\+ ": '|',
	"-  \\": '-',
	"-|/ ":  '+',
	"- | ":  '\\',
	"-|/-":  '+',
	"+  -":  '-',
	"-  /":  '-',
	"-||-":  '+',
	"-|  ":  '/',
	" |/ ":  '|',
	"-  +":  '-',
	" \\| ": '|',
	"-/|-":  '+',
	"  +-":  '/',
	"/ --":  '-',
	">  -":  '-',
	"+- \\": '-',
	"--| ":  '\\',
	"- \\-": '+',
	"/ -+":  '-',
	"\\ -+": '-',
	" v--":  '\\',
	"-- +":  '-',
	" +||":  '|',
	"-- -":  '-',
}

type Grid struct {
	Tracks   []string
	AllCarts Carts
}

func (g Grid) UncoverPoint(x, y int) rune {
	//fmt.Println("Looking under:", x, y)
	// check for boarders
	left, right, up, down := " ", " ", " ", " "
	if x > 0 {
		left = string(g.Tracks[y][x-1])
	}
	if x < len(g.Tracks[0])-1 {
		right = string(g.Tracks[y][x+1])
	}
	if y > 0 {
		up = string(g.Tracks[y-1][x])
	}
	if y < len(g.Tracks)-1 {
		down = string(g.Tracks[y+1][x])
	}
	key := left + up + down + right
	s := Reveal(key)
	//fmt.Printf("Found: %s => [%s]\n", key, string(s))
	return s
}

func (g Grid) String() string {
	return strings.Join(g.Tracks, "\n")
}

func (g *Grid) NextTrack(x, y int) byte {
	//fmt.Println("Trying to read:", c.X, nextY)
	//fmt.Println("X:", g.Tracks)
	return g.Tracks[y][x]
}

func (g *Grid) UpdateMove(c *Cart, xdif, ydif int) {
	// set where the cart was to should be underneath
	row := g.Tracks[c.Y]
	g.Tracks[c.Y] = row[:c.X] + string(g.UncoverPoint(c.X, c.Y)) + row[c.X+1:]

	// set the cart to the new location
	c.X += xdif
	c.Y += ydif

	// redraw the cart in the new location
	row = g.Tracks[c.Y]
	g.Tracks[c.Y] = row[:c.X] + string(c.Direction) + row[c.X+1:]
}

func (g *Grid) Click() {
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
		g.UpdateMove(c, nextX, nextY)
	}
}

// run simulation and return location and tick of first collision
func (g *Grid) Run(maxTicks int) (tick int, colliders Carts) {
	for i := 0; i < maxTicks; i++ {
		g.Click()
		//dump(g)
		collidings := g.AllCarts.Collisions()
		if len(collidings) > 0 {
			return i, collidings
		}
	}

	return tick, colliders
}

func dump(g *Grid) {
	fmt.Println(g)

	for _, c := range g.AllCarts {
		fmt.Println(string(c.Direction), c.X, c.Y)
	}
}
