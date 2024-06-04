package main

import (
	"fmt"
	"math"
)

func main() {
    var a, b float64
    
    fmt.Print("Ikkita manfiy bo'lmagan sonlarni kiriting: ")
    fmt.Scan(&a)
    fmt.Print("Ikkita manfiy bo'lmagan sonlarni kiriting: ")
    fmt.Scan(&b)

    Geometrik := math.Sqrt(a * b)
    
    fmt.Printf("Sonlarning o'rta geometrigi: %.2f\n", Geometrik)
}
