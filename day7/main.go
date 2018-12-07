package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

const extra_time = 60
const worker_count = 5

type Worker struct {
	id          int
	currentStep *Step
	startTime   int
	willTake    int
}

type Workers []*Worker

// isCompleted checks the clock, if done returns the step
func (w *Worker) isCompleted(clock int) *Step {
	if w.currentStep == nil {
		//fmt.Printf("Worker [%v], idle\n", w.id)
		return nil
	}

	if (clock - w.startTime) >= w.willTake {
		c := w.currentStep
		w.currentStep = nil
		c.done = true
		//fmt.Printf("Worker [%v], done with [%s], in time [%v]\n", w.id, c.id, (clock - w.startTime))
		return c
	}

	//fmt.Printf("Worker [%v], NOT done, in time [%v], should take [%v]\n", w.id, (clock - w.startTime), w.willTake)
	return nil
}

// assign accepts new work if it it's not busy
func (w *Worker) assign(s *Step, clock *int) bool {
	if w.currentStep != nil {
		return false
	}
	w.currentStep = s
	w.currentStep.assigned = true
	w.willTake = int(w.currentStep.id[0]) - 'A' + 1
	w.willTake += extra_time
	w.startTime = *clock

	//fmt.Printf("Worker Assigned [%v][%s], start [%v], willtake [%v]\n", w.id, w.currentStep.id, w.startTime, w.willTake)

	return true
}

func sequenceAndAssignSteps(steps *Steps, completed *Steps, workers *Workers, clock *int) {
	//fmt.Println("Clock:", *clock)
	//fmt.Println("Whole:", steps)

	if *clock > 1000 {
		fmt.Println("Too much time passed, exit")
		return
	}

	// first see if the workers are done
	for _, w := range *workers {
		done := w.isCompleted(*clock)
		if done != nil {
			steps.delete(done)
			*completed = append(*completed, done)
		}
	}

	// find ones that have no parents
	var noParents Steps
	for _, s := range *steps {
		if len(s.parents) == 0 && !s.assigned {
			noParents = append(noParents, s)
		}
	}

	if len(noParents) > 0 {
		sort.Slice(noParents, func(i, j int) bool {
			return noParents[i].id < noParents[j].id
		})

		assignWork(&noParents, workers, clock)
	}

	if len(*steps) > 0 {
		*clock++
		sequenceAndAssignSteps(steps, completed, workers, clock)
	}
	return
}

func assignWork(nexts *Steps, workers *Workers, clock *int) (completed Steps) {
	// try to assign each step
	for _, s := range *nexts {
		assigned := false
		// check each worker for availability
		for i, _ := range *workers {
			assigned = (*workers)[i].assign(s, clock)
			if assigned {
				break
			}
		}
		if !assigned {
			return completed
		}
	}

	return completed
}

func parse(d string) (steps Steps) {
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

	final := Steps{}
	w := Workers{}
	for i := 1; i < worker_count; i++ {
		w = append(w, &Worker{id: i})
	}

	clock := 0
	sequenceAndAssignSteps(&steps, &final, &w, &clock)

	fmt.Println("Time:", clock)

	seq := ""
	for _, s := range final {
		seq += s.id
	}

	fmt.Println("Final sequence:", seq)
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
