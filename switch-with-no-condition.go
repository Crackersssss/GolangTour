package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now()
	switch {
	case hour.Hour() < 12:
		fmt.Println("Good morning")
	case hour.Hour() < 19:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
