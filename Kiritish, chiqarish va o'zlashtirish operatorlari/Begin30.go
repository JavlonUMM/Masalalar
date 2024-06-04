
package main

import "fmt"

func main() {
    var alfa float64
    fmt.Print("Burchakni kiriting (0 < alfa < 2π): ")
    fmt.Scan(&alfa)

	const pi = 3.14
    
    daraja := alfa * 180 / pi
    
    fmt.Printf("Gradusda: %.2f\n", daraja)
}
