package main

import (
	"fmt"
)

// struct
type Edge struct {
	X, Y int
}

// Method
func (e Edge) Add() int {
	return e.X + e.Y
}

// C1
func (e *Edge) ScaleC1(multi int) {
	e.X *= multi
	e.Y *= multi
}
func process1() {
	e := Edge{3, 4}
	e.ScaleC1(10)
	fmt.Println("Result by method: ", e.Add())
}

// C2
func ScaleC2(e *Edge, multi int) {
	e.X *= multi
	e.Y *= multi
}
func process2() {
	e := Edge{3, 4}
	ScaleC2(&e, 10)
	fmt.Println("Result by func: ", e.Add())
}

// no struct
type Value int

func (v Value) Add() int {
	if v < 0 {
		return -1
	}
	return int(v)
}

func addProcess() {
	v := Value(-2)
	fmt.Println("Result no struct: ", v.Add())
}

// interface
type I interface {
	UserName()
}

type Name struct {
	FirstName string
	LastName  string
}

func (name Name) UserName() {
	fmt.Println("UserName by interface: ", name.FirstName+" "+name.LastName)
}

func nameProcess() {
	var i I = Name{"Trong", "Vuong"}
	i.UserName()
}

// stringer
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v year)", p.Name, p.Age)
}

func personProcess() {
	p1 := Person{"Trong", 20}
	p2 := Person{"Le", 25}
	fmt.Println(p1, ",", p2)
}

// custom error
type MyError struct {
	Study string
	Work  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("I am still learning %s, having an error %s", e.Study, e.Work)
}

func Run() error {
	return &MyError{
		"Golang",
		"Syntax",
	}
}

func runProcess() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	process1()
	process2()
	addProcess()
	nameProcess()
	personProcess()
	runProcess()
}
