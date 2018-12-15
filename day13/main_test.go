package main

import (
	"testing"
)

func TestGiven(t *testing.T) {
	g := parse(test_data)

	ticks := 15
	click, a, _ := g.Run(ticks)

	if a == nil {
		t.Error("DIdn't find collider in n ticks:", ticks)
		return
	}

	if a.X != 7 && a.Y != 3 {
		t.Error("Wrong Collision at:", a.X, a.Y, click)
	}
}

const test_data = `
/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `
