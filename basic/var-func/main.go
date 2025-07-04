package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) float64 {
	if (y == 0) {
		return 0.0
	}
	return float64(x) / float64(y)
}

func swap(x, y int) (int, int) {
	return y, x
}

func main () {
	var x, y int
	fmt.Println("Please enter two integers:")
	fmt.Scan(&x, &y)
	fmt.Println("Addition:", add(x, y))
	fmt.Println("Subtraction:", subtract(x, y))
	fmt.Println("Multiplication:", multiply(x, y))
	fmt.Println("Division:", divide(x, y))
	a, b := swap(x, y)
	fmt.Println("Swapped values:", a, b)
}