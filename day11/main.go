package main

import "fmt"

var dim = 300

type Cell struct {
	X          int
	Y          int
	PowerLevel int
	RackID     int
	SumAsHome  int
}

func NewCell(serialNumber int, x, y int) *Cell {
	c := Cell{X: x, Y: y}
	c.RackID = c.X + 10
	p := c.RackID * c.Y
	p += serialNumber
	p *= c.RackID
	p = (p / 100) % 10
	p -= 5
	c.PowerLevel = p
	return &c
}

type Grid [][]*Cell

// sum returns whether it's a viable candidate
func (g Grid) sum(c *Cell, size int) bool {
	// this cells location in the grid
	x := c.X - 1
	y := c.Y - 1

	// if it's at a border then the sum is 0
	if x+(size-1) > dim-1 || y+(size-1) > dim-1 {
		c.SumAsHome = 0
		return false
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			c.SumAsHome += g[x+i][y+j].PowerLevel
		}
	}
	return true
}

func (g Grid) FindMaxHome(size int) *Cell {
	var max *Cell
	for y := 0; y < dim; y++ {
		for x := 0; x < dim; x++ {
			c := g[x][y]
			ok := g.sum(c, size)
			if ok {
				if max == nil || c.SumAsHome > max.SumAsHome {
					max = c
				}
			}
		}
	}

	return max
}

func FindMaxHomeAndSize() (maxCell Cell, size int) {
	for i := 1; i <= 300; i++ {
		g := NewGrid(7511, 300)
		c := g.FindMaxHome(i)

		if c.SumAsHome > maxCell.SumAsHome {
			fmt.Printf("New Max at [%d]: %+v\n", i, maxCell)
			maxCell = *c
			size = i
		}
	}

	return maxCell, size
}

func NewGrid(serialNumber, size int) *Grid {
	g := Grid{}
	for x := 0; x < size; x++ {
		g = append(g, make([]*Cell, dim))
		for y := 0; y < size; y++ {
			c := NewCell(serialNumber, x+1, y+1)
			g[x][y] = c
		}
	}
	return &g
}

func main() {
	g := NewGrid(7511, 300)
	c := g.FindMaxHome(3)
	fmt.Println("Cell with most:", c)

	cell, size := FindMaxHomeAndSize()
	fmt.Printf("X,Y,size: [%d,%d,%d] has power %d\n", cell.X, cell.Y, size, cell.SumAsHome)
}
