package main

import "fmt"

func main() {
    var A, B, C float64
    
    fmt.Print("Birinchi nuqtani kiriting (A): ")
    fmt.Scan(&A)
    
    fmt.Print("Ikkinchi nuqtani kiriting (B): ")
    fmt.Scan(&B)

	fmt.Print("Uchinchi nuqtani kiriting (C): ")
    fmt.Scan(&C)
    
    AC := C - A
	BC := C - B
	Yigindisi := AC * BC

    
    fmt.Printf("AC kesmaning uzunligi: %.2f\n",  AC)
	fmt.Printf("BC kesmaning uzunligi: %.2f\n",  BC)
	fmt.Printf("Kesmaning yigindisi: %.2f\n",  Yigindisi)
}
