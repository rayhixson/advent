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

func TestNodeValue(t *testing.T) {
	tree := parseTree(test_license)
	v := tree.NodeValue()

	if v != 66 {
		t.Error("Wrong node value:", v)
	}
}

const test_license = `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`
