package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
type Key struct {
	Note string
}

type Keys []*Key
*/

type Row struct {
	Pots       []rune
	ZeroOffset int
}

func (r *Row) Prune() {
	for r.Pots[0] == '.' {
		r.Pots = r.Pots[1:]
		r.ZeroOffset++
	}
	for r.Pots[len(r.Pots)-1] == '.' {
		r.Pots = r.Pots[:len(r.Pots)-1]
	}
}

func (r Row) String() string {
	header := ""
	for i := r.ZeroOffset; i < len(r.Pots); i++ {
		if i%10 == 0 {
			n := i / 10
			if n < 0 {
				n *= -1
			}
			header += strconv.Itoa(n)
		} else {
			header += " "
		}
	}
	s := "\n"
	s += fmt.Sprintln(header)
	s += fmt.Sprintln(string(r.Pots))
	return s
}

func (r Row) SumLives() int {
	sum := 0
	for i, p := range r.Pots {
		if p == '#' {
			x := i + r.ZeroOffset
			//fmt.Println("Val:", x)
			sum += x
		}
	}
	return sum
}

func NewRow(s string, offset int) Row {
	// grow the pots each time
	p := "...." + s + "...."
	return Row{
		Pots:       []rune(p),
		ZeroOffset: -4 + offset,
	}
}

type Notes map[string]rune

func ReadNotes(data string) Notes {
	notes := make(Notes)

	var key string
	var nextState rune
	rdr := strings.NewReader(data)
	for {
		n, err := fmt.Fscanf(rdr, "%5s => %1c\n", &key, &nextState)
		if n == 0 {
			return notes
		}
		if err != nil {
			panic(err)
		}
		notes[key] = nextState
	}
}

func willLive(set []rune, keys *Notes) rune {
	v, ok := (*keys)[string(set)]
	if ok {
		return v
	}
	return '.'
}

func NextGen(row Row, keys *Notes) Row {
	// copy the row to make it bigger
	prevrow := NewRow(string(row.Pots), row.ZeroOffset)
	// we'll modify this one
	next := NewRow(string(row.Pots), row.ZeroOffset)
	for i := 0; i < (len(prevrow.Pots) - 5); i++ {
		x := prevrow.Pots[i : i+5]
		//lfmt.Println("Checking: ", i, string(x), string(willLive(x, keys)))
		next.Pots[i+2] = willLive(x, keys)
	}
	//fmt.Println(next.ZeroOffset, next)
	next.Prune()
	return next
}

func main() {
	notes := ReadNotes(keys)
	n := NewRow(initialState, 0)
	//gen := 50000000000
	gen := 20
	for i := 1; i <= gen; i++ {
		n = NextGen(n, &notes)
	}
	fmt.Printf("Gen %d: %s\n", gen, n)

	fmt.Println("Total", n.SumLives())
}

const initialState = `##.###.......#..#.##..#####...#...#######....##.##.##.##..#.#.##########...##.##..##.##...####..####`

const keys = `#.#.# => #
.##.. => .
#.#.. => .
..### => #
.#..# => #
..#.. => .
####. => #
###.. => #
#.... => .
.#.#. => #
....# => .
#...# => #
..#.# => #
#..#. => #
.#... => #
##..# => .
##... => .
#..## => .
.#.## => #
.##.# => .
#.##. => #
.#### => .
.###. => .
..##. => .
##.#. => .
...## => #
...#. => .
..... => .
##.## => .
###.# => #
##### => #
#.### => .
`
