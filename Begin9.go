package main

import "fmt"

func main() {
    var a, b float64
    
    fmt.Print("Ikkita manfiy bo'lmagan sonlarni kiriting (a va b): ")
    fmt.Scan(&a, &b)
    
    Geometrik := (a * b)/2
    
    fmt.Printf("Sonlarning o'rta geometrigi: %.2f\n", Geometrik)
}
