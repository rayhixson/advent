package main

import (
	"testing"
)

func TestGiven(t *testing.T) {
	g := parse(test_data)

	ticks := 15
	click, colliders := g.Run(ticks)

	if len(colliders) == 0 {
		t.Error("DIdn't find collider in n ticks:", ticks)
		return
	}
	if len(colliders) != 2 {
		t.Error("Wrong:", colliders)
		return
	}

	if colliders[0].X != 7 && colliders[0].Y != 3 {
		t.Error("Wrong Collision at:", colliders[0].X, colliders[0].Y, click)
	}
}

const test_data = `
/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `
