package main

import "fmt"

func main() {
    var son int
    fmt.Print("999 dan katta son kiriting: ")
    fmt.Scan(&son)
    
    Natija := (son / 100) % 10
    
    fmt.Printf("Sonning yuzliklar xonasi: %d\n", Natija)
}
