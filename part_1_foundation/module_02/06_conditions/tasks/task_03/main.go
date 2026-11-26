package main

import "fmt"

func main() {
	a := 10.0
	b := 4.0
	op := "/"

	var result float64

	switch op {
	case "+":
		result = a + b
		fmt.Printf("%.2f + %.2f = %.2f\n", a, b, result)

	case "-":
		result = a - b
		fmt.Printf("%.2f - %.2f = %.2f\n", a, b, result)

	case "*":
		result = a * b
		fmt.Printf("%.2f * %.2f = %.2f\n", a, b, result)

	case "/":
		if b == 0 {
			fmt.Println("Error: division by zero")
		} else {
			result = a / b
			fmt.Printf("%.2f / %.2f = %.2f\n", a, b, result)
		}

	default:
		fmt.Printf("Error: unknown operator %s\n", op)
	}
}
