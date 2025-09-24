package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// factorial function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// prime numbers checker
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true

}

// mini project fibonacci sequence
func fibonacci(n int) []int {
	if n < 0 {
		return []int{}
	}
	fibs := []int{0, 1}
	for i := 2; i < n; i++ {
		fibs = append(fibs, fibs[i-1]+fibs[i-2])
	}
	return fibs
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("Result: ", result)
	}

	fmt.Println(factorial(5))
	fmt.Println(isPrime(31))
	fmt.Println(fibonacci(10))

}
