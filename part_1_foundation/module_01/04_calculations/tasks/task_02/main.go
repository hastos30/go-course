package main

import "fmt"

func main() {
	price := 249.99
	quantity := 3
	discountPercent := 15

	subTotal := price * float64(quantity)
	discount := subTotal * (float64(discountPercent) / 100)
	total := subTotal - discount

	fmt.Printf("Subtotal:%8.2f\n", subTotal)
	fmt.Printf("Discount:%8.2f (%d%%)\n", discount, discountPercent)
	fmt.Printf("Total:   %8.2f\n", total)
}
