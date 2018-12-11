package main

import (
	"testing"
)

func TestCells(t *testing.T) {

	check := func(serial, x, y, power int) {
		c := NewCell(serial, x, y)
		if c.PowerLevel != power {
			t.Error("Wrong Power:", c)
		}
	}
	check(8, 3, 5, 4)
	check(57, 122, 79, -5)
	check(39, 217, 196, 0)
	check(71, 101, 153, 4)
}

func TestGiven(t *testing.T) {
	check := func(serial, x, y, power int) {
		g := NewGrid(serial, 300)
		c := g.FindMaxHome(3)
		if c.SumAsHome != power || c.X != x || c.Y != y {
			t.Errorf("Wrong cell [%+v], for serial %d", c, serial)
		}
	}

	check(18, 33, 45, 29)
	check(42, 21, 61, 30)
}

func TestGiven2(t *testing.T) {
	check := func(serial, x, y, power, size int) {
		g := NewGrid(serial, 300)
		c := g.FindMaxHome(size)
		if c.SumAsHome != power || c.X != x || c.Y != y {
			t.Errorf("Wrong cell [%+v], for serial %d", c, serial)
		}
	}

	check(18, 90, 269, 113, 16)
	check(42, 232, 251, 119, 12)
	check(7511, 235, 288, 147, 13)
}

/*
func TestWhatIThink(t *testing.T) {
	g := NewGrid(7511, 300)
	c := g.FindMaxHome(16)
	fmt.Printf("%+v\n", c)
	c = g.FindMaxHome(17)
	fmt.Printf("%+v\n", c)
	c = g.FindMaxHome(18)
	fmt.Printf("%+v\n", c)
	t.Error("<-------->")
}

func dump(g *Grid) {
	dim = len(*g)
	for y := 0; y < dim; y++ {
		for x := 0; x < dim; x++ {
			c := (*g)[x][y]
			g.sum(c, 3)
			fmt.Printf("%3d ", c.SumAsHome)
		}
		fmt.Println()
	}
}
*/
