package main

import "fmt"

func main() {
	name := "Viktor"
	age := 30
	height := 1.87
	isStudent := true

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %f\n", height)
	fmt.Printf("Height: %.2f\n", height)
	fmt.Printf("Height: %.1f\n", height)
	fmt.Printf("Student: %t\n", isStudent)

	fmt.Printf("Values: %v %v %v %v\n", name, age, height, isStudent)

	fmt.Printf("%s is %d years old and %.2f m tall\n", name, age, height)

	fmt.Println("Println:", name, age, height)

	fmt.Printf("|%8s|%6d|\n", "apples", 10)
	fmt.Printf("|%-8s|%-6d|\n", "pears", 3)
	fmt.Printf("|%8.2f|\n", 3.14159)

	discount := 15
	fmt.Printf("Discount: %d%%\n", discount)

	label := fmt.Sprintf("%s (%d)", name, age)
	fmt.Println("Lable:", label)
}
