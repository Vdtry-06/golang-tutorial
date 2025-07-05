package main

import (
	"fmt"
)

// type parameters
func Index[T comparable](arr []T, x T) int {
	for i, v := range arr {
		if v == x {
			return i
		}
	}
	return -1
}

func IndexProcess() {
	ai := []int{1, 2, 3, 4, 5}
	fmt.Println("Check if 3 is in arr at position: ", Index(ai, 3))
	as := []string{"Hi", "Hello", "Have a good Day"}
	fmt.Println("Check if 'too' is in arr at position: ", Index(as, "too"))
}

// Generic Type
var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
}
var floats = map[string]float64{
	"one":   1.2,
	"two":   2.2,
	"three": 3.2,
}

// Case 1
func SumInts(m map[string]int) int {
	var s int
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsProcess() {
	fmt.Println("Result by sum int: ", SumInts(numbers))
}

func SumFloatsProcess() {
	fmt.Println("Result by sum float: ", SumFloats(floats))
}

// Case 2
type Numbers interface {
	int | float64
}

func SumIntsOrFloats[K comparable, V Numbers](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloatsProcess() {
	fmt.Println("Result by sum IntOrFloat: ", SumIntsOrFloats(numbers))
	fmt.Println("Result by sum IntOrFloat: ", SumIntsOrFloats(floats))
}

func main() {
	IndexProcess()
	SumIntsProcess()
	SumFloatsProcess()
	SumIntsOrFloatsProcess()
}
