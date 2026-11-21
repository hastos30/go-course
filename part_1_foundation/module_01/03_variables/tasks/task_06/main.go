package main

import "fmt"

func main() {
	first := "red"
	second := "green"
	third := "blue"
	fmt.Println("=== Before ===")
	fmt.Println("First:", first)
	fmt.Println("Second:", second)
	fmt.Println("Third:", third)

	temp := first
	first = second
	second = third
	third = temp

	fmt.Println("=== After ===")
	fmt.Println("First:", first)
	fmt.Println("Second:", second)
	fmt.Println("Third:", third)
}
