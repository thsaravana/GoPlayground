package main

import "fmt"

func main() {
	// Slice initialization
	fmt.Println("Slice initialization")

	slice1 := make([]int, 5)
	fmt.Println(slice1)

	slice2 := make([]int, 3, 5)
	fmt.Println(slice2)

	slice3 := []int {1,2,3,4,5}
	fmt.Println(slice3)

	slice4 := []int {10: 100}
	fmt.Println(slice4)

	// nil slice
	fmt.Println("nil slice")
	var nilSlice []int
	fmt.Println(nilSlice)

	// empty slice
	fmt.Println("empty slice")
	emptySlice1 := make([]int, 0)
	fmt.Println(emptySlice1)

	emptySlice2 := []int {}
	fmt.Println(emptySlice2)
}