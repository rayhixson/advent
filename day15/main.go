package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type Point struct {
	X        int
	Y        int
	Distance int
	IsWall   bool
	Type     rune
	Soldier  *Unit
}

func NewPoint(x, y int, t rune) Point {
	return Point{
		X:        x,
		Y:        y,
		Type:     t,
		Distance: -1,
	}
}

func (p Point) Print(what string) string {
	switch what {
	case "turn":
		if p.Soldier != nil {
			return fmt.Sprintf("%d", p.Soldier.Turn)
		}
	case "distance":
		if p.Distance > -1 {
			return fmt.Sprintf("%d", p.Distance)
		}
	case "battle":
		return fmt.Sprintf("%c(%d, %d)", p.Type, p.X, p.Y)
	}

	return string(p.Type)
}

func (p *Point) MarkDistance(d int) int {
	if p.Type == '.' && p.Distance == -1 {
		p.Distance = d
		return 1
	}
	return 0
}

func (p *Point) Reset() {
	if p.Type == '?' {
		p.Type = '.'
	}
	p.Distance = -1
}

func (ps Points) Dump() {
	for _, p := range ps {
		fmt.Printf("[%v, %v] -> %v\n", p.X, p.Y, p.Distance)
	}
}

func (ps Points) Nearest() Points {
	near := Points{}
	min := 10000000
	for _, p := range ps {
		if p.Distance > -1 && p.Distance < min {
			min = p.Distance
		}
	}

	for _, p := range ps {
		if p.Distance == min {
			near = append(near, p)
		}
	}

	return near
}

func (ps Points) ReadingOrder() {
	sort.Slice(ps, func(i, j int) bool {
		if ps[i].Y == ps[j].Y {
			return ps[i].X < ps[j].X
		}
		return ps[i].Y < ps[j].Y
	})
}

type Grid []Points

func (g Grid) Neighbors(p *Point) Points {
	x, y := p.X, p.Y
	ps := Points{}
	if x > 0 {
		ps = append(ps, g[y][x-1])
	}
	if x < len(g[0]) {
		ps = append(ps, g[y][x+1])
	}
	if y > 0 {
		ps = append(ps, g[y-1][x])
	}
	if y < len(g) {
		ps = append(ps, g[y+1][x])
	}
	return ps
}

type Points []*Point

type Battlefield struct {
	Grid Grid
}

func NewBattlefield() Battlefield {
	return Battlefield{
		Grid: Grid{},
	}
}

func (b Battlefield) Dump(what string) {
	for _, row := range b.Grid {
		s := ""
		for _, v := range row {
			s += v.Print(what)
		}

		// summarize the units
		units := ""
		for _, v := range row {
			if v.Soldier != nil {
				units += fmt.Sprintf("%c(%v) ", v.Type, v.Soldier.HitScore)
			}
		}
		fmt.Printf("%v    %v\n", s, units)
	}
}

func (b Battlefield) Reset() {
	for _, row := range b.Grid {
		for _, v := range row {
			v.Reset()
		}
	}
}

func (b *Battlefield) SumUnits() int {
	sum := 0
	for _, row := range b.Grid {
		for _, v := range row {
			if v.Soldier != nil {
				sum += v.Soldier.HitScore
			}
		}
	}
	return sum
}

func (b *Battlefield) SortedUnits() Units {
	// refresh the list of units
	allUnits := Units{}
	for _, row := range b.Grid {
		for _, v := range row {
			if v.Soldier != nil {
				if v.Soldier.HitScore < 1 {
					panic("Zombie alert!!")
				}

				allUnits = append(allUnits, v.Soldier)
			}
		}
	}

	allUnits.ReadingOrder()
	return allUnits
}

// score the distance out from this point to everywhere in the grid
// marks the numbers
func (b *Battlefield) MarkDistance(p *Point) {
	distance := 0
	p.Distance = distance

	mark := func(dist int) int {
		c := 0
		for _, row := range b.Grid {
			for _, v := range row {
				if v.Distance == dist {
					for _, n := range b.Grid.Neighbors(v) {
						c += n.MarkDistance(dist + 1)
					}
				}
			}
		}
		return c
	}

	count := 1
	for count > 0 {
		count = mark(distance)
		distance++
	}
}

func (b Battlefield) FindEnemyToAttack(turner *Unit) *Unit {
	enemies := Points{}
	for _, n := range b.Grid.Neighbors(turner.Location) {
		if n.Soldier != nil && n.Type != turner.Location.Type {
			// he's an enemy
			enemies = append(enemies, n)
		}
	}

	// return least points
	minHit := 10000
	for _, e := range enemies {
		if e.Soldier.HitScore < minHit {
			minHit = e.Soldier.HitScore
		}
	}

	// find all that match
	options := Points{}
	for _, e := range enemies {
		if e.Soldier.HitScore == minHit {
			options = append(options, e)
		}
	}

	options.ReadingOrder()
	if len(options) > 0 {
		return options[0].Soldier
	}
	return nil
}

