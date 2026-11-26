package main

import "fmt"

func main() {
	age := 30
	isStudent := false
	hour := 10

	price := 200

	if age < 6 {
		price = 0
	} else {
		if age >= 65 || isStudent {
			price = price / 2
		}

		if hour < 12 {
			price = price - 30
		}
	}

	fmt.Printf("Price: %d UAH\n", price)
}
