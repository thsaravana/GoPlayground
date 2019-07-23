package main

import "fmt"

func main() {
	// Initializing
	fmt.Println("Initializing")

	map1 := make(map[int] string)
	fmt.Println(map1)

	map2 := map[int] string { 1 : "One", 2 : "Two" }
	fmt.Println(map2)

	// Working with Maps
	fmt.Println("Working with maps")
	map3 := map[int] string {}
	map3[1] = "One"
	map3[2] = "Two"
	map3[3] = "Three"
	map3[1] = "Uno"
	fmt.Println(map3)

	value, exists := map3[1]
	if exists {
		fmt.Printf("1 is %s\n", value)
	}

	value1, exists1 := map3[11]
	if exists1 {
		fmt.Printf("1 is %s\n", value1)
	} else {
		fmt.Println("Nothing to see here")
	}
}