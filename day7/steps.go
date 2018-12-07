package main

import "fmt"

type Step struct {
	id       string
	parents  []*Step
	done     bool
	assigned bool
}

func (s Step) String() string {
	ps := "("
	for _, p := range s.parents {
		ps += p.id
	}
	ps += ")"
	return fmt.Sprintf("%s: parents: %s", s.id, ps)
}

type Steps []*Step

func (ss *Steps) add(newid string, newparent *Step) *Step {
	for i, s := range *ss {
		if s.id == newid {
			if newparent != nil {
				(*ss)[i].parents = append(s.parents, newparent)
			}
			return s
		}
	}

	s := Step{id: newid}
	if newparent != nil {
		s.parents = append(s.parents, newparent)
	}
	*ss = append(*ss, &s)
	return &s
}

func (ss *Steps) delete(dels *Step) {
	index := -1
	for i, s := range *ss {
		if s.id == dels.id {
			index = i
		}
		// remove parent refs
		for j, p := range s.parents {
			if p.id == dels.id {
				(*ss)[i].parents = append(s.parents[:j], s.parents[j+1:]...)
			}
		}
	}

	// remove from array
	if index >= 0 {
		*ss = append((*ss)[:index], (*ss)[index+1:]...)
	}
}
