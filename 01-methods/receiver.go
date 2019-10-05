package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

type myFloat float64

// basic receiver function
func (v vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// pointer receiver
func (v *vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func receiver() {
	v := vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	f := myFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
