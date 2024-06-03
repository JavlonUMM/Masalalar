package main

import "fmt"

func main() {
    var A, B int
    fmt.Print("A ni kiriting: ")
    fmt.Scan(&A)
    fmt.Print("B ni kiriting: ")
    fmt.Scan(&B)
    
    C := A / B
    
    Qoldiq := A % B
    
    fmt.Printf("B kesmada A kesma: %d\n", C)
    fmt.Printf("Joylashmagan qismi: %d\n", Qoldiq)
}
