package main

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	b := parse(simple)
	b.SortedUnits()

	b.Dump("default")
	b.Dump("turn")

	t.Fail()
}

func TestDistance(t *testing.T) {
	b := parse(inrange)
	n := b.SortedUnits()[0]

	b.MarkDistance(n.Location)
	b.Dump("distance")

	t.Fail()
}

func TestMove(t *testing.T) {
	b := parse(dest)

	n := b.SortedUnits()[0]

	b.Dump("")
	var enemy *Unit
	for i := 0; i < 5; i++ {
		enemy = b.FindEnemyToAttack(n)

		if enemy == nil {
			target := b.FindTarget(n)
			nextstep := b.FindPathStep(n.Location, target)
			b.Move(n, nextstep)
		} else {
			break
		}
		b.Dump("")
	}
	if enemy != nil {
		fmt.Printf("Battle: %v vs %v\n", n.Location.Print("battle"), enemy.Location.Print("battle"))
	} else {
		fmt.Println("Dead end")
	}

	t.Fail()
}

func TestBattle(t *testing.T) {
	b := parse(fight)
	b.SortedUnits()

	b.Dump("")

	b.GoRound()
	b.Dump("")

	b.GoRound()
	b.Dump("")

	b.GoRound()
	b.Dump("")

	t.Fail()
}

func TestFinal(t *testing.T) {
	b := parse(fight)

	for i := 1; i < 100; i++ {
		fmt.Println("Round: ", i)
		if !b.GoRound() {
			fmt.Println("Final Round: ", i)
			break
		}
		b.Dump("")
	}

	t.Fail()
}

func TestA(t *testing.T) {
	b := parse(fight)
	b.Dump("")
	lastFullRound, sum := b.Battle()
	b.Dump("")

	fmt.Printf("Final Round * Sum: %v * %v = %v\n", lastFullRound, sum, lastFullRound*sum)

	t.Fail()
}

func TestB(t *testing.T) {
	b := parse(B)
	b.Dump("")
	lastFullRound, sum := b.Battle()
	b.Dump("")

	fmt.Printf("Final Round * Sum: %v * %v = %v\n", lastFullRound, sum, lastFullRound*sum)

	t.Fail()
}

const B = `
#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`

const fight = `
#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`

const dest = `
#######
#.E...#
#.....#
#...G.#
#######`

const inrange = `
#######
#E..G.#
#...#.#
#.G.#G#
#######`

const simple = `
#######
#.G.E.#
#E.G.E#
#.G.E.#
#######`
