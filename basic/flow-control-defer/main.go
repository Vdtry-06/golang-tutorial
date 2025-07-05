package main

import (
	"fmt"
	"time"
)

func lop() []int {
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	return arr
}

func condition(x, y int) string {
	if res := x; res < y {
		return fmt.Sprintf("%d is less than %d", x, y)
	}
	return fmt.Sprintf("%d is larger than %d", x, y)
}

func switchCase(x int) {
	switch res := 10 * x; x {
	case 1:
		fmt.Println(res + 1)
	case 2:
		fmt.Println(res + 2)
	default:
		fmt.Println(res + x)
	}
}

func switchWithTime() {
	today := time.Now().Weekday()
	switch time.Saturday {
		case today:
			fmt.Println("Today is Saturday")
		case today + 1:
			fmt.Println("Tomorrow is Sunday")
		case today - 1:
			fmt.Println("Yesterday was Friday")
		default:
			fmt.Println("Today is a Weekday")
	}

	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Good Morning")
		case t.Hour() < 18:
			fmt.Println("Good Afternoon")
		default:
			fmt.Println("Good Evening")
	}
}

func deferFunc() {
	defer fmt.Println("world")
	fmt.Print("Hello ")
}

func deferStack() {
	for i := 0; i < 10; i++ {
		defer fmt.Print(i, " ")
	}
	fmt.Println("Done")
}

func main () {
	arr := lop()
	fmt.Println("Arrays: ", arr)
	fmt.Println("Condition: ", condition(5, 10))
	switchCase(2)
	switchWithTime()
	deferFunc()
	deferStack()
}
