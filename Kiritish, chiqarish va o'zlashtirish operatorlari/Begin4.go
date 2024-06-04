package main

import "fmt"

func main() {
    var d float64
    const pi = 3.14
    
    fmt.Print("Aylananing diametrini kiriting: ")
    fmt.Scan(&d)
    
    L := pi * d
    
    fmt.Printf("Aylananing uzunligi: %.2f\n", L)
}
