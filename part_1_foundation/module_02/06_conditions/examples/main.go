package main

import "fmt"

func main() {
	age := 17
	price := 99.5
	name := "Viktor"

	fmt.Println(age == 18)
	fmt.Println(age != 18)
	fmt.Println(age < 18)
	fmt.Println(age >= 18)
	fmt.Println(price > 100.0)
	fmt.Println(name == "Viktor")

	isAdult := age >= 18
	fmt.Printf("isAdult = %t, type %T\n", isAdult, isAdult)

	temperature := 25

	if temperature > 30 {
		fmt.Println("Hot")
	}

	if age >= 18 {
		fmt.Println("Access granted")
	} else {
		fmt.Println("Access denied")
	}

	score := 74
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}
	fmt.Println("Done")

	hasTicket := true
	isVIP := false

	if age >= 18 && hasTicket {
		fmt.Println("Welcome to the concert")
	} else {
		fmt.Println("Entry not allowed")
	}

	if hasTicket || isVIP {
		fmt.Println("Can enter the hall")
	}

	if !isVIP {
		fmt.Println("Regular seat")
	}

	hour := 14
	isWorkingTime := hour >= 9 && hour < 18
	fmt.Println("Working time:", isWorkingTime)

	day := 6

	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Unknown day")
	}

	command := "stop"
	switch command {
	case "start":
		fmt.Println("Starting...")
	case "stop":
		fmt.Println("Stopping...")
	default:
		fmt.Println("Unknown command:", command)
	}

	switch {
	case score >= 90:
		fmt.Println("Switch grade: A")
	case score >= 75:
		fmt.Println("Switch grade: B")
	case score >= 60:
		fmt.Println("Switch grade: C")
	default:
		fmt.Println("Switch grade: F")
	}

	month := 11
	switch month {
	case 12, 1, 2:
		fmt.Println("Winter")
	case 6, 7, 8:
		fmt.Println("Summer")
	default:
		fmt.Println("Other season")
	}
}
