package main

import "fmt"

func main() {
    var son int
    fmt.Print("Ikki xonali sonni kiriting: ")
    fmt.Scan(&son)
    
    onliklar := son / 10
    birlar := son % 10
    
    fmt.Printf("O'nliklar xonasi: %d\n", onliklar)
    fmt.Printf("Birlar xonasi: %d\n", birlar)
}
