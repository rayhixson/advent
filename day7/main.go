package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type Step struct {
	id      string
	parents []*Step
}

func (s Step) String() string {
	ps := "("
	for _, p := range s.parents {
		ps += p.id
	}
	ps += ")"
	return fmt.Sprintf("%s: parents: %s", s.id, ps)
}

type Steps []Step

func (ss *Steps) add(newid string, newparent *Step) *Step {
	for i, s := range *ss {
		if s.id == newid {
			if newparent != nil {
				(*ss)[i].parents = append(s.parents, newparent)
			}
			return &s
		}
	}

	s := Step{id: newid}
	if newparent != nil {
		s.parents = append(s.parents, newparent)
	}
	*ss = append(*ss, s)
	return &s
}

func (ss Steps) delete(dels Step) Steps {
	index := -1
	for i, s := range ss {
		if s.id == dels.id {
			index = i
		}
		// remove parent refs
		for j, p := range s.parents {
			if p.id == dels.id {
				ss[i].parents = append(s.parents[:j], s.parents[j+1:]...)
			}
		}
	}

	// remove from array
	if index >= 0 {
		return append(ss[:index], ss[index+1:]...)
	}
	return ss
}

func determineSequence(steps Steps, sequence string) string {
	noParents := func() Steps {
		var ts Steps
		for _, s := range steps {
			if len(s.parents) == 0 {
				ts = append(ts, s)
			}
		}
		return ts
	}

	nexts := noParents()
	if len(nexts) == 0 {
		return sequence
	}

	sort.Slice(nexts, func(i, j int) bool {
		return nexts[i].id < nexts[j].id
	})
	sequence += nexts[0].id
	steps = steps.delete(nexts[0])

	if len(steps) > 0 {
		return determineSequence(steps, sequence)
	}
	return sequence
}

func parse(d string) []Step {
	steps := Steps{}

	scanner := bufio.NewScanner(strings.NewReader(d))
	for scanner.Scan() {
		val := scanner.Text()
		if val != "" {
			id := val[36:37]
			parent := val[5:6]
			// make sure parent is an object in the network
			s := steps.add(parent, nil)
			// and the id in question
			s = steps.add(id, s)
		}
	}

	return steps
}

func main() {
	steps := parse(data)

	final := determineSequence(steps, "")
	fmt.Println("Instruction sequence:", final)
}

const data = `
Step P must be finished before step R can begin.
Step V must be finished before step J can begin.
Step O must be finished before step K can begin.
Step S must be finished before step W can begin.
Step H must be finished before step E can begin.
Step K must be finished before step Y can begin.
Step B must be finished before step Z can begin.
Step N must be finished before step G can begin.
Step W must be finished before step I can begin.
Step L must be finished before step Y can begin.
Step U must be finished before step Q can begin.
Step R must be finished before step Z can begin.
Step Z must be finished before step E can begin.
Step C must be finished before step I can begin.
Step I must be finished before step Q can begin.
Step D must be finished before step E can begin.
Step A must be finished before step J can begin.
Step G must be finished before step Y can begin.
Step M must be finished before step T can begin.
Step E must be finished before step X can begin.
Step F must be finished before step T can begin.
Step X must be finished before step J can begin.
Step Y must be finished before step J can begin.
Step T must be finished before step Q can begin.
Step J must be finished before step Q can begin.
Step E must be finished before step Y can begin.
Step A must be finished before step T can begin.
Step P must be finished before step H can begin.
Step W must be finished before step R can begin.
Step Y must be finished before step Q can begin.
Step W must be finished before step M can begin.
Step O must be finished before step M can begin.
Step H must be finished before step R can begin.
Step N must be finished before step L can begin.
Step V must be finished before step W can begin.
Step S must be finished before step Q can begin.
Step D must be finished before step J can begin.
Step W must be finished before step E can begin.
Step V must be finished before step Y can begin.
Step O must be finished before step C can begin.
Step B must be finished before step T can begin.
Step W must be finished before step T can begin.
Step G must be finished before step T can begin.
Step D must be finished before step T can begin.
Step P must be finished before step E can begin.
Step P must be finished before step J can begin.
Step G must be finished before step E can begin.
Step Z must be finished before step M can begin.
Step K must be finished before step T can begin.
Step H must be finished before step U can begin.
Step P must be finished before step T can begin.
Step W must be finished before step A can begin.
Step A must be finished before step F can begin.
Step F must be finished before step Y can begin.
Step H must be finished before step M can begin.
Step T must be finished before step J can begin.
Step O must be finished before step S can begin.
Step P must be finished before step M can begin.
Step X must be finished before step T can begin.
Step S must be finished before step J can begin.
Step H must be finished before step C can begin.
Step B must be finished before step W can begin.
Step K must be finished before step N can begin.
Step E must be finished before step T can begin.
Step S must be finished before step Y can begin.
Step C must be finished before step G can begin.
Step R must be finished before step D can begin.
Step N must be finished before step U can begin.
Step O must be finished before step L can begin.
Step B must be finished before step F can begin.
Step S must be finished before step F can begin.
Step X must be finished before step Y can begin.
Step S must be finished before step D can begin.
Step R must be finished before step E can begin.
Step S must be finished before step A can begin.
Step S must be finished before step X can begin.
Step A must be finished before step G can begin.
Step E must be finished before step F can begin.
Step P must be finished before step A can begin.
Step A must be finished before step M can begin.
Step E must be finished before step Q can begin.
Step H must be finished before step W can begin.
Step W must be finished before step U can begin.
Step F must be finished before step Q can begin.
Step I must be finished before step J can begin.
Step H must be finished before step G can begin.
Step I must be finished before step G can begin.
Step P must be finished before step X can begin.
Step I must be finished before step D can begin.
Step R must be finished before step X can begin.
Step S must be finished before step I can begin.
Step Y must be finished before step T can begin.
Step R must be finished before step G can begin.
Step I must be finished before step X can begin.
Step B must be finished before step D can begin.
Step X must be finished before step Q can begin.
Step F must be finished before step X can begin.
Step V must be finished before step R can begin.
Step C must be finished before step J can begin.
Step L must be finished before step Q can begin.
Step K must be finished before step B can begin.
`
