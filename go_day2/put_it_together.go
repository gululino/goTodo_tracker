package main

import "fmt"

func main() {
	number := 7
	fmt.Printf("Multplication table for %d\n", number)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}

}
