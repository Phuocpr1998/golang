package main

import (
"fmt"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func (e Employee) displeyInfo()  {
	fmt.Printf("Name %s %s\nAge %d\nSalary %d",
		e.firstName, e.lastName, e.age, e.salary)
}

func (e Employee) changeName(newName string){
	e.firstName = newName
}

func (e *Employee) changeName2(newName string){
	e.firstName = newName
}

func main() {
	e := Employee{"Sam", "Anderson", 55, 6000}
	e.displeyInfo()

	// khong su dung con tro
	fmt.Printf("Employee name before change: %s", e.firstName)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.firstName)

	// su dung con tro
	fmt.Printf("\n\nEmployee name before change: %s", e.firstName)
	(&e).changeName2("Join")
	fmt.Printf("\nEmployee name after change: %s", e.firstName)
}

