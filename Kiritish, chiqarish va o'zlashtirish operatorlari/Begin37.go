package main

import "fmt"

func main() {
    var V1, V2, S float64
    fmt.Print("Birinchi avtomobilning tezligini kiriting: ")
    fmt.Scan(&V1)
    fmt.Print("Ikkinchi avtomobilning tezligini kiriting: ")
    fmt.Scan(&V2)
    fmt.Print("Ular orasidagi masofani kiriting: ")
    fmt.Scan(&S)

    T := S / (V1 + V2)

    fmt.Printf("Ular orasidagi masofa %v soatdan keyin bosiladi\n", T)
}
