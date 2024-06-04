package main

import "fmt"

func main() {
	var TF float64
	fmt.Print("Temperaturani Farengeytda kiriting: ")
	fmt.Scan(&TF)

	TC := (TF - 32) * 5 / 9

	fmt.Printf("Selsiyda %2f\n", TC)
}
