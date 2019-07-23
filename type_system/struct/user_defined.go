package main

import "fmt"

func main() {
	fmt.Println("User Defined types")

	// Initializing
	var bill student
	fmt.Println(bill)

	murray := student{
		name:     "Murray",
		id:       12,
		eligible: false,
	}
	fmt.Println(murray)

	joey := student{"Joey", 31, true}
	fmt.Println(joey)

	grade1 := class{
		name:     "Grade 1",
		students: []student{murray, joey},
	}
	fmt.Println(grade1)
}

type student struct {
	name     string
	id       int
	eligible bool
}

type class struct {
	students []student
	name     string
}
