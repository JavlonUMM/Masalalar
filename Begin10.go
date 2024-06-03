package main

import "fmt"

func main() {
    var a, b int
    
    fmt.Print("Nolga teng bo'lmagan ikkita sonni kiriting (a va b): ")
    fmt.Scan(&a, &b)
    
    Yigindisi := a + b
    Kopaytmasi := a * b
    a_ning_kvadrati := a * a
    b_ning_kvadrati := b * b
    
    fmt.Printf("Ikkita sonning yig'indisi: %.d\n", Yigindisi)
    fmt.Printf("Ikkita sonning ko'paytmasi: %.d\n", Kopaytmasi)
    fmt.Printf("Birinchi sonning kvadrati: %.d\n", a_ning_kvadrati)
    fmt.Printf("Ikkinchi sonning kvadrati: %.d\n", b_ning_kvadrati)
}
