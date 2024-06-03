package main

import "fmt"

func main() {
    var son int
    fmt.Print("Uch xonali son kiriting: ")
    fmt.Scan(&son)
    
    yuzlar := (son / 100) % 10
    
    fmt.Printf("Yuzlar xonasu: %d\n", yuzlar)
}
