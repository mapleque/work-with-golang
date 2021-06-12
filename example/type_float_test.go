package example

import (
	"fmt"
	"math"
)

func Example_float() {
	var maxFloat32 float32 = math.MaxFloat32
	var minFloat32 float32 = math.SmallestNonzeroFloat32
	var maxFloat64 float64 = math.MaxFloat64
	var minFloat64 float64 = math.SmallestNonzeroFloat64
	fmt.Printf("max float32: %v\n", maxFloat32)
	fmt.Printf("min float32: %v\n", minFloat32)
	fmt.Printf("max float64: %v\n", maxFloat64)
	fmt.Printf("min float64: %v\n", minFloat64)
	pi := math.Pi
	fmt.Printf("pi: %v\n", pi)
	var one = 1.0
	fmt.Printf("one: %v\n", one)
	fmt.Printf("one+one: %v\n", one+one)
	fmt.Printf("one*one: %v\n", one*one)
	fmt.Printf("one==one: %v\n", one == one)

	// Output:
	// max float32: 3.4028235e+38
	// min float32: 1e-45
	// max float64: 1.7976931348623157e+308
	// min float64: 5e-324
	// pi: 3.141592653589793
	// one: 1
	// one+one: 2
	// one*one: 1
	// one==one: true
}
