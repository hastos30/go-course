package main

import "fmt"

func main() {
	contribution := 1000.0
	const rate = 0.10
	years := 0

	for contribution < 2000.0 {
		contribution += contribution * rate
		years++

		fmt.Printf("year %d: %.2f\n", years, contribution)
	}

	fmt.Printf("years needed: %d\n", years)
}
