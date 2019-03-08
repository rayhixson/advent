package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
)

type Device [4]int

type Set [4]int

type Opcoder func(*Device, Set)

var Opcodes []Opcoder

type Sample struct {
	Before Device
	After  Device
	Input  Set
}

func (s Sample) String() string {
	return fmt.Sprintf("%v -> %v = %v", s.Before, s.Input, s.After)
}

func parse(reader io.Reader) (ss []*Sample) {
	scanner := bufio.NewScanner(reader)
	s := &Sample{}
	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {
			continue
		}

		if strings.Contains(val, "Before") {
			fmt.Sscanf(val, "Before: [%d, %d, %d, %d]", &s.Before[0], &s.Before[1], &s.Before[2], &s.Before[3])
		} else if strings.Contains(val, "After") {
			fmt.Sscanf(val, "After: [%d, %d, %d, %d]", &s.After[0], &s.After[1], &s.After[2], &s.After[3])
			ss = append(ss, s)
			s = &Sample{}
		} else {
			fmt.Sscanf(val, "%d %d %d %d", &s.Input[0], &s.Input[1], &s.Input[2], &s.Input[3])
		}
	}
	return ss
}

func findMatches(s *Sample) (matches []Opcoder) {
	for _, f := range Opcodes {
		orig := s.Before
		f(&orig, s.Input)
		if orig == s.After {
			//fmt.Printf("YES (%v could be %v)\n", s.Input[0], opName(f))
			matches = append(matches, f)
		}
	}
	return
}

func main() {
	/*
		samples := parse(strings.NewReader(test))
	*/
	r, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	samples := parse(r)

	total := 0
	for _, s := range samples {
		matches := findMatches(s)
		if len(matches) >= 3 {
			//fmt.Println(s)
			//fmt.Println("Matches: ", len(matches))
			total++
		}
	}

	fmt.Printf("Total samples [%v] matching 3 or more: %v\n", len(samples), total)
}

func opName(i interface{}) string {
	return strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), ".")[1]
}

const test = `
Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]`

func addr(v *Device, in Set) { (*v)[in[3]] = (*v)[in[1]] + (*v)[in[2]] }
func addi(v *Device, in Set) { (*v)[in[3]] = (*v)[in[1]] + in[2] }
func mulr(v *Device, in Set) { (*v)[in[3]] = (*v)[in[1]] * (*v)[in[2]] }
func muli(v *Device, in Set) { (*v)[in[3]] = (*v)[in[1]] * in[2] }
func banr(v *Device, in Set) { (*v)[in[3]] = (*v)[in[1]] & (*v)[in[2]] }
func bani(v *Device, in Set) { (*v)[in[3]] = (*v)[in[1]] & in[2] }
func borr(v *Device, in Set) { (*v)[in[3]] = (*v)[in[1]] | (*v)[in[2]] }
func bori(v *Device, in Set) { (*v)[in[3]] = (*v)[in[1]] | in[2] }
func setr(v *Device, in Set) { (*v)[in[3]] = (*v)[in[1]] }
func seti(v *Device, in Set) { (*v)[in[3]] = in[1] }
func gtir(v *Device, in Set) { (*v)[in[3]] = gt(in[1], (*v)[in[2]]) }
func gtri(v *Device, in Set) { (*v)[in[3]] = gt((*v)[in[1]], in[2]) }
func gtrr(v *Device, in Set) { (*v)[in[3]] = gt((*v)[in[1]], (*v)[in[2]]) }
func eqir(v *Device, in Set) { (*v)[in[3]] = eq(in[1], (*v)[in[2]]) }
func eqri(v *Device, in Set) { (*v)[in[3]] = eq((*v)[in[1]], in[2]) }
func eqrr(v *Device, in Set) { (*v)[in[3]] = eq((*v)[in[1]], (*v)[in[2]]) }

func gt(x, y int) int {
	if x > y {
		return 1
	}
	return 0
}

func eq(x, y int) int {
	if x == y {
		return 1
	}
	return 0
}

func init() {
	Opcodes = []Opcoder{
		addr,
		addi,
		mulr,
		muli,
		banr,
		bani,
		borr,
		bori,
		setr,
		seti,
		gtir,
		gtri,
		gtrr,
		eqir,
		eqri,
		eqrr,
	}
}
