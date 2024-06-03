package main

import "fmt"

func main() {
    var a, b int
    
    fmt.Print("Nolga teng bo'lmagan ikkita sonni kiriting (a va b): ")
    fmt.Scan(&a, &b)
    
    Yigindisi := a + b
    Kopaytmasi := a * b
    a_ning_moduli := (a)
    b_ning_moduli := (b)
    
    fmt.Printf("Ikkita sonning yig'indisi: %.d\n", Yigindisi)
    fmt.Printf("Ikkita sonning ko'paytmasi: %.d\n", Kopaytmasi)
    fmt.Printf("Birinchi sonning moduli: %.d\n", a_ning_moduli)
    fmt.Printf("Ikkinchi sonning moduli: %.d\n", b_ning_moduli)
}

// 2-usul

/*package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b float64
    
    fmt.Print("Ikkita sonni kiriting (nolga teng bo'lmagan a va b): ")
    fmt.Scan(&a, &b)
    
    sum := a + b
    product := a * b
    abs_a := math.Abs(a)
    abs_b := math.Abs(b)
    
    fmt.Printf("Ikkita sonning yig'indisi: %.2f\n", sum)
    fmt.Printf("Ikkita sonning ko'paytmasi: %.2f\n", product)
    fmt.Printf("Birinchi sonning moduli: %.2f\n", abs_a)
    fmt.Printf("Ikkinchi sonning moduli: %.2f\n", abs_b)
}
