package main

import "fmt"

func main() {
	totalSeconds := 86399

	seconds := totalSeconds % 60
	minutes := (totalSeconds / 60) % 60
	hours := (totalSeconds / 60) / 60

	fmt.Printf("%d seconds = %d h %d min %d s\n", totalSeconds, hours, minutes, seconds)
}
