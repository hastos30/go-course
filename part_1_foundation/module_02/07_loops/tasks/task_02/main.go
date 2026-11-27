package main

import "fmt"

func main() {
	sum := 0
	count := 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
		count++
	}
	fmt.Println("sum of even numbers 1..100:", sum)
	fmt.Println("count:", count)
}
