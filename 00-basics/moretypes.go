package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X float64
	Y float64
}

var m map[string]vertex

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func moretypes() {
	// Pointers
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// Struct
	v := vertex{1, 2}
	g := &v
	g.X = 15
	fmt.Println(v)

	// Arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices
	var s []int = primes[1:4]
	// s := primes[1:4] // literals
	// s := make([]int, 3, 4) // make
	fmt.Println(s)

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	// we can add more than one element at a time.
	s = append(s, 2, 3, 4)

	// print with loop
	for i, v := range s {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Map
	m = make(map[string]vertex)
	m["Bell Labs"] = vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// mutating map
	ma := make(map[string]int)

	ma["Answer"] = 42
	fmt.Println("The value:", ma["Answer"])

	ma["Answer"] = 48
	fmt.Println("The value:", ma["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", ma["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// Function values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
