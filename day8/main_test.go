package main

import (
	"strings"
	"testing"
)

func TestGiven(t *testing.T) {
	tree := parseTree(test_license)
	r := tree.String()

	if strings.TrimSpace(r) != test_license {
		t.Errorf("Should have matched:\n %s\n and\n %s\n", r, test_license)
	}

	sum := tree.CountMetaData()
	if sum != 138 {
		t.Error("Wrong meta data sum:", sum)
	}
}

/*
func TestHyphen(t *testing.T) {
	tree := parseTree(test_license)
	tree.Dump()

	t.Error("XXX")
}
*/

const test_license = `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`
