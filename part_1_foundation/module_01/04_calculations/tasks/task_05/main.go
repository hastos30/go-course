package main

import "fmt"

func main() {
	// Программа считает расход топлива: литры на 100 км
	distance := 420
	fuel := 31.5

	consumption := fuel / float64(distance) * 100 // привел int к float64
	fmt.Printf("Consumption: %.2f l/100km\n", consumption)

	// Сколько полных баков по 50 литров нужно на 1000 км
	trip := 1000
	tank := 50
	litersNeeded := float64(trip) * consumption / 100
	fullTanks := float64(int(litersNeeded) / tank)
	fmt.Printf("Liters needed: %.1f\n", litersNeeded) // передал нужный аргумент
	fmt.Printf("Full tanks: %.f\n", fullTanks)        // исправил %d на %.f
}
