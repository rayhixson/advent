package main

import (
	"fmt"
	"strconv"
)

func find10After(n int) ([]int, string) {
	r := doit(n + 10)
	ten := (*r)[n : n+10]

	s := ""
	for _, n := range ten {
		s += strconv.Itoa(n)
	}

	return ten, s
}

func findSequence(seq ...int) int {
	r := doit(100000000)

	len := len(seq)
	for i, _ := range *r {
		section := (*r)[i : i+len]
		eq := true
		for j, v := range section {
			if v != seq[j] {
				eq = false
			}
		}
		if eq {
			fmt.Printf("section [%v] == seq [%v]\n", section, seq)
			return i
		}
	}
	return -1
}

func doit(n int) *[]int {
	recipes := make([]int, 0, 500000)
	elf1 := 0
	elf2 := 1
	recipes = append(recipes, 3, 7, 1, 0)
	//dump(&recipes, elf1, elf2)

	for i := 0; i < n; i++ {
		sum := recipes[elf1] + recipes[elf2]
		if sum < 10 {
			recipes = append(recipes, sum)
		} else {
			ones := sum % 10
			//fmt.Println("--ones>", ones)
			tens := (sum - ones) / 10 % 10
			//fmt.Println("--tens>", tens)
			recipes = append(recipes, tens)
			recipes = append(recipes, ones)
		}
		elf1 = rotate(&recipes, elf1)
		//fmt.Println("e1>", elf1)
		elf2 = rotate(&recipes, elf2)
		//fmt.Println("e2>", elf2)

		//dump(&recipes, elf1, elf2)
	}

	return &recipes
}

func rotate(recipes *[]int, elf int) int {
	inc := (*recipes)[elf] + 1
	return (elf + inc) % len(*recipes)
}

func dump(recipes *[]int, e1, e2 int) {
	s := ""
	for i := 0; i < len(*recipes); i++ {
		format := " %d "
		if i == e1 {
			format = "(%d)"
		} else if i == e2 {
			format = "[%d]"
		}
		s += fmt.Sprintf(format, (*recipes)[i])
	}

	fmt.Println(s)
}

func main() {
	_, s := find10After(509671)
	fmt.Println("Ten: ", s)

	c := findSequence(5, 0, 9, 6, 7, 1)
	fmt.Println("Found after:", c)
}
