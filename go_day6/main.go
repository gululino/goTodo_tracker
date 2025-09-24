package main

import "fmt"

func main() {
	// declare an array of interger

	var arr [5]int

	// initialize

	arr = [5]int{10, 20, 30, 40, 50}

	// accessing elements
	fmt.Println("Elements at index 0:", arr[0])
	fmt.Println("Full array: ", arr)

	slices()
	maps()
	structss()

}

func slices() {
	// slices are dynamically sized and backed by an array
	// append to add elements to a slice
	// create a slice using shorthand
	slice := []int{10, 20, 30, 40, 50}

	// append elements
	slice = append(slice, 60, 70)
	fmt.Println("After appending:", slice)

	//slice a slice
	subSlice := slice[1:4] // includes index 1 to 3 excludes 4
	fmt.Println("Sliced portion:", subSlice)

	// modify the original slice
	slice[0] = 100
	fmt.Println("Modified original slice:", slice)
}

func maps() {
	myMap := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	// access vlaues by key
	fmt.Println("Alice's age:", myMap["Alice"])

	// add a new key-value pair
	myMap["Manzi"] = 288
	fmt.Println("updated map: ", myMap)

	// check if a key exists
	value, exists := myMap["Diana"]
	if exists {
		fmt.Println("Diana's age is ", value)
	} else {
		fmt.Println("Diana not found")
	}

	// iterate throug maps
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

// alright lets look at struct
func structss() {
	type Person struct {
		Name string
		Age  int
	}
	// create a struct instance
	P := Person{Name: "Alice", Age: 35}

	// accessing fields
	fmt.Println("Name: ", P.Name)
	fmt.Println("Age:", P.Age)

	// Modify Fields
	P.Age = 26
	fmt.Println("Updated age: ", P.Age)

	// anonmous struct

	// anonymous := stuct {
	// 	Brand string,
	// 	Price int
	// }{
	// 	Brand: "Tesla"
	// 	Price: 500000
	// }

	// fmt.println("Anonymous Struct:", anonymous)

}
