package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	id   int
	x    int
	y    int
	size int
}

// every x and y with an id of the closest point
type grid [][]int

func findArea(c *point, g grid) {
	xedge := len(g) - 1
	yedge := len(g[xedge]) - 1
	for x, ax := range g {
		for y, v := range ax {
			if v == c.id {
				// if it's a border point then it's infinite
				if x == xedge || y == yedge || x == 0 || y == 0 {
					c.size = -1
					return
				}
				c.size++
			}
		}
	}
}

func findLargestArea(coords []point) (point, grid) {
	g := populateGrid(coords)

	big := point{}
	for _, c := range coords {
		findArea(&c, g)
		if c.size > big.size {
			big = c
		}
	}
	return big, g
}

func findClosestID(x, y int, g grid, coords []point) (coordID int) {
	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		}
		return x
	}

	// find the distance from this one point to every coord
	var minCoordA point
	var minCoordB point
	minDisA := 9999999999
	minDisB := 9999999999
	for _, c := range coords {
		d := abs(x-c.x) + abs(y-c.y)
		if d < minDisA {
			minDisA = d
			minCoordA = c
		} else {
			if d == minDisA {
				minDisB = d
				minCoordB = c
			}
		}
	}

	if minDisA == minDisB && minCoordA.id != minCoordB.id {
		// two equally distant coords - return not a valid id
		return -1
	}
	return minCoordA.id
}

func populateGrid(coords []point) grid {
	maxX, maxY := 0, 0
	for _, c := range coords {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	// empty grid
	g := grid{}
	for i := 0; i < (maxX + 1); i++ {
		g = append(g, []int{})
		for j := 0; j < (maxY + 1); j++ {
			g[i] = append(g[i], -1)
		}
	}
	fmt.Println("Grid size:", len(g), len(g[len(g)-1]))

	// place the coords
	for _, c := range coords {
		g[c.x][c.y] = c.id
	}

	// label the nearest to each point
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[len(g)-1]); j++ {
			g[i][j] = findClosestID(i, j, g, coords)
		}
	}

	return g
}

func dumpGrid(g grid) {
	for y := 0; y < len(g[0]); y++ {
		for x, _ := range g {
			c := g[x][y]
			if c == -1 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", c)
			}
		}
		fmt.Println()
	}
}

func parse(data string) (coords []point) {
	scanner := bufio.NewScanner(strings.NewReader(data))
	for i := 0; scanner.Scan(); i++ {
		val := scanner.Text()
		if val != "" {
			parts := strings.Split(val, ",")
			c := point{
				id: i,
				x:  num(parts[0]),
				y:  num(parts[1]),
			}
			coords = append(coords, c)
		}
	}
	return coords
}

func num(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	coords := parse(data)

	c, _ := findLargestArea(coords)

	fmt.Printf("Largest: %+v\n", c)
}

const data = `
118, 274
102, 101
216, 203
208, 251
309, 68
330, 93
91, 179
298, 278
201, 99
280, 272
141, 312
324, 290
41, 65
305, 311
198, 68
231, 237
164, 224
103, 189
216, 207
164, 290
151, 91
166, 250
129, 149
47, 231
249, 100
262, 175
299, 237
62, 288
228, 219
224, 76
310, 173
80, 46
312, 65
183, 158
272, 249
57, 141
331, 191
163, 359
271, 210
142, 137
349, 123
55, 268
160, 82
180, 70
231, 243
133, 353
246, 315
164, 206
229, 97
268, 94
`
