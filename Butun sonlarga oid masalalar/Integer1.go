package main

import "fmt"

func main() {
    var L int
    fmt.Print("Santimetr kiriting: ")
    fmt.Scan(&L)
    
    metr := L / 100
    
    fmt.Printf("Metr: %d\n", metr)
}
