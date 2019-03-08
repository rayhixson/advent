package main

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"runtime"
	"strings"
)

type Device [4]int

type Set [4]int

type Opcoder func(Device, Set) Device

var Opcodes *[]Opcoder

func addr(v Device, in Set) Device {
	v[in[3]] = v[in[1]] + v[in[2]]
	return v
}

func mulr(v Device, in Set) Device {
	v[in[3]] = v[in[1]] * v[in[2]]
	return v
}

func findMatches(before Device, after Device, input Set) {
	for _, f := range *Opcodes {
		final := f(before, input)
		if final == after {
			fmt.Printf("YES (%v could be %v) --> %v\n", input[0], opName(f), final)
		}
	}
}

func parse(reader io.Reader) (before Device, input Set, after Device) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {
			continue
		}

		if strings.Contains(val, "Before") {
			fmt.Sscanf(val, "Before: [%d, %d, %d, %d]", &before[0], &before[1], &before[2], &before[3])
		} else if strings.Contains(val, "After") {
			fmt.Sscanf(val, "After: [%d, %d, %d, %d]", &after[0], &after[1], &after[2], &after[3])
		} else {
			fmt.Sscanf(val, "%d %d %d %d", &input[0], &input[1], &input[2], &input[3])
		}
	}
	return
}

func main() {
	Opcodes = &[]Opcoder{
		addr,
		mulr,
	}

	before, input, after := parse(strings.NewReader(test))

	fmt.Println(before)
	fmt.Println(after)
	findMatches(before, after, input)
}

func opName(i interface{}) string {
	return strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), ".")[1]
}

const test = `
Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]`
