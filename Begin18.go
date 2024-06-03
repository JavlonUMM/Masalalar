package main

import "fmt"

func main() {
    var A, B, C float64

	fmt.Println("A, B va C nuqtalarini kiriting (A < C < B):")
    
    fmt.Print("Birinchi nuqtani kiriting (A): ")
    fmt.Scan(&A)
    
    fmt.Print("Ikkinchi nuqtani kiriting (B): ")
    fmt.Scan(&B)

	fmt.Print("Uchinchi nuqtani kiriting (C): ")
    fmt.Scan(&C)
    
    AC := C - A
	BC := B - C
	Kopaytmasi := AC + BC

    
    fmt.Printf("AC kesmaning uzunligi: %.2f\n",  AC)
	fmt.Printf("BC kesmaning uzunligi: %.2f\n",  BC)
	fmt.Printf("Kesmaning yigindisi: %.2f\n",  Kopaytmasi)
}
