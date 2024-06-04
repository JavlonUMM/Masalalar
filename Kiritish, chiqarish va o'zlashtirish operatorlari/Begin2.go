package main

import (
	"fmt"
	"math"
)

func main() {
    var a float64
    fmt.Print("Kvadratning tomoni kiriting: ")
    fmt.Scan(&a)
    
    S := math.Pow(a,10)
    fmt.Printf("Kvadratning yuzasi: %d\n", S)
}
