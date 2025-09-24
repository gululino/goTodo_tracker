package main

import "fmt"

func main() {

	// if else
	age := 18
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println(" You are a minor ")
	}

	// if with short statement
	if x := 10; x > 5 {
		fmt.Println("x is greater than 5")
	}

	// switch
	grade := "A"

	switch grade {
	case "A":
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Good")
	case "C":
		fmt.Println("Needs Improvement")
	}

	var x, y = 100, 2
	if x%y == 0 {
		fmt.Printf("%d is even number", x)

	} else {
		fmt.Printf("%d is odd number ", x)
	}

	// loops
	// basic for loop
	for i := 1; i <= 15; i++ {
		fmt.Println(i)
	}

	// while like Loop
	z := 9
	for z == 1  {
		fmt.Println(z * )
		z--
	}

	// range loop (iterating over a collection)
	nums := []int{1, 4, 3, 5, 6, 7, 8, 90}
	for index, value := range nums {
		fmt.Printf("Index : %d, Value: %d\n", index, value)
	}

	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Println(i)
		}
	}

}
