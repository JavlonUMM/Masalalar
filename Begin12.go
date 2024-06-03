package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b float64
    
    fmt.Print("To'g'ri uchburchakning katetlarini kiriting (a va b): ")
    fmt.Scan(&a, &b)
    
    c := math.Sqrt(a*a + b*b)
    P := a + b + c
    
    fmt.Printf("Uchburchakning gipotenuzasi: %.2f\n", c)
    fmt.Printf("Uchburchakning perimetri: %.2f\n", P)
}
