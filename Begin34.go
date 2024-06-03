package main

import "fmt"

func main() {
    var X, A, Y, B float64
    fmt.Print("Shokolad kilogrammini kiriting: ")
    fmt.Scan(&X)
    fmt.Print("1 kg shokolad narxini kiriting: ")
    fmt.Scan(&A)
    fmt.Print("Konfet kilogrammini kiriting: ")
    fmt.Scan(&Y)
    fmt.Print("1 kg konfet narxini kiriting: ")
    fmt.Scan(&B)

    shokalad_narxi := A

    konfet_narxi := B

    qimmatligi := shokalad_narxi / konfet_narxi

    fmt.Printf("1 kg shokolad 1 kg konfetdan %.2f marta qimmatdir\n", qimmatligi)
}
