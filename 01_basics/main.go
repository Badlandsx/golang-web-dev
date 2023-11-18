package main

import (
	"fmt"
)

// This variable scope is package level
var testScopePackage int

func main() {
	helloWorld()
	polymorphism()
	slice()
	maps()
}

// Hello world
func helloWorld() {
	x := "Hello, class!"
	fmt.Println(x)
}

// Polymorphism
type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func polymorphism() {
	p1 := person{
		"Miss",
		"Moneypenny",
	}

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	saySomething(p1)
	saySomething(sa1)
}

// slice
func slice() {
	x := []int{3, 4, 6, 8, 10, 100}
	fmt.Println(x)
}

// maps
func maps() {
	// y scope is in this function block
	y := map[string]int{
		"test": 10,
		"JP":   30,
	}
	fmt.Println(y)
}
