package main

import "fmt"

func main() {
    var A, B float64
    fmt.Print("A ni kiriting: ")
    fmt.Scan(&A)
    fmt.Print("B ni kiriting: ")
    fmt.Scan(&B)
    
    ALmashtirish := A
    A = B
    B = ALmashtirish
    
    fmt.Println("A ning qiymati:", A)
    fmt.Println("B ning qiymati:", B)
}
