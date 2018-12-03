package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	data := "#1346 @ 284,9: 26x12"

	c := ParseClaim(data)
	if c.id != 1346 ||
		c.xOffset != 284 ||
		c.yOffset != 9 ||
		c.width != 26 ||
		c.height != 12 {
		t.Errorf("Unexpected parse of %v --> %+v", data, c)
	}
}

/*
func TestDisplay(t *testing.T) {
	//data := "#1 @ 1,3: 4x4" //, "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	data := "#2 @ 3,1: 4x4" //, "#3 @ 5,5: 2x2"}

	c := ParseClaim(data)

	for i := 0; i < c.width; i++ {
		xval := i + c.xOffset
		for j := 0; j < c.height; j++ {
			yval := j + c.yOffset

			fmt.Printf("%d,%d\n", xval, yval)
		}
	}
	t.Error("no")
}
*/

func TestOverlaps(t *testing.T) {
	a := Claim{
		xOffset: 0,
		yOffset: 0,
		width:   2,
		height:  2,
	}

	b := Claim{
		xOffset: 1,
		yOffset: 1,
		width:   1,
		height:  1,
	}

	c := Claim{
		xOffset: 1,
		yOffset: 1,
		width:   1,
		height:  1,
	}

	d := Claim{
		xOffset: 1,
		yOffset: 1,
		width:   2,
		height:  2,
	}

	e := Claim{
		xOffset: 0,
		yOffset: 0,
		width:   3,
		height:  3,
	}

	test := func(test int, a, b Claim, expect int) {
		points := PointsMap{}
		a.findOverlaps(b, &points)
		count := points.countNonEmpty()
		if expect != count {
			t.Errorf("Test %d: %d found, expected %d", test, count, expect)
		}
	}
	test(1, a, b, 1)
	test(2, b, a, 1)
	test(3, b, c, 1)
	test(4, c, b, 1)
	test(5, c, d, 1)
	test(6, e, d, 4)
}

func TestContains(t *testing.T) {
	c := ParseClaim("#1 @ 1,3: 4x4")

	test := func(x, y int, within bool) {
		if c.contains(x, y) != within {
			t.Errorf("Wrong: %d,%d", x, y)
		}
	}

	test(3, 1, false)
	test(3, 2, false)
	test(3, 3, true)
	test(3, 4, true)
	test(4, 1, false)
	test(4, 2, false)
	test(4, 3, true)
	test(4, 4, true)
	test(5, 1, false)
	test(5, 2, false)
	test(5, 3, false)
	test(5, 4, false)
	test(6, 1, false)
	test(6, 2, false)
	test(6, 3, false)
	test(6, 4, false)
}

func TestGiven(t *testing.T) {
	data := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}

	claims := []Claim{}
	for _, c := range data {
		claims = append(claims, ParseClaim(c))
	}

	points := PointsMap{}
	findOverlappingSquares(claims, &points)

	count := points.countNonEmpty()
	if count != 4 {
		t.Errorf("Incorrect overlapping squares found: %d, should have been 4", count)
	}
}
