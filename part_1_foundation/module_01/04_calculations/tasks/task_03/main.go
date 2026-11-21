package main

import "fmt"

func main() {
	a := 17
	b := 4
	c := 2.5

	fmt.Println(a / b)                   // 4
	fmt.Println(a % b)                   // 1
	fmt.Println(float64(a) / float64(b)) // 4.25
	fmt.Println(float64(a / b))          // 4
	fmt.Println(int(c) * b)              // 8
	fmt.Println(c * float64(b))          // 10

	a += 3
	b *= 2
	fmt.Printf("a=%d b=%d\n", a, b)         // a=20 b=8
	fmt.Printf("%.1f|%5d|%-5d|\n", c, a, b) //2.5|   20|8    |
	fmt.Printf("%d%%\n", a/b*100)           // 200%
	fmt.Printf("%v %T\n", a/b, c)           //2 float64
}
