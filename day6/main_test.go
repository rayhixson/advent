package main

import "testing"

func TestGivenAreas(t *testing.T) {
	coords := parse(data)
	g := populateGrid(coords)

	dumpGrid(g)

	for _, c := range coords {
		switch c.id {
		case 1:
			checkArea(t, g, c, -1)
		case 2:
			checkArea(t, g, c, -1)
		case 3:
			checkArea(t, g, c, -1)
		case 4:
			checkArea(t, g, c, 9)
		case 5:
			checkArea(t, g, c, 17)
		case 6:
			checkArea(t, g, c, -1)
		}
	}
}

func checkArea(t *testing.T, g grid, c point, expected int) {
	findArea(&c, g)
	if c.size != expected {
		t.Errorf("Wrong Area %+v, expected: %d\n", c, expected)
	}
}

func TestGivenLargest(t *testing.T) {
	coords := parse(data)

	c, _ := findLargestArea(coords)

	if c.x != 5 && c.y != 5 && c.size != 17 {
		t.Errorf("Wrong largest: %+v\n", c)
	}
}

const data = `
1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
`
