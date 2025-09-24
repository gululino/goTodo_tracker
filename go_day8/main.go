package main

import "fmt"

// define a struct
type Person struct {
	Name string
	Age  int
}

// methods for Person struct
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

// method with a pointer reciever (modieis the struct)
func (p *Person) HaveBirthday() {
	p.Age++ // modifies the original struct
	fmt.Println(p.Age)
}

func main() {
	// create an instance of the struct
	p := Person{Name: "Alice", Age: 89}

	// access struct fields
	fmt.Println("Name:", p.Name)
	fmt.Println("Age", p.Age)

	p.Greet()
	p.HaveBirthday()

}
