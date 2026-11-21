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

	price2 := 19.5
	quantity2 := 3
	total2 := price2 * float64(quantity2)
	fmt.Println("Total:", total2)

	sum := 17
	count := 5
	fmt.Println("Integer average:", sum/count)
	fmt.Println("Float average:", float64(sum)/float64(count))

	fmt.Println("Too late:", float64(sum/count))

	temperature := 23.9
	rounded := int(temperature)
	fmt.Println("int(23.9) =", rounded)
	cold := -23.9
	fmt.Println("int(-23.9) =", int(cold))

	fmt.Println("temerature:", temperature)
	fmt.Printf("Types: %T and %T\n", temperature, rounded)
}
