package main

import "fmt"

func main() {
    var x float64
    fmt.Print("x ni kiriting: ")
    fmt.Scan(&x)
    
    y := 3*x*x*x*x*x*x - 6*x*x - 7
    
    fmt.Printf("y ning qiymati: %.2f\n", y)
}
