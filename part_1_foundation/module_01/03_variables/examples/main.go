package main

import "fmt"

func main() {
	var age int = 25
	fmt.Println("Age:", age)

	age = 26
	fmt.Println("Age:", age)

	age = age + 1
	fmt.Println("Age:", age)

	var name string = "Viktor"
	fmt.Println("Name:", name)

	var price float64 = 19.99
	var isActive bool = true
	fmt.Println("Price:", price)
	fmt.Println("Active:", isActive)

	var emptyInt int
	var emptyFloat float64
	var emptyString string
	var emptyBool bool

	fmt.Println("Zero int:", emptyInt)
	fmt.Println("Zero float64:", emptyFloat)
	fmt.Println("Zero string:", emptyString, "<- end")
	fmt.Println("Zero bool:", emptyBool)

	city := "Kyiv"
	year := 2026
	temperature := 21.5
	fmt.Println(city, year, temperature)
	fmt.Printf("%T %T %T\n", city, year, temperature)

	const pi = 3.14159
	radius := 2.0
	area := pi * radius * radius
	fmt.Println("Area:", area)

	x := 10
	y := x
	x = 20
	fmt.Println("x:", x, "y:", y)

}
