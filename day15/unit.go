package main

import "sort"

type Unit struct {
	HitScore    int
	AttackPower int
	Turn        int
	Location    *Point
}

const ElfPower = 13

func NewUnit(x, y int, where *Point) *Unit {
	u := Unit{
		HitScore:    200,
		AttackPower: 3,
		Location:    where,
	}

	if where.Type == 'E' {
		u.AttackPower = ElfPower
	}

	return &u
}

type Units []*Unit

func (us Units) ReadingOrder() {
	sort.Slice(us, func(i, j int) bool {
		if us[i].Location.Y == us[j].Location.Y {
			return us[i].Location.X < us[j].Location.X
		}
		return us[i].Location.Y < us[j].Location.Y
	})
}

func (u Unit) IdentifyTargets(b *Battlefield) *Units {
	// identify "in range squares" open squares adjacent to targets

	// no targets ends combat
	return nil
}

// Move returns the unit that it can attack, otherwise it didn't move or didn't get to a target
func (u *Unit) Move(b *Battlefield, targets *Units) *Unit {
	// find the shortest path to every target

	// if a target has multiple short paths, choose by reading order

	// if already adjacent then no move

	// if it can't get to a target then don't return anything
	return nil
}

func (u Unit) Find(target *Unit, b *Battlefield) {
	// identify all in range squars for this target

}

func (u *Unit) Fight(enemy *Unit) {
	enemy.HitScore -= u.AttackPower
}
