package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64
	fmt.Print("A nuqtasini kiriting: ")
	fmt.Scan(&A)
	fmt.Print("B nuqtasini kiriting: ")
	fmt.Scan(&B)
	fmt.Print("C nuqtasini kiriting: ")
	fmt.Scan(&C)

	AC := math.Abs(C - A)
	BC := math.Abs(C - B)

	yigindi := AC + BC

	fmt.Printf("AC kesma uzunligi = %.2f\n", AC)
	fmt.Printf("BC kesma uzunligi = %.2f\n", BC)
	fmt.Printf("AC va BC kesmalarini uzunliklari yig'indisi = %.2f\n", yigindi)
}
