
package main

import "fmt"

func main() {
    // Radianda berilgan burchakni kiriting
    var alfa float64
    fmt.Print("Radianda berilgan burchakni kiriting (0 < alfa < 2π): ")
    fmt.Scan(&alfa)

	const pi = 3.14
    
    // Radianni gradusga o'tkazish
    daraja := alfa * 180 / pi
    
    fmt.Printf("Berilgan burchakning qiymati gradusda: %.2f\n", daraja)
}
