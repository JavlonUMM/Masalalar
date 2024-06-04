package main

import "fmt"

func main() {
    var x1, x2 float64
    
    fmt.Print("Birinchi sonni kiriting (x1): ")
    fmt.Scan(&x1)
    
    fmt.Print("Ikkinchi sonni kiriting (x2): ")
    fmt.Scan(&x2)
    
    Orasidagi_masofa := (x2 - x1)
    
    fmt.Printf("Ikkita son o'rtasidagi masofa: %.2f\n",  Orasidagi_masofa)
}
