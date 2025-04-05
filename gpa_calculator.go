/*

Fundamentals of Go Tasks
Task: Student Grade Calculator
Create a Go console application that allows students to calculate their average grade based on different subjects. The application should prompt the student to enter their name and the number of subjects they have taken. For each subject, the student should enter the subject name and the grade obtained (numeric value). After entering all subjects and grades, the application should display the student's name, individual subject grades, and the calculated average grade.
Requirements:
* Use variables and data types to store student data.
* Use conditional statements to validate input (e.g., ensure grade values are within a valid range).
* Implement loops to handle multiple subjects and grades.
* Utilize collections (e.g., List, Dictionary) to store subject names and corresponding grades.
* Define a method to calculate the average grade based on the entered grades.
* Use string interpolation to display the results in a user-friendly format.
* Write test for your code [Optional]

*/

package main

import (
	"fmt"
	"strings" // To convert subject name and grades to uppercase
)

func main() {
	// Possible Grades
	grades := map[string]int{
		"A": 5,
		"B": 4,
		"C": 3,
		"D": 2,
		"E": 1,
		"F": 0,
	}

	var name string
	fmt.Print("What is your name: ")
	fmt.Scan(&name)

	scores := make(map[string]string)
	var noCourses int
	var total_unit float64 = 0
	var scored float64 = 0
	var subject string
	var unit float64
	var grade string

	fmt.Print("How many courses did you offer: ")
	fmt.Scan(&noCourses)

	for i := 0; i < noCourses; i++ {
		fmt.Print(i+1, ". Subject: ")
		fmt.Scan(&subject)
		fmt.Print(i+1, ". Units: ")
		fmt.Scan(&unit)
		fmt.Print(i+1, ". Grade (e.g. A, B, C, etc.): ")
		fmt.Scan(&grade)

		grade = strings.ToUpper(grade)
		subject = strings.ToUpper(subject)

		scores[subject] = grade
		total_unit += unit
		scored += float64(grades[grade]) * unit
	}
	fmt.Println(total_unit, scored/total_unit)

	fmt.Printf("***%s RESULT***\n", strings.ToUpper(name))
	for key, val := range scores {
		fmt.Printf("%s - %s\n", key, val)
	}
	fmt.Printf("GPA: %.2f", scored/total_unit) // Precision to 2 decimal places.
}
