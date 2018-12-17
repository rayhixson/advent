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
	count := findSequence("51589")

	if count != 5 {
		t.Error("Wrong:", count)
	}
}
