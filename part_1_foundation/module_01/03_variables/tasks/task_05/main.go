package main

import "fmt"

func main() {
	name := "Viktor"
	age := 31
	height := 1.87
	isStudent := false

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Student:", isStudent)
	fmt.Printf("%T %T %T %T\n", name, age, height, isStudent)

	age = age + 1

	fmt.Println("Update age:", age)
}
