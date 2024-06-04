package main

import "fmt"

func main() {
    var a int
    fmt.Print("Kubning tomoni kiriting: ")
    fmt.Scan(&a)
    
    V := a * a * a
    S := 6 * a * a
    
    fmt.Printf("Kubning hajmi: %d\n", V)
    fmt.Printf("Kubning to'la sirti: %d\n", S)
}
