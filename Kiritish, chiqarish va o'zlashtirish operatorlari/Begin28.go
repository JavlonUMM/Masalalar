package main

import (
	"fmt"
	"math"
)

func main() {
    var A float64
    fmt.Print("A ni kiriting: ")
    fmt.Scan(&A)
    
    A2 := math.Pow(A, 2)
    A3 := A2 * A
    A5 := A3 * A2
    A10 := math.Pow(A5, 2)
    A15 := A10 * A5
    
    fmt.Printf("A ning kvadrati (A^2): %.2f\n", A2)
    fmt.Printf("A ning kubi (A^3): %.2f\n", A3)
    fmt.Printf("A ning beshinchi darajasi (A^5): %.2f\n", A5)
    fmt.Printf("A ning o'ninchi darajasi (A^10): %.2f\n", A10)
    fmt.Printf("A ning o'n beshinchi darajasi (A^15): %.2f\n", A15)
}
