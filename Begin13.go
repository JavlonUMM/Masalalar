package main

import (
	"fmt"
	"math"
)

func main() {
	var R1, R2 float64
	fmt.Print("R1 radiusini kiriting: ")
	fmt.Scan(&R1)
	fmt.Print("R2 radiusini kiriting: ")
	fmt.Scan(&R2)

	S1 := math.Pi * math.Pow(R1, 2)
	S2 := math.Pi * math.Pow(R2, 2)

	S3 := math.Pi * (R1 - R2) * (R1 + R2)

	fmt.Printf("S1 = %.2f\n", S1)
	fmt.Printf("S2 = %.2f\n", S2)
	fmt.Printf("S3 = %.2f\n", S3)
}
