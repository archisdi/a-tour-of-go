package main

import (
	"fmt"
	"math"
)

// function
func add(x int, y int) int {
	return x + y
}

// multiple result
func swap(x, y string) (string, string) {
	return y, x
}

// named return function
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func funcAndVar() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(add(42, 13))

	// multiple returns function
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// variables
	var c, python, java bool
	var i int
	fmt.Println(i, c, python, java)

	// variable initializer
	var x, y int = 1, 2
	var l, golang, js = true, false, "no!"
	fmt.Println(l, x, y, golang, js)

	// short var declaration
	name := "archie"
	fmt.Println(name)

	// constant
	const Pi = 3.14
	fmt.Println(Pi)
}
