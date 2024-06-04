package main

import "fmt"

func main() {
    var A, B float64
    fmt.Print("A ni kiriting: ")
    fmt.Scan(&A)
    fmt.Print("B ni kiriting: ")
    fmt.Scan(&B)

    x := -B / A
    
	fmt.Printf("Tenglamaning yechimi: %2f\n", x)
}
