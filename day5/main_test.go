package main

import "testing"

func TestGiven(t *testing.T) {
	res := extract("dabAcCaCBAcCcaDA")
	if res != "dabCBAcaDA" {
		t.Errorf("Wrong Given: %v", res)
	}
}

func TestEmpty(t *testing.T) {
	res := extract("bacbAaBCAB")
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
