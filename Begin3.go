package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("To'g'ri to'rtburchakning tomonlarini kiriting (a va b): ")
    fmt.Scan(&a, &b)
    
    S := a * b
    P := 2 * (a + b)
    
    fmt.Printf("To'g'ri to'rtburchakning yuzasi: %d\n", S)
    fmt.Printf("To'g'ri to'rtburchakning perimetri: %d\n", P)
}
