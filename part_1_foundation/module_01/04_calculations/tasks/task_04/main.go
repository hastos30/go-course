package main

import (
	"fmt"
	"math"
)

func main() {
	grade1 := 86
	grade2 := 92
	grade3 := 79

	middleGrade := (float64(grade1) + float64(grade2) + float64(grade3)) / 3
	fmt.Printf("%.2f\n", middleGrade)

	middleGrade = math.Round(middleGrade)
	fmt.Printf("%d\n", int(middleGrade))
}
