package main

import "fmt"

func main() {
	x := 15
	y := 0
	flag := !(x > 10)

	if y != 0 && x/y > 1 {
		fmt.Println("A")
	} else if x%2 == 1 || flag {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// B

	switch {
	case x > 10:
		fmt.Println("D")
	case x > 5:
		fmt.Println("E")
	}
	// D

	switch x % 4 {
	case 0, 1:
		fmt.Println("F")
	case 3:
		fmt.Println("G")
	default:
		fmt.Println("H")
	}

	// G

	fmt.Printf("%t %t\n", flag, x >= 15 && !flag) // false, true
}
