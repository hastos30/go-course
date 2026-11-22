package main

import "fmt"

func main() {
	a := 17
	b := 5
	c := a / b
	d := a % b
	e := float64(a) / float64(b)
	f := float64(a / b)
	x := 7
	x += c * d
	fmt.Printf("%d %d %.1f %.1f %d\n", c, d, e, f, x)
	fmt.Printf("%v %T\n", 10/4*2.0, 10/4*2.0)

}

// 3 2 3.4 3.0 13
// 4 float64
