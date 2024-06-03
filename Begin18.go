package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64
	fmt.Print("A nuqtani kiriting: ")
	fmt.Scan(&A)
	fmt.Print("B nuqtani koordinatasini kiriting: ")
	fmt.Scan(&B)
	fmt.Print("C nuqtani koordinatasini kiriting: ")
	fmt.Scan(&C)

	AC := math.Abs(C - A)
	BC := math.Abs(C - B)
	kopaytmasi := AC * BC

	fmt.Printf("AC kesma uzunligi = %.2f\n", AC)
	fmt.Printf("BC kesma uzunligi = %.2f\n", BC)
	fmt.Printf("AC va BC kesmalar uzunligi ko'paytmasi = %.2f\n", kopaytmasi)
}