func (b Battlefield) AnyEnemiesLeft(allUnits Units, u *Unit) bool {
	for _, v := range allUnits {
		if v.HitScore > 0 && v.Location.Type != u.Location.Type {
			return true
		}
	}
	return false
}

// return target to approach
func (b Battlefield) FindTarget(turner *Unit) *Point {
	b.Reset()

	// identify all the 'in range' target spots we could move to
	rangers := Points{}
	for _, row := range b.Grid {
		for _, v := range row {
			if v.Soldier != nil && v.Soldier != turner && v.Type != turner.Location.Type {
				for _, n := range b.Grid.Neighbors(v) {
					if n.Type == '.' {
						n.Type = '?'
						rangers = append(rangers, n)
					}
				}
			}
		}
	}

	b.Reset()
	b.MarkDistance(turner.Location)
	rangers = rangers.Nearest()
	if len(rangers) == 0 {
		return nil
	}

	rangers.ReadingOrder()
	return rangers[0]
}

func (b Battlefield) FindPathStep(p *Point, target *Point) *Point {
	b.Reset()

	b.MarkDistance(target)
	neighbors := b.Grid.Neighbors(p)
	options := neighbors.Nearest()
	options.ReadingOrder()
	if len(options) > 0 {
		return options[0]
	}
	return nil
}

func (b *Battlefield) Move(u *Unit, next *Point) {
	next.Type = u.Location.Type
	u.Location.Type = '.'

	u.Location.Soldier = nil
	next.Soldier = u
	u.Location = next
}

func (b Battlefield) BuryTheDead(u *Unit) {
	if u.HitScore <= 0 {
		if u.Location.Type == 'E' {
			panic("Elf died!!!")
		}
		u.Location.Type = '.'
		u.Location.Soldier = nil
		u.Location = nil
	}
}

// sort all units by turn order then stick with it for the whole round
// each live unit takes a turn()
// returns false if no targets are found by a unit
func (b Battlefield) GoRound() bool {
	allUnits := b.SortedUnits()

	for _, turner := range allUnits {
		if turner.HitScore <= 0 {
			//fmt.Println("Skipping dead: ", turner)
			continue
		}

		if !b.AnyEnemiesLeft(allUnits, turner) {
			// this ends the battle
			return false
		}

		enemy := b.FindEnemyToAttack(turner)

		if enemy == nil {
			target := b.FindTarget(turner)
			if target == nil {
				// we can't move
				continue
			}

			nextstep := b.FindPathStep(turner.Location, target)
			if nextstep != nil {
				b.Move(turner, nextstep)

				// then try for an enemy again
				enemy = b.FindEnemyToAttack(turner)
			}
		}

		if enemy != nil {
			turner.Fight(enemy)
			b.BuryTheDead(enemy)
		}
	}

	return true
}

func (b *Battlefield) Battle() (lastFullRound int, sumPoints int) {
	max := 200
	for i := 1; i < max; i++ {
		//fmt.Println("Round: ", i)
		if !b.GoRound() {
			//fmt.Println("Final Round: ", i)
			lastFullRound = i - 1
			break
		}
		//b.Dump("")
	}

	return lastFullRound, b.SumUnits()
}

func parse(data string) *Battlefield {
	b := NewBattlefield()
	scanner := bufio.NewScanner(strings.NewReader(data))
	row := 0
	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {
			continue
		}

		colOfPoints := Points{}
		for col, c := range val {
			p := NewPoint(col, row, c)
			colOfPoints = append(colOfPoints, &p)
			switch c {
			case 'E', 'G':
				u := NewUnit(col, row, &p)
				p.Soldier = u
			}
		}
		b.Grid = append(b.Grid, colOfPoints)
		row++
	}
	return &b
}

func main() {
	b := parse(input)
	//b.Dump("")
	lastFullRound, sum := b.Battle()
	//b.Dump("")

	fmt.Printf("Final Round * Sum: %v * %v = %v\n", lastFullRound, sum, lastFullRound*sum)
}

const input = `
################################
##########G###.#################
##########..G#G.G###############
##########G......#########...###
##########...##.##########...###
##########...##.#########G..####
###########G....######....######
#############...............####
#############...G..G.......#####
#############.............######
############.............E######
######....G..G.........E....####
####..G..G....#####.E.G.....####
#####...G...G#######........####
#####.......#########........###
####G.......#########.......####
####...#....#########.#.....####
####.#..#...#########E#..E#..###
####........#########..E.#######
###......#..G#######....########
###.......G...#####.....########
##........#............#########
#...##.....G......E....#########
#.#.###..#.....E.......###.#####
#######................###.#####
##########.......E.....###.#####
###########...##........#...####
###########..#####.............#
############..#####.....#......#
##########...######...........##
#########....######..E#....#####
################################`
