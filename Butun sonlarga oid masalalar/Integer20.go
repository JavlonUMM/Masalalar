package main

import "fmt"

func main() {
    var sekund int
    fmt.Print("Sekundni kiriting: ")
    fmt.Scan(&sekund)
    
    soat := sekund / 3600
    
    fmt.Printf("Soatlar soni: %d\n", soat)
}
