package main

import "fmt"

func main() {
	n := -7

	sign := ""
	parity := ""

	if n > 0 {
		sign = "positive"
	} else if n < 0 {
		sign = "negative"
	} else {
		sign = "zero"
	}

	if n%2 == 0 {
		parity = "even"
	} else {
		parity = "odd"
	}

	fmt.Printf("%d is %s and %s\n", n, sign, parity)
}
