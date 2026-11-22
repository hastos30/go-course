package main

import "fmt"

func main() {
	item := "Pen"
	price := 45
	quantity := 11
	const discountPercent = 12
	const boxSize = 4
	delivery := 180
	people := 5

	subTotal := price * quantity
	discountAmount := float64(subTotal) * (float64(discountPercent) / 100)
	total := float64(subTotal) - discountAmount + float64(delivery)
	pricePerPerson := total / float64(people)
	fullBoxes := quantity / boxSize
	leftOver := quantity % boxSize

	fmt.Println("Receipt")
	fmt.Printf("%-12s %s\n", "Item:", item)
	fmt.Printf("%-12s %d RUB\n", "Price:", price)
	fmt.Printf("%-12s %d\n", "Quantity:", quantity)
	fmt.Printf("%-12s %d RUB\n", "Subtotal:", subTotal)
	fmt.Printf("%-12s %d%% (-%.2f RUB)\n", "Discount:", discountPercent, discountAmount)
	fmt.Printf("%-12s %d RUB\n", "Delivery:", delivery)
	fmt.Printf("%-12s %.2f RUB\n", "Total:", total)
	fmt.Printf("%-12s %.2f RUB (%d people)\n", "Per person:", pricePerPerson, people)
	fmt.Printf("%-12s %d full, %d left over\n", "Boxes:", fullBoxes, leftOver)
}
