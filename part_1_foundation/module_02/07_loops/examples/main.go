package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println("iteration", i)
	}
	fmt.Println("loop finished")

	sum := 0
	for i := 1; i <= 10; i++ {
		sum += 1
	}
	fmt.Println("sum 1..10 =", sum)

	balance := 100
	months := 0

	for balance < 150 {
		balance += 20
		months++
	}

	fmt.Println("months:", months, "balance:", balance)

	n := 1
	for {
		if n*n > 50 {
			break
		}
		n++
	}
	fmt.Println("first n with n*n > 50", n)

	for i := 1; i <= 6; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("odd:", i)
	}
}
