package main

import "fmt"

func main() {
    var a int
    fmt.Print("Kvadratning tomoni kiriting: ")
    fmt.Scan(&a)
    
    P := 4 * a
    fmt.Printf("Kvadratning perimetri: %.d\n", P)
}
