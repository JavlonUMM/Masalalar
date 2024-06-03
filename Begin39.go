package main

import (
    "fmt"
    "math"
)

func main() {
    // Koeffisentlarni kiritamiz
    var A, B, C float64
    fmt.Print("A ni kiriting: ")
    fmt.Scan(&A)
    fmt.Print("B ni kiriting: ")
    fmt.Scan(&B)
    fmt.Print("C ni kiriting: ")
    fmt.Scan(&C)

    D := math.Pow(B, 2) - 4*A*C

    if D < 0 {
        fmt.Println("Tenglama ildizlarga ega emas.")
    } else {
        x1 := (-B + math.Sqrt(D)) / (2 * A)
        x2 := (-B - math.Sqrt(D)) / (2 * A)
        fmt.Printf("Tenglamaning yechimlari: %2f, %2f\n", x1, x2)
    }
}
