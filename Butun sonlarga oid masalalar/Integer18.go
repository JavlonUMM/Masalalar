package main

import "fmt"

func main() {
    var son int
    fmt.Print("999 dan katta son kiriting: ")
    fmt.Scan(&son)
    
    Natija := (son / 1000) % 10
    
    fmt.Printf("Sonning minglik xonasi: %d\n", Natija)
}
