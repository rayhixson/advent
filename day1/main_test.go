package main

import "testing"

func TestBasic(t *testing.T) {
	data := `
-4
+9
+6
`
	shouldEqual := 11
	n, err := sum(data)
	if err != nil {
		t.Error(err)
	}
	if n != shouldEqual {
		t.Errorf("Result should have been %d, not %d", shouldEqual, n)
	}
}

func TestMulti(t *testing.T) {
	data := `
-400
+900
+600
`
	shouldEqual := 1100
	n, err := sum(data)
	if err != nil {
		t.Error(err)
	}
	if n != shouldEqual {
		t.Errorf("Result should have been %d, not %d", shouldEqual, n)
	}
}

func TestGarbage(t *testing.T) {
	n, err := sum(`a, b`)

	if err != nil {
		t.Error(err)
	}

	if n != 0 {
		t.Error("Garbage should return 0")
	}
}
