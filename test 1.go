package main

import (
	"fmt"
	"math"
)

func add(a int, b int) int {
	return a + b
}

func s(r float64) float64 {
	s := math.Pi * r * r
	return s
}

func main() {
	result := add(1, 2)
	fmt.Println(result)
	area := s(5.0)
	fmt.Println(area)
}
