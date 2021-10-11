package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add1(2, 2))
}

func add(x int, y int) int {
	return x + y
}

func add1(x, y int) int {
	return x + y
}
