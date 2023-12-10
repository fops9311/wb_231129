package main

import (
	"fmt"
	"math"
)

var v1 float64 = (math.Pow(2, 2)) + 0.0000001
var v2 float64 = (math.Pow(2, 2)) + 0.0000001

func main() {
	fmt.Printf("v1: %.10f\n", v1)
	fmt.Printf("v2: %.10f\n", v2)
	fmt.Printf("v1+v2: %.10f\n", v1+v2)
	fmt.Printf("v1*v2: %.10f\n", v1*v2)
	fmt.Printf("v1/v2: %.10f\n", v1/v2)
}
