package main

import "fmt"

func main() {
    var a, b, c int
    fmt.Print("Paralelepepidning tomonlati kiriting (a, b, c): ")
    fmt.Scan(&a, &b, &c)
    
    V := a * b* c
    S := 2 * (a * b + b * c + a * c)
    
    fmt.Printf("Paralelepepidning hajmi: %d\n", V)
    fmt.Printf("Paralelepepidning to'la sirti: %d\n", S)
}
