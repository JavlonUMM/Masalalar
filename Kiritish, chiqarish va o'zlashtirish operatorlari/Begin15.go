package main

import "fmt"

func main() {
    var S float64
    const pi = 3.14
    
    fmt.Print("Aylananing yuzasini kiriting (S): ")
    fmt.Scan(&S)
	
	R := S / (2 * pi)
	d := 2 * R
    
    fmt.Printf("Aylananing radiusi (R): %.2f\n", R)
	fmt.Printf("Aylananing diametri (d): %.2f\n", d)
}
