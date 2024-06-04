package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a three-digit number: ")
	fmt.Scan(&number)

	hundreds := number / 100
	tens := (number % 100) / 10
	units := number % 10

	result := hundreds*100 + tens*10 + units

	fmt.Printf("Result: %d\n", result)
}
