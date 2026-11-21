package main

import "fmt"

func main() {
	apples := 10
	pears := 3

	fmt.Println("Sum:", apples+pears)
	fmt.Println("Difference:", apples-pears)
	fmt.Println("Product:", apples*pears)

	price := 19.5
	quantity := 2.0
	total := price * quantity
	fmt.Println("Total:", total)

	result := 2 + 3*4
	fmt.Println("2 + 3 * 4 =", result)

	result = (2 + 3) * 4
	fmt.Println("(2 + 3) * 4 =", result)

	score := 100
	score += 50
	fmt.Println("Score after +=", score)
	score -= 30
	fmt.Println("Score after -=", score)
	score *= 2
	fmt.Println("Score after *=", score)

	fmt.Println("7 / 2 =", 7/2)
	fmt.Println("7 % 2 =", 7%2)

	fmt.Println("7.0 / 2.0 =", 7.0/2.0)

	candies := 17
	kids := 5
	perKid := candies / kids
	left := candies % kids
	fmt.Println("Each kid gets:", perKid)
	fmt.Println("Left over:", left)

	fmt.Println("Check:", perKid*kids+left)

	totalMinutes := 135
	hours := totalMinutes / 60
	minutes := totalMinutes % 60
	fmt.Println(totalMinutes, "minutes =", hours, "h", minutes, "min")

	number := 14
	fmt.Println(number, "% 2 =", number%2)

	fmt.Println("-7 / 2 =", -7/2)
	fmt.Println("-7 % 2 =", -7%2)

	
}
