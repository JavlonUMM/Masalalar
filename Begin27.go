package main

import "fmt"

func main() {
    var A float64
    fmt.Print("A ni kiriting: ")
    fmt.Scan(&A)
    
    A2 := A * A
    A4 := A2 * A2
    A8 := A4 * A4
    
    fmt.Printf("A ning kvadrati (A^2): %.2f\n", A2)
    fmt.Printf("A ning to'rtinchi darajasi (A^4): %.2f\n", A4)
    fmt.Printf("A ning sakkizinchi darajasi (A^8): %.2f\n", A8)
}
