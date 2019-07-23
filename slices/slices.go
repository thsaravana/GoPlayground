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


	// Using slices
	fmt.Println("Using slices")
	slice5 := []int {1,2,3,4,5}
	slice5[0] = -1
	fmt.Println(slice5)

	slice6 := slice5[1:3]
	fmt.Println(slice6)

	// Appending slices
	fmt.Println("Appending slices")
	slice7 := []int {1,2,3,4,5}
	slice8 := append(slice7, 6)
	fmt.Println(slice8)

	slice9 := []int {1,2}
	slice10 := append(slice9, 3, 4, 5)
	fmt.Println(slice10)

	slice11 := []int {6,7,8}
	slice12 := append(slice10, slice11...)
	fmt.Println(slice12)


	// Iterating over slices
	slice13 := []int {1,2,3,4,5}
	
	for index, value := range slice13 {
		fmt.Printf("Index: %d  Value: %d\n", index, value)
	}

	for index := 0; index < len(slice13); index++ {
		fmt.Println(slice13[index])
	}
}