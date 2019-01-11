package main

import (
	"testing"
)

func TestGiven(t *testing.T) {
	_, s := find10After(9)

	if s != "5158916779" {
		t.Error("Wrong:", s)
	}

	_, s = find10After(5)
	if s != "0124515891" {
		t.Error("Wrong:", s)
	}

	_, s = find10After(18)
	if s != "9251071085" {
		t.Error("Wrong:", s)
	}

	_, s = find10After(2018)
	if s != "5941429882" {
		t.Error("Wrong:", s)
	}
}

func TestPart2(t *testing.T) {
	count := findSequence(5, 1, 5, 8, 9)

	if count != 9 {
		t.Error("Wrong:", count)
	}

	count = findSequence(0, 1, 2, 4, 5)

	if count != 5 {
		t.Error("Wrong:", count)
		panic("not 5")
	}

	count = findSequence(9, 2, 5, 1, 0)

	if count != 18 {
		t.Error("Wrong:", count)
		panic("not 18")
	}

	count = findSequence(5, 9, 4, 1, 4)

	if count != 2018 {
		t.Error("Wrong:", count)
		panic("not 2018")
	}

}
