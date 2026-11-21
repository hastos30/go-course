package main

import "fmt"

func main() {
	var count int
	var label string
	count = count + 3
	label = "items"
	total := count
	count = count * 2
	fmt.Println(count, total, label)
	const limit = 10
	fmt.Println(limit - total)
}
