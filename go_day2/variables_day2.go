package main

import "fmt"

func main() {
	// Explicit declaration
	var name string
	name = "John Doe"
	fmt.Println(name)

	// initialization
	var age int = 25
	fmt.Println(age)

	// type inferred and short declaration
	fullname := "John Joseph"
	fmt.Println(fullname)

	// multiple declaration
	var x, y int = 10, 20
	z := x + y
	fmt.Println(z)

	// Constants
	const Pi = 3.14
	fmt.Println(Pi)

	var rectangle = 4 * 6
	fmt.Println(rectangle)
}
