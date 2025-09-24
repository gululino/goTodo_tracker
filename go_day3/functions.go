package main

import "fmt"

// defining functions
func greet(name string) string {
	return "Hello, " + name
}

//functions with multiple retrns
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// named return values
func addSubtract(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

// Variadic functions
func sumu(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println(greet("Zeinab"))
	quotient, remainder := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
	sum, diff := addSubtract(10, 8)
	fmt.Printf("Sum: %d, Difference: %d\n", sum, diff)
	fmt.Println(sumu(1, 4, 6, 7, 8, 9))
}
