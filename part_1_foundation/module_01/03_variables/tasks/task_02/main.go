package main

import "fmt"

func main() {
	a := 5
	b := 8
	fmt.Println("Before: a =", a, "b =", b)

	temp := a
	a = b
	b = temp
	fmt.Println("After: a =", a, "b =", b)
}
