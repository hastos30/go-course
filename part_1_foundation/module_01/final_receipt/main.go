package main

import "fmt"

func main() {
	item := "Notebook"
	price := 349
	quantity := 7
	const discount = 15
	const boxSize = 3
	delivery := 250
	perPerson := 3

	subTotal := price * quantity
	discountAmount := float64(subTotal) * (float64(discount) / 100)
	total := float64(subTotal) - discountAmount + float64(delivery)
	pricePerPerson := total / float64(perPerson)
	fullBoxes := quantity / boxSize
	leftOver := quantity % boxSize

	fmt.Println("Receipt")
	fmt.Printf("Item:		%s\n", item)
	fmt.Printf("Price:		%d RUB\n", price)
	fmt.Printf("Quantity:	%d\n", quantity)
	fmt.Printf("Subtotal:	%d RUB\n", subTotal)
	fmt.Printf("Discount:	%d%% (-%.2f RUB)\n", discount, discountAmount)
	fmt.Printf("Delivery:	%d RUB\n", delivery)
	fmt.Printf("Total: 		%.2f RUB\n", total)
	fmt.Printf("Per person:	%.2f RUB (%d people)\n", pricePerPerson, perPerson)
	fmt.Printf("Boxes:		%d full, %d left over\n", fullBoxes, leftOver)
}
