package main

import "fmt"

func main() {
	var name string = "Viktor"
	var age int = 31
	var height float64 = 1.87
	var isStudent bool = false

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Student:", isStudent)

	age = age + 1

	fmt.Println("Update age:", age)
}
