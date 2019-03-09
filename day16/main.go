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

var KnownOpcodes [16]Opcoder

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
		// also if it's a known opcode don't compare
		known := false
		for _, k := range KnownOpcodes {
			if k != nil && opName(f) == opName(k) {
				known = true
				break
			}
		}
		if known {
			continue
		}

		orig := s.Before
		f(&orig, s.Input)
		if orig == s.After {
			//fmt.Printf("YES (%v could be %v)\n", s.Input[0], opName(f))
			matches = append(matches, f)
		}
	}
	return
}

func findThree(samples []*Sample) {
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

func testEachOpCodeAgainstSamples(samples []*Sample) {

	for i := 0; i < 16; i++ {
		opCounts := make(map[string]int)
		sampleCount := 0
		for _, s := range samples {
			if s.Input[0] == i {
				sampleCount++
				//fmt.Println("Sample: ", s)
				matches := findMatches(s)
				//fmt.Printf("%v: %v\n", i, names(matches))
				for _, m := range matches {
					n := opName(m)
					opCounts[n] = opCounts[n] + 1
				}
			}
		}
		if len(opCounts) > 0 {
			fmt.Printf("%v (%v) ==> %v\n", i, sampleCount, opCounts)
		}
		//break
	}

}

func runTestProgram(reader io.Reader) Device {
	scanner := bufio.NewScanner(reader)
	d := Device{}
	for scanner.Scan() {
		s := Set{}
		fmt.Sscanf(scanner.Text(), "%d %d %d %d", &s[0], &s[1], &s[2], &s[3])
		KnownOpcodes[s[0]](&d, s)
	}

	return d
}

func names(oa []Opcoder) (s string) {
	for _, o := range oa {
		s += opName(o) + " "
	}
	return s
}

func main() {
	/*
		samples := parse(strings.NewReader(test))
	*/
	/*
		r, err := os.Open("input")
		if err != nil {
			panic(err)
		}

		samples := parse(r)

		//findThree(samples)
		testEachOpCodeAgainstSamples(samples)
	*/

	r, err := os.Open("test_program")
	if err != nil {
		panic(err)
	}
	s := runTestProgram(r)
	fmt.Println("Final: ", s)
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
	KnownOpcodes[11] = eqri
	KnownOpcodes[10] = eqrr
	KnownOpcodes[7] = eqir
	KnownOpcodes[15] = gtri
	KnownOpcodes[13] = gtrr
	KnownOpcodes[4] = gtir
	KnownOpcodes[2] = banr
	KnownOpcodes[3] = bani
	KnownOpcodes[5] = setr
	KnownOpcodes[8] = seti
	KnownOpcodes[14] = mulr
	KnownOpcodes[1] = muli
	KnownOpcodes[0] = bori
	KnownOpcodes[12] = borr
	KnownOpcodes[6] = addr
	KnownOpcodes[9] = addi

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
