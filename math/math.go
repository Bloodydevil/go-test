package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Printf("Now you have %g problems", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 1337))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
