package main

import "fmt"

// Define a student struct

type Student struct {
	ID     int
	Name   string
	Grades []int
}

// add a method to calculate the average grades for a student
func (s Student) AverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	total := 0
	for _, grade := range s.Grades {
		total += grade
	}

	return float64(total) / float64(len(s.Grades))
}

// Main program
func main() {
	students := make(map[int]Student)

	// add students
	students[1] = Student{
		ID:     1,
		Name:   "Alice",
		Grades: []int{90, 88, 88},
	}

	students[2] = Student{
		ID:     1,
		Name:   "Manzi",
		Grades: []int{90, 93, 88},
	}

	// add a new student dyanmically

	newStudent := Student{
		ID:     3,
		Name:   "Zeinab",
		Grades: []int{98, 82, 100},
	}

	students[newStudent.ID] = newStudent

	// Display all students
	fmt.Println("All students: ")
	for _, student := range students {
		fmt.Printf("ID: %d, Name: %s, Grades: :%v, Average: %.2f\n",
			student.ID, student.Name, student.Grades, student.AverageGrade())
	}

	// calculate and display the average grade for the class
	classAverage := calculateClassAverage(students)
	fmt.Printf("\nClass Avearge Grade: %.2f\n", classAverage)

	// search for a student by ID
	StudentId := 2
	if student, found := students[StudentId]; found {
		fmt.Printf("\nDetails for Student Id %d:\n", StudentId)
		fmt.Printf("Name: %s, Grades: %v, Average: %.2f\n",
			student.Name, student.Grades, student.AverageGrade())
	} else {
		fmt.Printf("\nStudent ID %d not found. \n", StudentId)
	}
}

// function to calculate the class average

func calculateClassAverage(students map[int]Student) float64 {
	total := 0
	count := 0
	for _, student := range students {
		for _, grade := range student.Grades {
			total += grade
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return float64(total) / float64(count)
}
