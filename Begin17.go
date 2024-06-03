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

<<<<<<< HEAD
	AC := math.Abs(C - A)
	BC := math.Abs(C - B)
=======
	fmt.Print("Uchinchi nuqtani kiriting (C): ")
    fmt.Scan(&C)
    
    AC := C - A
	BC := C - B
	Yigindisi := AC * BC
>>>>>>> 0fd0c8e716e3c1f20b4cef7893806467050da676

	yigindi := AC + BC

	fmt.Printf("AC kesma uzunligi = %.2f\n", AC)
	fmt.Printf("BC kesma uzunligi = %.2f\n", BC)
	fmt.Printf("AC va BC kesmalarini uzunliklari yig'indisi = %.2f\n", yigindi)
}
