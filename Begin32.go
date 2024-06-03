package main

import "fmt"

func main() {
	var TC float64
	fmt.Print("Temperaturani selsiyda kiriting: ")
	fmt.Scan(&TC)

	TF := (TC * 9 / 5) + 32

	fmt.Printf("Farengeytda %2f\n", TF)
}
