package main

import (
    "fmt"
    "math"
)


func main() {
    var x float64
    fmt.Print("x ni kiriting: ")
    fmt.Scan(&x)
    
    y := 4*(math.Pow(x-3, 6)) - 7*(math.Pow(x-3, 3)) + 2
    
    fmt.Printf("y ning qiymati: %.2f\n", y)
}
