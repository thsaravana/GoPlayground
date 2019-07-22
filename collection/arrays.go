package main

import "fmt"

func main() {

	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]int {1,2,3,4,5}
	fmt.Println(array2)

	array3 := [...]int {1,2,3}
	fmt.Println(array3)

	array4 := [5]int {0:10, 4:10}
	fmt.Println(array4)

	array4[3] = 80
	fmt.Println(array4[3])

	var array5 = array4
	array5[4] = 3000
	fmt.Println(array5)
	fmt.Println(array4)

}