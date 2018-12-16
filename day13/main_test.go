package main

import (
	"fmt"
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

func TestCrash(t *testing.T) {
	g := parse(crash_test)

	dump(g)
	ticks := 5
	_, a := g.Run2(ticks)

	fmt.Println(a)
	if a.X != 6 || a.Y != 4 {
		t.Error("Wrong", a)
	}
}

const test_data = `
/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `

const crash_test = `
/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`
