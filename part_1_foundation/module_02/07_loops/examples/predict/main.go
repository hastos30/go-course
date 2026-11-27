package main

import "fmt"

func main() {
	total := 0
	for i := 10; i > 0; i -= 3 {
		if i == 4 {
			continue
		}
		total += i
		fmt.Println("i:", i, "total:", total)
	}
	fmt.Println("result:", total)

	count := 0
	x := 1
	for {
		x *= 2
		count++
		if x >= 20 {
			break
		}
	}
	fmt.Println("x:", x, "count:", count)
}
