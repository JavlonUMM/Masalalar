package main

import "fmt"

func main() {
    var alfa float64
    fmt.Print("Burchakni kiriting (0° < alfa < 360°): ")
    fmt.Scan(&alfa)
	const pi = 3.14
    
    radian := alfa * (pi / 180)
    
    fmt.Printf("Radianda: %.2f\n", radian)
}
