package main

import "fmt"

func main() {
    var L float64
    const pi = 3.14
    
    fmt.Print("Aylananing uzunligini kiriting (L): ")
    fmt.Scan(&L)
	
    R := L / (2 * pi)
    
    S := pi * (R * R)
    
    fmt.Printf("Aylananing radiusi (R): %.2f\n", R)
    fmt.Printf("Aylananing yuzasi (S): %.2f\n", S)
}
