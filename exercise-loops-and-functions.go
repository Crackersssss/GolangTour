package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("this is math.Sqrt():", math.Sqrt(10000))
	fmt.Println("this is Sqrt():", Sqrt(10000))
}

func Sqrt(x float64) float64 {
	if x < 0 {
		fmt.Println("X has to be greater than zero!")
		return -1
	}

	z := float64(10)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
