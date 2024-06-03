package main

import "fmt"

func main() {
    var a, b float64
    
    fmt.Print("Ikkita sonni kiriting (a va b): ")
    fmt.Scan(&a, &b)
    
    Arifmetik := (a + b) / 2
    
    fmt.Printf("Sonlarning o'rta arifmetigi: %.2f\n", Arifmetik)
}
