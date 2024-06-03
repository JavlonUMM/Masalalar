package main

import "fmt"

func main() {
    var a int
    fmt.Print("Kvadratning tomoni kiriting: ")
    fmt.Scan(&a)
    
    S := a * a
    fmt.Printf("Kvadratning yuzasi: %d\n", S)
}
