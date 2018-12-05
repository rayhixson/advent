package main

import "testing"

func TestGiven(t *testing.T) {
	res := reduce("dabAcCaCBAcCcaDA")
	if res != "dabCBAcaDA" {
		t.Errorf("Wrong Given: %v", res)
	}
}

func TestEmpty(t *testing.T) {
	res := reduce("bacbAaBCAB")
	if res != "" {
		t.Errorf("Wrong Empty: %v", res)
	}
}

func TestStack(t *testing.T) {
	s := stack{}
	s = s.Push('a')
	s = s.Push('b')
	s = s.Push('c')

	var x int
	s, x = s.Pop()
	if x != 'c' {
		t.Error("Wrong", x)
	}
	s, x = s.Pop()
	if x != 'b' {
		t.Error("Wrong", x)
	}
	s, x = s.Pop()
	if x != 'a' {
		t.Error("Wrong", x)
	}
}

func TestRemoveUnit(t *testing.T) {
	s := "dabAcCaCBAcCcaDA"

	if shorter := removeAndReduce(s, 'A'); len(shorter) != 6 {
		t.Error("Wrong len for A remove:", len(shorter))
	}
	if shorter := removeAndReduce(s, 'B'); len(shorter) != 8 {
		t.Error("Wrong len for A remove:", len(shorter))
	}
	if shorter := removeAndReduce(s, 'C'); len(shorter) != 4 {
		t.Error("Wrong len for A remove:", len(shorter))
	}
	if shorter := removeAndReduce(s, 'D'); len(shorter) != 6 {
		t.Error("Wrong len for A remove:", len(shorter))
	}
}
