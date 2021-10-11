package main

import "fmt"

func main() {
	fmt.Println(split(10))
}

func split(sum int) (x, y int) {
	x = sum + 1
	y = sum - x
	return
}
