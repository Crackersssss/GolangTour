package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	p := &v.Y
	*p = 1e9
	fmt.Println(v.X, v.Y)
}