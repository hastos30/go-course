package main

import "fmt"

func main() {
	count := 0
	number := 1000
	temp := number
	for {
		if temp == 1 {
			break
		}
		temp /= 2
		count++
	}

	fmt.Println("Source number:", number)
	fmt.Println("Counter divided by 2:", count)
	fmt.Println("Result:", temp)
}
