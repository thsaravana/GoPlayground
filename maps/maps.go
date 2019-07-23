package main

import "fmt"

func main() {
	// Initializing
	fmt.Println("Initializing")

	map1 := make(map[int] string)
	fmt.Println(map1)

	map2 := map[int] string { 1 : "One", 2 : "Two" }
	fmt.Println(map2)
}