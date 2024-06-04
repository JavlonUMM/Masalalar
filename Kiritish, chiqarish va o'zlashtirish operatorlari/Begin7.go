package main

import "fmt"

func main() {
    var R float64
    const pi = 3.14
    
    fmt.Print("Doiraning radiusini kiriting: ")
    fmt.Scan(&R)
    
    L := 2 * pi * R
    S := pi * R * R
    
    fmt.Printf("Doiraning uzunligi: %.2f\n", L)
    fmt.Printf("Doiraning yuzasi: %.2f\n", S)
}
