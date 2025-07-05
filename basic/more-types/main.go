package main

import (
	"fmt"
)

// struct
type Edge struct {
	X, Y int
}

// pointer
func pointer(p *int) {
	x := 2
	p = &x
	fmt.Println("value is: ", *p)
	*p = 3
	fmt.Println("value is changed: ", *p)
}

var (
	e1 = Edge{1, 2}
	p = &Edge{1, 2} // address pointer return value type: *
	e2 = Edge{X: 1} // get X = 1, and Y = 0
	e3 = Edge{} // get X = 0, and Y = 0
)

func edgeProcess() {
	fmt.Println("Edge: ", Edge{1, 2})
	e := Edge{1, 2}
	e.X = 2
	fmt.Println("Change x:", e)
	fmt.Println("List Edges: ", e1, *p, e2, e3)
}

// array
func arrProcess() {
	arr := make([]int, 0)
	for i := 0; i < 5; i++ {
		arr = append(arr, i)
	}
	fmt.Println("Array is: ", arr)
}

// slice
func sliceProcess() {
	names := [4]string {
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("Slice names before: ", names)
	a := names[0: 2]
	b := names[1: 3]
	fmt.Println("List slices a, b: ", a, b)
	b[0] = "XXX"
	fmt.Println("List slices when changed b[0]: ", a, b)
	fmt.Println("Slice names then: ", names)
}

// map
var m2 = map[string]Edge {
	"d1": {
		1, 
		2,
	},
	"d2": {
		3,
		4,
	},
}

func mapProcess() {
	m := make(map[string]Edge)
	m["d1"] = Edge{
		1,
		2,
	}
	fmt.Println("Value in map m of key d1:", m["d1"])
	fmt.Println("List map m2: ", m2)
	_, ok := m["d2"]
	fmt.Println("Check if d2 is in map m: ", ok)
}

// function value
func funcVal(fn func(int, int) int) int {
	return fn(3, 4)
}

func funcValProcess() {
	ch := func(x, y int) int {
		return x + y
	}
	fmt.Println("Value in ch func: ", ch(1, 2))
	fmt.Println("Result funcVal: ", funcVal(ch))
}

// closure
func adder() func(int) int {
	s := 0
	return func(x int) int {
		s += x
		return s
	}
}
func closureProcess() {
	sum, sub := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Print(sum(i), " ", sub(-i), "\n")
	}
}

func main () {
	var p *int
	pointer(p)
	edgeProcess()
	arrProcess()
	sliceProcess()
	mapProcess()
	funcValProcess()
	closureProcess()
}