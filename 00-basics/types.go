package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func types() {

	// basic types
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// zero values
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// type conversion
	var x, y int = 3, 4
	var n float64 = math.Sqrt(float64(x*x + y*y))
	var m uint = uint(n)
	fmt.Println(x, y, m)
}
