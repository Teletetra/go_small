package main

import (
	"fmt"
)

func PrintSlice[T any](items []T) {
	for _, v := range items {
		fmt.Println(v)
	}
}

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// Sum works only for numeric types
func Sum[T Number](a, b T) T {
	return a + b
}


type Box[T any] struct {
	value T
}

func (b Box[T]) Get() T {
	return b.value
}

func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}


func main() {

	// Generic function example
	fmt.Println("=== PrintSlice Example ===")
	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"Go", "Generics", "Awesome"})

	// Generic sum
	fmt.Println("\n=== Sum Example ===")
	fmt.Println(Sum(10, 20))
	fmt.Println(Sum(3.5, 2.5))

	// Generic struct
	fmt.Println("\n=== Generic Struct Example ===")
	intBox := Box[int]{value: 100}
	strBox := Box[string]{value: "Hello"}

	fmt.Println(intBox.Get())
	fmt.Println(strBox.Get())

	// Comparable constraint
	fmt.Println("\n=== Contains Example ===")
	fmt.Println(Contains([]int{1, 2, 3}, 2))
	fmt.Println(Contains([]string{"a", "b", "c"}, "z"))
}