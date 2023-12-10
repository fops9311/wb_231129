package main

import (
	"fmt"
	"math"
	"math/big"
)

var a float64 = (math.Pow(2, 20)) + 0.0000001
var b float64 = (math.Pow(2, 20)) + 0.0000001

func main() {
	fmt.Printf("a: %.10f\n", a)
	fmt.Printf("b: %.10f\n", b)
	fmt.Printf("a+b: %.10f\n", a+b)
	fmt.Printf("a*b: %.10f\n", a*b)
	fmt.Printf("a/b: %.10f\n", a/b)

	//пакет биг для больших чисел
	a1 := new(big.Int)
	a1.SetString("24000000000000000000000000000000000000000000000000000000", 10)
	fmt.Println(a1.Text(10))
	b1 := new(big.Int)
	b1.SetString("24000000000000000000000000000000000000000000000000000000", 10)
	fmt.Println(b1.Text(10))
	//результат
	c1 := new(big.Int)
	c1.Add(a1, b1)
	fmt.Printf("a1+b1: %s\n", c1.Text(10))
	c1.Mul(a1, b1)
	fmt.Printf("a1*b1: %s\n", c1.Text(10))
	c1.Div(a1, b1)
	fmt.Printf("a1/b1: %s\n", c1.Text(10))

}
