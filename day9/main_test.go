package main

import "testing"

func TestGiven(t *testing.T) {

	check := func(playerCount, lastMarble, highScoreExpect int) {
		g := NewGameBoard(playerCount, lastMarble)
		g.Play()
		if g.HighScore() != highScoreExpect {
			t.Errorf("Wrong high score: %+v\n", g.HighScore())
		}
	}

	check(9, 25, 32)
	check(10, 1618, 8317)
	check(13, 7999, 146373)
	check(17, 1104, 2764)
	check(21, 6111, 54718)
	check(30, 5807, 37305)
}
