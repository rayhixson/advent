package main

type stack []int

func (s stack) Push(c int) stack {
	return append(s, c)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	if l < 1 {
		panic("empty")
	}

	return s[:(l - 1)], s[l-1]
}

func (s stack) Peek() *int {
	l := len(s)
	if l == 0 {
		return nil
	}
	return &s[l-1]
}

func (s stack) Reverse() stack {
	c := 0
	l := len(s)
	if l < 2 {
		return s
	}
	if l%2 > 0 {
		c = (l - 1) / 2
	} else {
		c = l / 2
	}

	for i := 0; i < c; i++ {
		x := s[i]
		s[i] = s[l-i-1]
		s[l-i-1] = x
	}
	return s
}

func (s stack) Print() string {
	var x string
	l := len(s)
	for i := (l - 1); i >= 0; i-- {
		x += string(s[i])
	}
	return x
}
