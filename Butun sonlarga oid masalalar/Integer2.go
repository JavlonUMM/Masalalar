package main

import "fmt"

func main() {
    var M int
    fmt.Print("Kilogrammni kiriting: ")
    fmt.Scan(&M)
    
    tonna := M / 1000
    
    fmt.Printf("Undagi to'liq tonnalar soni: %d\n", tonna)
}
