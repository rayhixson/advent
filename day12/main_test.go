package main

import (
	"fmt"
	"testing"
)

func TestGiven(t *testing.T) {
	notes := ReadNotes(test_keys)

	//fmt.Println(notes)

	n := NewRow(test_init, 0)

	fmt.Println("Iniit:", n)
	gen := 20
	for i := 1; i <= gen; i++ {
		n = NextGen(n, &notes)
		fmt.Printf("Gen %d: %s\n", i, n)
	}

	sum, xsum := n.SumLives(0)
	fmt.Println("Total", sum)

	if sum != 325 {
		t.Error("Wrong sum: ", sum, xsum)
	}
}

const test_init = `#..#.#..##......###...###`

const test_keys = `...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #
`
