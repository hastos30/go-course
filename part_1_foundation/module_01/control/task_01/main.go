package main

import "fmt"

const (
	metersInKm      = 1000
	secondsInMinute = 60
	secondsInHour   = 3600
)

func main() {
	distanceMeters := 42195
	totalSeconds := 12345

	distanceKm := float64(distanceMeters) / metersInKm

	hours := totalSeconds / secondsInHour
	remainingSeconds := totalSeconds % secondsInHour
	minutes := remainingSeconds / secondsInMinute
	seconds := remainingSeconds % secondsInMinute

	totalHours := float64(totalSeconds) / secondsInHour
	speed := distanceKm / totalHours

	fmt.Printf("Distance: %.3f km\n", distanceKm)
	fmt.Printf("Time: %dh %dm %ds\n", hours, minutes, seconds)
	fmt.Printf("Speed: %.2f km/h\n", speed)
}
