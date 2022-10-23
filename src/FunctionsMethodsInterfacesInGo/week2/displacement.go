package main

import (
	"fmt"
	"math"
	"strconv"
)

func genDisplaceFn(a, v, s float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + (v * t) + s
	}

	return fn
}

func main() {
	var acceleration float64
	var initialVelocity float64
	var initialDisplacement float64
	var time float64

	fmt.Println("Please enter acceleration:")
	fmt.Scan(&acceleration)
	fmt.Println("Please enter initial velocity:")
	fmt.Scan(&initialVelocity)
	fmt.Println("Please enter initial displacement:")
	fmt.Scan(&initialDisplacement)
	fmt.Println("Please enter time:")
	fmt.Scan(&time)

	displacementTimeFunc := genDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Println("Displacement: " + strconv.FormatFloat(displacementTimeFunc(time), 'E', -1, 64))
}
