package main

import "fmt"

func main() {
	bill := student{"Bill", 10, 40, false}
	phoebe := student{"Phoebe", 11, 32, true}

	fmt.Println(bill.isValid())
	fmt.Println(phoebe.isValid())

	phoebe.eligible = false
	fmt.Println(phoebe.isValid())

	fmt.Println(bill)
	bill.changeAge(90)
	fmt.Println(bill)
}

type student struct {
	name     string
	id       int
	age      int
	eligible bool
}

func (s student) isValid() bool {
	return s.age > 30 && s.eligible
}

func (s *student) changeAge(age int) {
	s.age = age
}
