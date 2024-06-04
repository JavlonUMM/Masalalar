package main

import "fmt"

func main() {
    var a, b float64
    
    fmt.Print("Nolga teng bo'lmagan ikkita sonni kiriting (a va b): ")
    fmt.Scan(&a, &b)
    
    Yigindisi := a + b
    Kopaytmasi := a * b
    a_ning_kvadrati := a * a
    b_ning_kvadrati := b * b
    
    fmt.Printf("Ikkita sonning yig'indisi: %.2f\n", Yigindisi)
    fmt.Printf("Ikkita sonning ko'paytmasi: %.2f\n", Kopaytmasi)
    fmt.Printf("Birinchi sonning kvadrati: %.2f\n", a_ning_kvadrati)
    fmt.Printf("Ikkinchi sonning kvadrati: %.2f\n", b_ning_kvadrati)
}
