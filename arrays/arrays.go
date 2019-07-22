package main

import "fmt"

func main() {

	// Different array initializations
	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]int {1,2,3,4,5}
	fmt.Println(array2)

	array3 := [...]int {1,2,3}
	fmt.Println(array3)

	array4 := [5]int {0:10, 4:10}
	fmt.Println(array4)

	// Using Arrays
	array4[3] = 80
	fmt.Println(array4[3])

	// Array Copy via Assigning
	var array5 = array4
	array5[4] = 3000
	fmt.Println(array5)
	fmt.Println(array4)

	// Multidimensional arrays
	var multi1 [4][2]int
	fmt.Println(multi1)

	multi2 := [4][2]int {{1,1},{2,2}, {3,3}, {4,4} }
	fmt.Println(multi2)

	multi3 := [...][2]int {{1,1},{2,2}, {3,3}, {4,4} }
	fmt.Println(multi3)

	multi3[0][1] = 45
	fmt.Println(multi3)


	arrayValue := [4]int {1,2,3,4}
	// Array pass by value
	fmt.Println("Array pass by value:")
	passByValue(arrayValue)
	fmt.Println(arrayValue)

	// Array pass by reference
	fmt.Println("Array pass by reference:")
	passByReference(&arrayValue)
	fmt.Println(arrayValue)
}

func passByValue(array [4]int) {
	array[0] = 98
	fmt.Println(array)
}

func passByReference(array *[4]int) {
	array[0] = 90
	fmt.Println(*array)
}