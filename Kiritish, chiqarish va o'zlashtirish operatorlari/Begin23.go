package main

import "fmt"

func main() {
    var A, B, C float64
    fmt.Print("A ni kiriting: ")
    fmt.Scan(&A)
    fmt.Print("B ni kiriting: ")
    fmt.Scan(&B)
    fmt.Print("C ni kiriting: ")
    fmt.Scan(&C)
    
    ALmashtirish := A
    A = B
    B = C
    C = ALmashtirish

    fmt.Println("A ning qiymati:", A)
    fmt.Println("B ning qiymati:", B)
    fmt.Println("C ning qiymati:", C)
}
