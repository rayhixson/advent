package main

import "testing"

func TestBasic(t *testing.T) {
	data := []int{-4, +9, +6}
	shouldEqual := 11
	n := sum(data)

	if n != shouldEqual {
		t.Errorf("Result should have been %d, not %d", shouldEqual, n)
	}
}

func TestMulti(t *testing.T) {
	data := []int{-400, +900 + 600}

	shouldEqual := 1100
	n := sum(data)
	if n != shouldEqual {
		t.Errorf("Result should have been %d, not %d", shouldEqual, n)
	}
}

func TestFreq(t *testing.T) {
	compareFreq(0, t, +1, -1)
	compareFreq(10, t, +3, +3, +4, -2, -4)
	compareFreq(5, t, -6, +3, +8, +5, -6)
	compareFreq(14, t, +7, +7, -2, -7, -4)
}

func compareFreq(result int, t *testing.T, seq ...int) {
	repeater := findRepeat(seq)
	if repeater != result {
		t.Errorf("Failed to find %v, got %v, given: %v", result, repeater, seq)
	}
}
