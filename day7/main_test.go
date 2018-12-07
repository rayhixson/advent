package main

import (
	"fmt"
	"testing"
)

func TestGiven(t *testing.T) {
	steps := parse(tdata)

	final := Steps{}
	w := Workers{&Worker{id: 1}, &Worker{id: 2}}
	clock := 0
	sequenceAndAssignSteps(&steps, &final, &w, &clock)

	fmt.Println("Time:", clock)

	fmt.Println(final)
	seq := ""
	for _, s := range final {
		seq += s.id
	}
	if seq != "CABFDE" {
		t.Error("Wraong ", final)
	}
}

const tdata = `
Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
`
