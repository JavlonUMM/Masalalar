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
    A4 := math.Pow(A2, 2)
    A8 := math.Pow(A4, 2)
    
    fmt.Printf("A ning kvadrati (A^2): %.2f\n", A2)
    fmt.Printf("A ning to'rtinchi darajasi (A^4): %.2f\n", A4)
    fmt.Printf("A ning sakkizinchi darajasi (A^8): %.2f\n", A8)
}
