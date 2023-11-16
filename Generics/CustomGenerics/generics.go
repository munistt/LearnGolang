package main

import (
	"fmt"
)

func main() {

	a1 := []int{1, 2, 3}
	a2 := []float64{1.2, 3.4, 5.6}
	s1 := []string{"foo", "bar", "baz"}

	a11 := add(a1)
	a12 := add(a2)
	a13 := add(s1)

	fmt.Printf("Sum of %v : %v\n", a1, a11)
	fmt.Printf("Sum of %v : %v\n", a2, a12)
	fmt.Printf("Sum of %v : %v\n", s1, a13)
}

type addable interface {
	int | float64 | string
}

func add[V addable](s []V) V {

	var result V
	for _, v := range s {
		result += v
	}
	return result
}
