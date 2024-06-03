package main

import "fmt"

func main() {
    var sekund int
    fmt.Print("Sekundni kiriting: ")
    fmt.Scan(&sekund)

    minut := sekund / 60
    
    fmt.Printf("O'tgan minut: %d\n", minut)
}
