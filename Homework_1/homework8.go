package main

import "fmt"

func main() {
	var second int

	fmt.Print("Enter the second: ")
	fmt.Scan(&sekund)

	hour := second / 3600
	minute := (second % 3600) / 60
	seconds := (second % 3600) % 60

	fmt.Printf("Result: %d %d %d \n", hour, minute, seconds)
}
