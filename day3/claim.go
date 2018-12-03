package main

import (
	"strconv"
	"strings"
)

type Claim struct {
	id      int
	xOffset int
	yOffset int
	width   int
	height  int
}

func ParseClaim(line string) Claim {
	c := Claim{}

	parts := strings.Split(line, " ")
	c.id = num(parts[0][1:])
	xy := strings.Split(parts[2], ",")
	c.xOffset = num(xy[0])
	c.yOffset = num(xy[1][:len(xy[1])-1])
	dim := strings.Split(parts[3], "x")
	c.width = num(dim[0])
	c.height = num(dim[1])

	return c
}

func (c Claim) findOverlaps(other Claim, points *PointsMap) {
	for i := 0; i < c.width; i++ {
		xval := i + c.xOffset
		for j := 0; j < c.height; j++ {
			yval := j + c.yOffset
			if other.contains(xval, yval) {
				(*points)[xval][yval] += 1
			}
		}
	}
}

func (c Claim) contains(x, y int) bool {
	if x >= c.xOffset && x < c.xOffset+c.width &&
		y >= c.yOffset && y < c.yOffset+c.height {
		return true
	}
	return false
}

func num(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}
